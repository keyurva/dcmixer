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

	pb "github.com/datacommonsorg/mixer/internal/proto"
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

// nodeValueOrDcid returns the DCID if present on the EntityInfo node, falling back to Value.
func nodeValueOrDcid(n *pb.EntityInfo) string {
	if n == nil {
		return ""
	}
	if n.GetDcid() != "" {
		return n.GetDcid()
	}
	return n.GetValue()
}

// InspectIndicatorNodes traverses the indicator ontology and retrieves unified metadata for Topics and StatVars.
func (s *Service) InspectIndicatorNodes(
	ctx context.Context,
	req *pbv2.InspectIndicatorNodesRequest,
) (*pbv2.InspectIndicatorNodesResponse, error) {
	if req == nil || len(req.GetDcids()) == 0 {
		return &pbv2.InspectIndicatorNodesResponse{}, nil
	}

	topicDcids, svDcids := partitionIndicatorDcids(req.GetDcids())

	result := &pbv2.InspectIndicatorNodesResponse{}
	if len(topicDcids) > 0 {
		topics, err := s.inspectTopicNodes(ctx, topicDcids)
		if err != nil {
			return nil, err
		}
		result.Topics = topics
	}

	if len(svDcids) > 0 {
		statVars, err := s.inspectStatVarNodes(ctx, svDcids, req.GetPlaceDcids(), req.GetProperties())
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
) ([]*pbv2.InspectIndicatorNodesResponse_TopicInspection, error) {
	resp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    topicDcids,
		Property: exprRelevantVars,
	})
	if err != nil {
		return nil, err
	}

	var out []*pbv2.InspectIndicatorNodesResponse_TopicInspection
	for _, topicDcid := range topicDcids {
		inspection := &pbv2.InspectIndicatorNodesResponse_TopicInspection{
			TopicDcid: topicDcid,
		}
		if resp.Data != nil && resp.Data[topicDcid] != nil {
			arcs := resp.Data[topicDcid].Arcs
			if arcs != nil && arcs["relevantVariable"] != nil {
				for _, node := range arcs["relevantVariable"].Nodes {
					val := nodeValueOrDcid(node)
					if strings.HasPrefix(val, topicDcidPrefix) {
						inspection.ChildTopics = append(inspection.ChildTopics, val)
					} else if val != "" {
						inspection.HeadlineVariables = append(inspection.HeadlineVariables, val)
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
	properties []string,
) ([]*pbv2.InspectIndicatorNodesResponse_StatVarInspection, error) {
	nodeResp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    svDcids,
		Property: wildcardPropertyQuery,
	})
	if err != nil {
		return nil, err
	}

	checkPlaces := placeDcids
	if len(checkPlaces) == 0 {
		checkPlaces = []string{"country/USA"}
	}
	obsResp, _ := s.mixer.V2Observation(ctx, &pbv2.ObservationRequest{
		Variable: &pbv2.DcidOrExpression{Dcids: svDcids},
		Entity:   &pbv2.DcidOrExpression{Dcids: checkPlaces},
		Select:   []string{"variable", "entity", "facet"},
	})

	bulkResp, _ := s.mixer.V2BulkVariableInfo(ctx, &pbv1.BulkVariableInfoRequest{
		Nodes: svDcids,
	})
	summaryBySv := indexVariableSummaries(bulkResp)

	propSet := make(map[string]bool)
	for _, p := range properties {
		propSet[strings.TrimSpace(p)] = true
	}

	var out []*pbv2.InspectIndicatorNodesResponse_StatVarInspection
	for _, svDcid := range svDcids {
		var dims []*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary
		if len(propSet) > 0 {
			if prefixQuerier, ok := s.mixer.(interface {
				GetStatVarDimensionsByPrefix(ctx context.Context, seedDcid string, properties []string) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error)
			}); ok {
				allDims, _ := prefixQuerier.GetStatVarDimensionsByPrefix(ctx, svDcid, properties)
				for _, d := range allDims {
					if propSet[d.GetDimension()] {
						d.SampleSlices = nil
						dims = append(dims, d)
					}
				}
			}
			if len(dims) == 0 {
				targetSvgIds := extractNonRootSvgIds(nodeResp, svDcid, s.defaultSvgRoot)
				allDims, err := s.fetchDimensionSummariesForSvgs(ctx, targetSvgIds, placeDcids)
				if err != nil {
					return nil, err
				}
				for _, d := range allDims {
					if propSet[d.GetDimension()] {
						d.SampleSlices = nil
						dims = append(dims, d)
					}
				}
			}
		}
		inspection := &pbv2.InspectIndicatorNodesResponse_StatVarInspection{
			SeedDcid:   svDcid,
			RootSvg:    s.defaultSvgRoot,
			Dimensions: dims,
		}
		populateSeedName(inspection, nodeResp, svDcid)
		populateSeedMetadataFromObservations(inspection, svDcid, obsResp)
		if len(inspection.Provenances) == 0 {
			populateSeedMetadata(inspection, summaryBySv[svDcid])
		}
		out = append(out, inspection)
	}
	return out, nil
}

// populateSeedName sets the human-readable name from the StatVar node properties.
func populateSeedName(
	inspection *pbv2.InspectIndicatorNodesResponse_StatVarInspection,
	resp *pbv2.NodeResponse,
	svDcid string,
) {
	if resp == nil || resp.Data == nil || resp.Data[svDcid] == nil || resp.Data[svDcid].Arcs == nil {
		return
	}
	if nameNodes := resp.Data[svDcid].Arcs[propName]; nameNodes != nil && len(nameNodes.Nodes) > 0 {
		inspection.Name = nodeValueOrDcid(nameNodes.Nodes[0])
	}
}

// populateSeedMetadataFromObservations attaches provenance names and date boundaries from V2Observation facets.
func populateSeedMetadataFromObservations(
	inspection *pbv2.InspectIndicatorNodesResponse_StatVarInspection,
	svDcid string,
	obsResp *pbv2.ObservationResponse,
) {
	if obsResp == nil || obsResp.ByVariable == nil || obsResp.ByVariable[svDcid] == nil {
		return
	}
	seenProv := make(map[string]bool)
	for _, byEntity := range obsResp.ByVariable[svDcid].ByEntity {
		for _, orderedFacet := range byEntity.OrderedFacets {
			facet := obsResp.Facets[orderedFacet.FacetId]
			if facet == nil {
				continue
			}
			provName := facet.GetImportName()
			if provName != "" && !seenProv[provName] {
				seenProv[provName] = true
				inspection.Provenances = append(inspection.Provenances, provName)
			}
			if inspection.EarliestDate == "" || (orderedFacet.EarliestDate != "" && orderedFacet.EarliestDate < inspection.EarliestDate) {
				inspection.EarliestDate = orderedFacet.EarliestDate
			}
			if inspection.LatestDate == "" || orderedFacet.LatestDate > inspection.LatestDate {
				inspection.LatestDate = orderedFacet.LatestDate
			}
		}
	}
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
func populateSeedMetadata(inspection *pbv2.InspectIndicatorNodesResponse_StatVarInspection, infoResp *pbv1.VariableInfoResponse) {
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
		val := nodeValueOrDcid(node)
		if val != "" && val != rootSvg {
			svgIds = append(svgIds, val)
		}
	}
	return svgIds
}

// fetchDimensionSummariesForSvgs fetches linked member StatVars across target and child StatVarGroups and groups them by constraint property.
func (s *Service) fetchDimensionSummariesForSvgs(
	ctx context.Context,
	svgIds []string,
	placeDcids []string,
) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error) {
	if len(svgIds) == 0 {
		return nil, nil
	}

	expandedSvgIds := s.expandChildSvgIds(ctx, svgIds)

	memberResp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    expandedSvgIds,
		Property: exprLinkedMembers,
	})
	if err != nil {
		return nil, err
	}

	candidateSvs := collectCandidateSvs(memberResp, expandedSvgIds)
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

// expandChildSvgIds includes direct 1-hop child StatVarGroups (<-specializationOf) so breakdown dimensions in child groups are discovered.
func (s *Service) expandChildSvgIds(ctx context.Context, svgIds []string) []string {
	seen := make(map[string]bool)
	var all []string
	for _, id := range svgIds {
		if id != "" && !seen[id] {
			seen[id] = true
			all = append(all, id)
		}
	}

	childResp, err := s.mixer.V2Node(ctx, &pbv2.NodeRequest{
		Nodes:    svgIds,
		Property: "<-specializationOf",
	})
	if err != nil || childResp == nil || childResp.Data == nil {
		return all
	}

	const maxChildGroups = 40
	for _, svgId := range svgIds {
		nodeData := childResp.Data[svgId]
		if nodeData == nil || nodeData.Arcs == nil || nodeData.Arcs["specializationOf"] == nil {
			continue
		}
		for _, node := range nodeData.Arcs["specializationOf"].Nodes {
			val := nodeValueOrDcid(node)
			if val != "" && !seen[val] {
				seen[val] = true
				all = append(all, val)
				if len(all) >= maxChildGroups {
					return all
				}
			}
		}
	}
	return all
}

// collectCandidateSvs collects unique StatisticalVariable DCIDs from linkedMemberOf arcs, sampling up to maxSvsPerGroup per group.
func collectCandidateSvs(resp *pbv2.NodeResponse, svgIds []string) []string {
	seen := make(map[string]bool)
	var svs []string
	if resp == nil || resp.Data == nil {
		return nil
	}
	const (
		maxSvsPerGroup       = 10
		maxTotalCandidateSvs = 120
	)
	for _, svgId := range svgIds {
		nodeData := resp.Data[svgId]
		if nodeData == nil || nodeData.Arcs == nil || nodeData.Arcs[propLinkedMemberOf] == nil {
			continue
		}
		countForGroup := 0
		for _, node := range nodeData.Arcs[propLinkedMemberOf].Nodes {
			val := nodeValueOrDcid(node)
			if val != "" && !seen[val] {
				seen[val] = true
				svs = append(svs, val)
				countForGroup++
				if countForGroup >= maxSvsPerGroup || len(svs) >= maxTotalCandidateSvs {
					break
				}
			}
		}
		if len(svs) >= maxTotalCandidateSvs {
			break
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
) ([]*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary, error) {
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
			dimName := nodeValueOrDcid(cPropNode)
			if dimName == "" {
				continue
			}
			sliceVal := extractSliceValueForDimension(arcs, dimName, sv)
			dimMap[dimName] = append(dimMap[dimName], fmt.Sprintf("%s -> %s", sliceVal, sv))
		}
	}

	var summaries []*pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary
	for dim, slices := range dimMap {
		sort.Strings(slices)
		capCount := len(slices)
		if capCount > defaultSampleCap {
			capCount = defaultSampleCap
		}
		summaries = append(summaries, &pbv2.InspectIndicatorNodesResponse_DimensionSliceSummary{
			Dimension:      dim,
			AvailableCount: int32(len(slices)),
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
		val := nodeValueOrDcid(arcs[dimName].Nodes[0])
		if val != "" {
			return val
		}
	}
	return svDcid
}
