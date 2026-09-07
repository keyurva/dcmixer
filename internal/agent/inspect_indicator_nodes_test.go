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
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
	"github.com/google/go-cmp/cmp"
)

func TestInspectIndicatorNodes(t *testing.T) {
	ctx := context.Background()

	mockMixer := &mockVMMixerServer{
		nodeData: map[string]*pbv2.LinkedGraph{
			"dc/topic/Economy": {
				Arcs: map[string]*pbv2.Nodes{
					"relevantVariable": {
						Nodes: []*pb.EntityInfo{
							{Value: "Count_Person_Employed"},
							{Value: "dc/topic/Employment"},
						},
					},
				},
			},
			"Annual_Emissions_GreenhouseGas_NonBiogenic": {
				Arcs: map[string]*pbv2.Nodes{
					"memberOf": {
						Nodes: []*pb.EntityInfo{
							{Value: "dc/g/CustomRoot"},
							{Value: "dc/g/Emissions_Sector"},
						},
					},
				},
			},
			"dc/g/Emissions_Sector": {
				Arcs: map[string]*pbv2.Nodes{
					"linkedMemberOf": {
						Nodes: []*pb.EntityInfo{
							{Value: "Annual_Emissions_GreenhouseGas_Transportation_NonBiogenic"},
						},
					},
				},
			},
			"Annual_Emissions_GreenhouseGas_Transportation_NonBiogenic": {
				Arcs: map[string]*pbv2.Nodes{
					"constraintProperties": {
						Nodes: []*pb.EntityInfo{
							{Value: "emissionSource"},
						},
					},
					"emissionSource": {
						Nodes: []*pb.EntityInfo{
							{Value: "Transportation"},
						},
					},
				},
			},
		},
		obsData: map[string]*pb.Facet{
			"Annual_Emissions_GreenhouseGas_Transportation_NonBiogenic": {
				ImportName: "TestEmissions",
			},
		},
	}

	service := NewService(mockMixer, NewCache(mockMixer), &ServiceOptions{
		DefaultSvgRoot: "dc/g/CustomRoot",
	})

	req := &InspectIndicatorNodesRequest{
		Dcids:     []string{"dc/topic/Economy", "Annual_Emissions_GreenhouseGas_NonBiogenic"},
		PlaceDcid: "geoId/06",
	}

	got, err := service.InspectIndicatorNodes(ctx, req)
	if err != nil {
		t.Fatalf("InspectIndicatorNodes returned unexpected error: %v", err)
	}

	want := &InspectIndicatorNodesResult{
		Topics: []*TopicInspection{
			{
				TopicDcid:         "dc/topic/Economy",
				HeadlineVariables: []string{"Count_Person_Employed"},
				ChildTopics:       []string{"dc/topic/Employment"},
			},
		},
		StatVars: []*StatVarInspection{
			{
				SeedDcid: "Annual_Emissions_GreenhouseGas_NonBiogenic",
				RootSvg:  "dc/g/CustomRoot",
				Dimensions: []*DimensionSliceSummary{
					{
						Dimension:      "emissionSource",
						AvailableCount: 1,
						SampleSlices: []string{
							"Transportation -> Annual_Emissions_GreenhouseGas_Transportation_NonBiogenic",
						},
					},
				},
			},
		},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("InspectIndicatorNodes mismatch (-want +got):\n%s", diff)
	}
}
