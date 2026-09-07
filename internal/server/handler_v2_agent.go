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
	"strings"
	"sync"

	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
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
// It delegates incoming RPC requests directly to the isolated agent.Service layer
// and populates tabular constraint_values for returned StatVars and topic member variables.
func (s *Server) V2AgentSearchIndicators(
	ctx context.Context,
	in *pbv2.SearchIndicatorsRequest,
) (*pbv2.SearchIndicatorsResponse, error) {
	resp, err := s.agentService.SearchIndicators(ctx, in)
	if err != nil || resp == nil {
		return resp, err
	}
	s.populateConstraintValuesTable(ctx, resp)
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
func (s *Server) GetStatVarDimensionsByPrefix(ctx context.Context, seedDcid string) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error) {
	if sc, ok := s.spannerStalenessTimestampProvider.(interface {
		GetStatVarDimensionsByPrefix(ctx context.Context, seedDcid string) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error)
	}); ok && sc != nil {
		return sc.GetStatVarDimensionsByPrefix(ctx, seedDcid)
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

// populateConstraintValuesTable extracts seed StatVar DCIDs from SearchIndicatorsResponse
// and builds a consolidated pbv2.Table with columns: ["seed_dcid", "property", "value_dcid", "value_name"].
func (s *Server) populateConstraintValuesTable(ctx context.Context, resp *pbv2.SearchIndicatorsResponse) {
	seenSeeds := make(map[string]bool)
	var seedDcids []string

	addSeed := func(dcid string) {
		if dcid != "" && !seenSeeds[dcid] && len(seedDcids) < 20 {
			seenSeeds[dcid] = true
			seedDcids = append(seedDcids, dcid)
		}
	}

	if resp.VariableCandidates != nil {
		for _, row := range resp.VariableCandidates.GetRows() {
			if len(row.GetValues()) > 0 {
				addSeed(row.GetValues()[0].GetStringValue())
			}
		}
	}
	for _, v := range resp.GetVariables() {
		addSeed(v.GetDcid())
	}
	if resp.TopicCandidates != nil {
		for _, row := range resp.TopicCandidates.GetRows() {
			if len(row.GetValues()) > 4 && row.GetValues()[4].GetListValue() != nil {
				for _, val := range row.GetValues()[4].GetListValue().GetValues() {
					addSeed(val.GetStringValue())
				}
			}
		}
	}
	for _, t := range resp.GetTopics() {
		for _, mv := range t.GetMemberVariables() {
			addSeed(mv)
		}
	}

	if len(seedDcids) == 0 {
		return
	}

	type seedDims struct {
		seedDcid string
		dims     []*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary
	}

	results := make([]seedDims, len(seedDcids))
	var wg sync.WaitGroup
	for i, seed := range seedDcids {
		wg.Add(1)
		go func(idx int, dcid string) {
			defer wg.Done()
			dims, err := s.GetStatVarDimensionsByPrefix(ctx, dcid)
			if err == nil && len(dims) > 0 {
				results[idx] = seedDims{seedDcid: dcid, dims: dims}
			}
		}(i, seed)
	}
	wg.Wait()

	table := &pbv2.Table{
		Columns: []string{"seed_dcid", "property", "value_dcid", "value_name"},
		Rows:    []*structpb.ListValue{},
	}
	for _, res := range results {
		if res.seedDcid == "" {
			continue
		}
		for _, dim := range res.dims {
			prop := dim.GetDimension()
			for _, cvStr := range dim.GetConstraintValues() {
				parts := strings.SplitN(cvStr, "|||", 2)
				valDcid := parts[0]
				valName := valDcid
				if len(parts) == 2 && parts[1] != "" {
					valName = parts[1]
				}
				row, err := structpb.NewList([]any{
					res.seedDcid,
					prop,
					valDcid,
					valName,
				})
				if err == nil {
					table.Rows = append(table.Rows, row)
				}
			}
		}
	}

	if len(table.Rows) > 0 {
		resp.ConstraintValues = table
	}
}

