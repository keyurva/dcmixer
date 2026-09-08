// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server implements core Mixer handlers and data routing.
package server

import (
	"context"
	"sort"
	"strings"

	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
	"github.com/datacommonsorg/mixer/internal/server/topic"
	"google.golang.org/protobuf/types/known/structpb"
)

// V2AgentResolvePlaces implements API for mixer.V2AgentResolvePlaces.
// It delegates incoming RPC requests directly to the isolated agent.Service layer.
func (s *Server) V2AgentResolvePlaces(
	ctx context.Context,
	in *pbv2.ResolvePlacesRequest,
) (*pbv2.ResolvePlacesResponse, error) {
	return s.agentService.ResolvePlaces(ctx, in)
}

// V2AgentSearchIndicators implements API for mixer.V2AgentSearchIndicators.
// It delegates incoming RPC requests directly to the isolated agent.Service layer.
func (s *Server) V2AgentSearchIndicators(
	ctx context.Context,
	in *pbv2.SearchIndicatorsRequest,
) (*pbv2.SearchIndicatorsResponse, error) {
	resp, err := s.agentService.SearchIndicators(ctx, in)
	if err != nil || resp == nil {
		return resp, err
	}
	s.populateConstraintPropertiesTable(ctx, resp)
	return resp, nil
}

// V2AgentGetObservations implements API for mixer.V2AgentGetObservations.
// It delegates incoming RPC requests directly to the isolated agent.Service layer.
func (s *Server) V2AgentGetObservations(
	ctx context.Context,
	in *pbv2.GetObservationsRequest,
) (*pbv2.GetObservationsResponse, error) {
	return s.agentService.GetObservations(ctx, in)
}

// V2AgentGetVariableMetadata implements API for mixer.V2AgentGetVariableMetadata.
// It delegates incoming RPC requests directly to the isolated agent.Service layer.
func (s *Server) V2AgentGetVariableMetadata(
	ctx context.Context,
	in *pbv2.GetVariableMetadataRequest,
) (*pbv2.GetVariableMetadataResponse, error) {
	return s.agentService.GetVariableMetadata(ctx, in)
}

// V2AgentInspectIndicatorNodes implements API for mixer.V2AgentInspectIndicatorNodes.
// It delegates incoming RPC requests directly to the isolated agent.Service layer.
func (s *Server) V2AgentInspectIndicatorNodes(
	ctx context.Context,
	in *pbv2.InspectIndicatorNodesRequest,
) (*pbv2.InspectIndicatorNodesResponse, error) {
	return s.agentService.InspectIndicatorNodes(ctx, in)
}

// GetStatVarDimensionsByPrefix delegates to the underlying SpannerClient if available.
func (s *Server) GetStatVarDimensionsByPrefix(ctx context.Context, seedDcid string, properties []string) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error) {
	if sc, ok := s.spannerStalenessTimestampProvider.(interface {
		GetStatVarDimensionsByPrefix(ctx context.Context, seedDcid string, properties []string) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error)
	}); ok && sc != nil {
		return sc.GetStatVarDimensionsByPrefix(ctx, seedDcid, properties)
	}
	return nil, nil
}

// GetStatVarConstraintPropertiesByPrefix delegates to the underlying SpannerClient if available.
func (s *Server) GetStatVarConstraintPropertiesByPrefix(ctx context.Context, dcids []string) ([]string, error) {
	if sc, ok := s.spannerStalenessTimestampProvider.(interface {
		GetStatVarConstraintPropertiesByPrefix(ctx context.Context, dcids []string) ([]string, error)
	}); ok && sc != nil {
		return sc.GetStatVarConstraintPropertiesByPrefix(ctx, dcids)
	}
	return nil, nil
}

// V2AgentGetStatVarsByConstraints implements API for mixer.V2AgentGetStatVarsByConstraints.
func (s *Server) V2AgentGetStatVarsByConstraints(
	ctx context.Context,
	in *pbv2.GetStatVarsByConstraintsRequest,
) (*pbv2.GetStatVarsByConstraintsResponse, error) {
	if sc, ok := s.spannerStalenessTimestampProvider.(interface {
		GetStatVarsByConstraints(ctx context.Context, req *pbv2.GetStatVarsByConstraintsRequest) (*pbv2.GetStatVarsByConstraintsResponse, error)
	}); ok && sc != nil {
		return sc.GetStatVarsByConstraints(ctx, in)
	}
	return &pbv2.GetStatVarsByConstraintsResponse{SeedDcid: in.GetSeedDcid()}, nil
}

// populateConstraintPropertiesTable extracts candidate StatVar DCIDs from SearchIndicatorsResponse
// and builds a consolidated pbv2.Table with columns: ["dcid", "constraint_properties"] using in-memory TopicCache.
func (s *Server) populateConstraintPropertiesTable(ctx context.Context, resp *pbv2.SearchIndicatorsResponse) {
	seen := make(map[string]bool)
	var candidateDcids []string

	addCandidate := func(dcid string) {
		if dcid != "" && !seen[dcid] {
			seen[dcid] = true
			candidateDcids = append(candidateDcids, dcid)
		}
	}

	if resp.VariableCandidates != nil {
		for _, row := range resp.VariableCandidates.GetRows() {
			if len(row.GetValues()) > 0 {
				addCandidate(row.GetValues()[0].GetStringValue())
			}
		}
	}
	if resp.TopicCandidates != nil {
		for _, row := range resp.TopicCandidates.GetRows() {
			if len(row.GetValues()) >= 5 {
				for _, mv := range row.GetValues()[4].GetListValue().GetValues() {
					addCandidate(mv.GetStringValue())
				}
			}
		}
	}
	for _, v := range resp.GetVariables() {
		addCandidate(v.GetDcid())
	}
	for _, t := range resp.GetTopics() {
		for _, mv := range t.GetMemberVariables() {
			addCandidate(mv)
		}
	}

	if len(candidateDcids) == 0 || s.topicExpander == nil {
		return
	}

	infoProvider, ok := s.topicExpander.(interface {
		GetStatVarInfos(ctx context.Context, dcids []string) (map[string]*topic.StatVarInfo, error)
	})
	if !ok {
		return
	}

	infos, err := infoProvider.GetStatVarInfos(ctx, candidateDcids)
	if err != nil || len(infos) == 0 {
		return
	}

	table := &pbv2.Table{
		Columns: []string{"dcid", "constraint_properties"},
		Rows:    []*structpb.ListValue{},
	}
	signatureCounts := make(map[string]int)
	for _, dcid := range candidateDcids {
		info := infos[dcid]
		if info == nil || len(info.ConstraintProperties) == 0 {
			continue
		}
		var props []string
		for prop := range info.ConstraintProperties {
			props = append(props, prop)
		}
		sort.Strings(props)
		sig := info.PopulationType + "|" + strings.Join(props, ",")
		if signatureCounts[sig] >= 2 {
			continue
		}
		signatureCounts[sig]++
		var propList []any
		for _, prop := range props {
			propList = append(propList, prop)
		}
		row, err := structpb.NewList([]any{
			dcid,
			propList,
		})
		if err == nil {
			table.Rows = append(table.Rows, row)
		}
	}

	if len(table.Rows) > 0 {
		resp.ConstraintProperties = table
	}
}

