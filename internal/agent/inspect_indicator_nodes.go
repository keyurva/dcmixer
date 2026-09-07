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

package agent

import (
	"context"
	"fmt"
	"sort"
	"strings"

	pbv1 "github.com/datacommonsorg/mixer/internal/proto/v1"
	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
)

const (
	topicDcidPrefix   = "dc/topic/"
	defaultSampleCap  = 5
	exprMemberOf      = "->memberOf"
	exprLinkedMembers = "<-linkedMemberOf"
	exprRelevantVars  = "->relevantVariable"
)

// InspectIndicatorNodesRequest specifies the input DCIDs and optional place filters
// for inspecting statistical indicator metadata and ontology neighborhoods.
type InspectIndicatorNodesRequest struct {
	Dcids      []string `json:"dcids"`
	PlaceDcids []string `json:"place_dcids,omitempty"`
}

// DimensionSliceSummary captures a populated breakdown dimension and sample slice mappings.
type DimensionSliceSummary struct {
	Dimension      string   `json:"dimension"`
	AvailableCount int      `json:"available_count"`
	SampleSlices   []string `json:"sample_slices"`
}

// StatVarInspection contains unified metadata, provenance, and breakdown dimensions for a StatisticalVariable.
type StatVarInspection struct {
	SeedDcid     string                   `json:"seed_dcid"`
	Name         string                   `json:"name,omitempty"`
	Provenances  []string                 `json:"provenances,omitempty"`
	EarliestDate string                   `json:"earliest_date,omitempty"`
	LatestDate   string                   `json:"latest_date,omitempty"`
	RootSvg      string                   `json:"root_svg"`
	Dimensions   []*DimensionSliceSummary `json:"dimensions"`
}

// TopicInspection contains the 1-hop curated headline variables and child topics for a Topic.
type TopicInspection struct {
	TopicDcid         string   `json:"topic_dcid"`
	HeadlineVariables []string `json:"headline_variables"`
	ChildTopics       []string `json:"child_topics"`
}

// InspectIndicatorNodesResult aggregates unified metadata and ontology inspections for Topics and StatVars.
type InspectIndicatorNodesResult struct {
	Topics   []*TopicInspection   `json:"topics,omitempty"`
	StatVars []*StatVarInspection `json:"stat_vars,omitempty"`
}

// InspectIndicatorNodes traverses the indicator ontology and retrieves unified metadata for Topics and StatVars.
func (s *Service) InspectIndicatorNodes(
	ctx context.Context,
	req *InspectIndicatorNodesRequest,
) (*InspectIndicatorNodesResult, error) {
	if req == nil || len(req.Dcids) == 0 {
		return &InspectIndicatorNodesResult{}, nil
	}

	topicDcids, svDcids := partitionIndicatorDcids(req.Dcids)

	result := &InspectIndicatorNodesResult{}
	if len(topicDcids) > 0 {
		topics, err := s.inspectTopicNodes(ctx, topicDcids)
		if err != nil {
			return nil, err
		}
		result.Topics = topics
	}

	if len(svDcids) > 0 {
		statVars, err := s.inspectStatVarNodes(ctx, svDcids, req.PlaceDcids)
		if err != nil {
			return nil, err
		}
		result.StatVars = statVars
	}

	return result, nil
}

// partitionIndicatorDcids splits input DCIDs into Topic DCIDs and StatisticalVariable DCIDs.
func partitionIndicatorDcids(dcids []string) ([]string, []string) {
	var topics []string
	var statVars []string
	for _, dcid := range dcids {
		trimmed := strings.TrimSpace(dcid)
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, topicDcidPrefix) {
			topics = append(topics, trimmed)
		} else {
			statVars = append(statVars, trimmed)
		}
	}
	return topics, statVars
}

// inspectTopicNodes retrieves direct 1-hop members for Topic DCIDs without recursive expansion bloat.
func (s *Service) inspectTopicNodes(
	ctx context.Context,
	topicDcids []string,
) ([]*TopicInspection, error) {
	resp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    topicDcids,
		Property: exprRelevantVars,
	})
	if err != nil {
		return nil, err
	}

	var out []*TopicInspection
	for _, topicDcid := range topicDcids {
		inspection := &TopicInspection{
			TopicDcid: topicDcid,
		}
		if resp.Data != nil && resp.Data[topicDcid] != nil {
			arcs := resp.Data[topicDcid].Arcs
			if arcs != nil && arcs["relevantVariable"] != nil {
				for _, node := range arcs["relevantVariable"].Nodes {
					if strings.HasPrefix(node.Value, topicDcidPrefix) {
						inspection.ChildTopics = append(inspection.ChildTopics, node.Value)
					} else if node.Value != "" {
						inspection.HeadlineVariables = append(inspection.HeadlineVariables, node.Value)
					}
				}
			}
		}
		out = append(out, inspection)
	}
	return out, nil
}

// inspectStatVarNodes resolves seed metadata, provenances, and ontology breakdown dimensions.
func (s *Service) inspectStatVarNodes(
	ctx context.Context,
	svDcids []string,
	placeDcids []string,
) ([]*StatVarInspection, error) {
	svgResp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    svDcids,
		Property: exprMemberOf,
	})
	if err != nil {
		return nil, err
	}

	bulkResp, err := s.mixer.V2BulkVariableInfo(ctx, &pbv1.BulkVariableInfoRequest{
		Nodes: svDcids,
	})
	if err != nil {
		return nil, err
	}
	summaryBySv := indexVariableSummaries(bulkResp)

	var out []*StatVarInspection
	for _, svDcid := range svDcids {
		targetSvgIds := extractNonRootSvgIds(svgResp, svDcid, s.defaultSvgRoot)
		dims, err := s.fetchDimensionSummariesForSvgs(ctx, targetSvgIds, placeDcids)
		if err != nil {
			return nil, err
		}
		inspection := &StatVarInspection{
			SeedDcid:   svDcid,
			RootSvg:    s.defaultSvgRoot,
			Dimensions: dims,
		}
		populateSeedMetadata(inspection, summaryBySv[svDcid])
		out = append(out, inspection)
	}
	return out, nil
}

// indexVariableSummaries indexes VariableInfoResponse objects by node DCID.
func indexVariableSummaries(resp *pbv1.BulkVariableInfoResponse) map[string]*pbv1.VariableInfoResponse {
	m := make(map[string]*pbv1.VariableInfoResponse)
	if resp == nil {
		return m
	}
	for _, item := range resp.GetData() {
		if item.GetNode() != "" {
			m[item.GetNode()] = item
		}
	}
	return m
}

// populateSeedMetadata attaches provenance names and date boundaries from BulkVariableInfo.
func populateSeedMetadata(inspection *StatVarInspection, infoResp *pbv1.VariableInfoResponse) {
	if infoResp == nil || infoResp.GetInfo() == nil {
		return
	}
	seenProv := make(map[string]bool)
	for _, provSummary := range infoResp.GetInfo().GetProvenanceSummary() {
		provName := provSummary.GetImportName()
		if provName != "" && !seenProv[provName] {
			seenProv[provName] = true
			inspection.Provenances = append(inspection.Provenances, provName)
		}
		for _, series := range provSummary.GetSeriesSummary() {
			if inspection.EarliestDate == "" || (series.GetEarliestDate() != "" && series.GetEarliestDate() < inspection.EarliestDate) {
				inspection.EarliestDate = series.GetEarliestDate()
			}
			if inspection.LatestDate == "" || series.GetLatestDate() > inspection.LatestDate {
				inspection.LatestDate = series.GetLatestDate()
			}
		}
	}
}

// extractNonRootSvgIds extracts StatVarGroup IDs for a StatVar while filtering out the root container.
func extractNonRootSvgIds(resp *pbv2.NodeResponse, svDcid string, rootSvg string) []string {
	if resp == nil || resp.Data == nil || resp.Data[svDcid] == nil {
		return nil
	}
	arcs := resp.Data[svDcid].Arcs
	if arcs == nil || arcs[propMemberOf] == nil {
		return nil
	}
	var svgIds []string
	for _, node := range arcs[propMemberOf].Nodes {
		if node.Value != "" && node.Value != rootSvg {
			svgIds = append(svgIds, node.Value)
		}
	}
	return svgIds
}

// fetchDimensionSummariesForSvgs fetches linked member StatVars and groups them by constraint property.
func (s *Service) fetchDimensionSummariesForSvgs(
	ctx context.Context,
	svgIds []string,
	placeDcids []string,
) ([]*DimensionSliceSummary, error) {
	if len(svgIds) == 0 {
		return nil, nil
	}

	memberResp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    svgIds,
		Property: exprLinkedMembers,
	})
	if err != nil {
		return nil, err
	}

	candidateSvs := collectCandidateSvs(memberResp, svgIds)
	if len(candidateSvs) == 0 {
		return nil, nil
	}

	if len(placeDcids) > 0 {
		candidateSvs, err = s.filterSvsByPlaceAvailability(ctx, candidateSvs, placeDcids)
		if err != nil {
			return nil, err
		}
	}

	return s.groupSvsByConstraintDimension(ctx, candidateSvs)
}

// collectCandidateSvs collects unique StatisticalVariable DCIDs from linkedMemberOf arcs.
func collectCandidateSvs(resp *pbv2.NodeResponse, svgIds []string) []string {
	seen := make(map[string]bool)
	var svs []string
	if resp == nil || resp.Data == nil {
		return nil
	}
	for _, svgId := range svgIds {
		nodeData := resp.Data[svgId]
		if nodeData == nil || nodeData.Arcs == nil || nodeData.Arcs[propLinkedMemberOf] == nil {
			continue
		}
		for _, node := range nodeData.Arcs[propLinkedMemberOf].Nodes {
			if node.Value != "" && !seen[node.Value] {
				seen[node.Value] = true
				svs = append(svs, node.Value)
			}
		}
	}
	return svs
}

// filterSvsByPlaceAvailability filters candidate StatVars to those with observations at any requested placeDcid.
func (s *Service) filterSvsByPlaceAvailability(
	ctx context.Context,
	candidateSvs []string,
	placeDcids []string,
) ([]string, error) {
	obsResp, err := s.mixer.V2Observation(ctx, &pbv2.ObservationRequest{
		Variable: &pbv2.DcidOrExpression{Dcids: candidateSvs},
		Entity:   &pbv2.DcidOrExpression{Dcids: placeDcids},
		Select:   []string{"variable", "entity"},
	})
	if err != nil {
		return nil, err
	}

	var available []string
	if obsResp == nil || obsResp.ByVariable == nil {
		return nil, nil
	}
	for _, sv := range candidateSvs {
		byVar := obsResp.ByVariable[sv]
		if byVar != nil && len(byVar.ByEntity) > 0 {
			available = append(available, sv)
		}
	}
	return available, nil
}

// groupSvsByConstraintDimension groups StatVars by constraint property and formats compact slice maps.
func (s *Service) groupSvsByConstraintDimension(
	ctx context.Context,
	svs []string,
) ([]*DimensionSliceSummary, error) {
	if len(svs) == 0 {
		return nil, nil
	}

	propResp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    svs,
		Property: wildcardPropertyQuery,
	})
	if err != nil {
		return nil, err
	}

	dimMap := make(map[string][]string)
	for _, sv := range svs {
		if propResp.Data == nil || propResp.Data[sv] == nil || propResp.Data[sv].Arcs == nil {
			continue
		}
		arcs := propResp.Data[sv].Arcs
		cProps := arcs["constraintProperties"]
		if cProps == nil {
			continue
		}
		for _, cPropNode := range cProps.Nodes {
			dimName := cPropNode.Value
			if dimName == "" {
				continue
			}
			sliceVal := extractSliceValueForDimension(arcs, dimName, sv)
			dimMap[dimName] = append(dimMap[dimName], fmt.Sprintf("%s -> %s", sliceVal, sv))
		}
	}

	var summaries []*DimensionSliceSummary
	for dim, slices := range dimMap {
		sort.Strings(slices)
		capCount := len(slices)
		if capCount > defaultSampleCap {
			capCount = defaultSampleCap
		}
		summaries = append(summaries, &DimensionSliceSummary{
			Dimension:      dim,
			AvailableCount: len(slices),
			SampleSlices:   slices[:capCount],
		})
	}

	sort.Slice(summaries, func(i, j int) bool {
		if summaries[i].AvailableCount != summaries[j].AvailableCount {
			return summaries[i].AvailableCount > summaries[j].AvailableCount
		}
		return summaries[i].Dimension < summaries[j].Dimension
	})

	return summaries, nil
}

// extractSliceValueForDimension extracts the constraint value for a specific dimension property.
func extractSliceValueForDimension(arcs map[string]*pbv2.Nodes, dimName string, svDcid string) string {
	if arcs[dimName] != nil && len(arcs[dimName].Nodes) > 0 {
		return arcs[dimName].Nodes[0].Value
	}
	return svDcid
}
