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
	"fmt"
)

// McpTools holds dependencies for the MCP tools.
type McpTools struct {
	// We don't have any dependencies for the skeletal infra, but we will when we add real tools.
}

// NewMcpTools initializes McpTools.
func NewMcpTools() *McpTools {
	return &McpTools{}
}

// Echo is the business logic for the echo tool.
// This is a temp tool to verify MCP infra and will be removed when we add real tools.
func (m *McpTools) Echo(input EchoInput) (EchoOutput, error) {
	return EchoOutput{
		Response: fmt.Sprintf("Echo: %s", input.Message),
	}, nil
}
