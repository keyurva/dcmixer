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

// Package mcp implements the DC MCP server.
package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// DcMcpServer wraps the SDK's MCP server.
type DcMcpServer struct {
	sdkServer *server.MCPServer
}

// NewDcMcpServer initializes the server with basic capabilities.
func NewDcMcpServer() *DcMcpServer {
	s := server.NewMCPServer(
		"Data Commons MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Instantiate tools
	tools := NewMcpTools()

	// Define tool
	echoTool := mcp.NewTool("echo",
		mcp.WithDescription("Echoes back the input message"),
		mcp.WithString("message", mcp.Required(), mcp.Description("The message to echo")),
	)

	// Register tool with handler
	s.AddTool(echoTool, MakeHandler(tools.Echo))

	return &DcMcpServer{sdkServer: s}
}

// MakeHandler is a generic factory that converts a business logic function
// into a handler expected by the mcp sdk.
// It handles JSON decoding of input arguments and JSON encoding of the output.
func MakeHandler[T_IN any, T_OUT any](bizLogic func(T_IN) (T_OUT, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Get arguments as map
		args := req.GetArguments()

		// Convert map to struct using JSON as a bridge
		jsonBytes, err := json.Marshal(args)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}

		var input T_IN
		if err := json.Unmarshal(jsonBytes, &input); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unmarshal arguments: %v", err)), nil
		}

		// Call the business logic
		output, err := bizLogic(input)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		// Return the output struct directly using the SDK's JSON helper
		return mcp.NewToolResultJSON(output)
	}
}
