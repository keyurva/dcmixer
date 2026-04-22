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

// Package mcp defines the tools exposed by the MCP server.
package mcp

import (
	"context"
	"reflect"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

type TestInput struct {
	Message string `json:"message"`
}

type TestOutput struct {
	Response string `json:"response"`
}

type TestTools struct{}

func (t *TestTools) Echo(input TestInput) (TestOutput, error) {
	return TestOutput{Response: "Got: " + input.Message}, nil
}

func TestMakeHandler(t *testing.T) {
	tools := &TestTools{}

	// Create handler using generic factory
	handler := MakeHandler(tools.Echo)

	// Create a mock request
	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "test_tool",
			Arguments: map[string]any{"message": "hello"},
		},
	}

	// Call the handler
	result, err := handler(context.Background(), req)
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if result.IsError {
		t.Fatalf("result is error")
	}

	// Verify StructuredContent
	if result.StructuredContent == nil {
		t.Fatalf("StructuredContent is nil")
	}

	// Verify the full response in StructuredContent
	expectedContent := TestOutput{Response: "Got: hello"}
	if !reflect.DeepEqual(result.StructuredContent, expectedContent) {
		t.Errorf("Expected StructuredContent to be %+v, got %+v", expectedContent, result.StructuredContent)
	}
}
