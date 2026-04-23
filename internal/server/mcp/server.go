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
	"embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// DcMcpServer wraps the SDK's MCP server.
type DcMcpServer struct {
	sdkServer  *server.MCPServer
	httpServer *server.StreamableHTTPServer
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

	// Create streamable HTTP server
	httpSrv := server.NewStreamableHTTPServer(s)

	dcMcpSrv := &DcMcpServer{sdkServer: s, httpServer: httpSrv}

	// Register the echo tool using the package-level helper
	registerTool(dcMcpSrv, "echo", tools.Echo)

	return dcMcpSrv
}

// ServeHTTP satisfies the http.Handler interface.
func (s *DcMcpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.httpServer.ServeHTTP(w, r)
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

		// Log tool invocation
		slog.Info("Invoking MCP tool", "name", req.Params.Name, "arguments", args)

		// Call the business logic
		output, err := bizLogic(input)
		if err != nil {
			slog.Error("MCP tool execution failed", "name", req.Params.Name, "error", err)
			return mcp.NewToolResultError(err.Error()), nil
		}

		slog.Info("MCP tool execution completed", "name", req.Params.Name)

		// Return the output struct directly using the SDK's JSON helper
		return mcp.NewToolResultJSON(output)
	}
}

// TODO(keyurs): Support loading instructions from a custom path (e.g., GCS).
//
//go:embed instructions/*
var instructionsFS embed.FS

func loadDescription(toolName string) string {
	path := fmt.Sprintf("instructions/%s.md", toolName)
	content, err := instructionsFS.ReadFile(path)
	if err != nil {
		slog.Warn("Failed to load tool description from embedded files", "tool", toolName, "error", err)
		return "Fallback description for " + toolName
	}
	return string(content)
}

// registerTool is a generic helper to register a tool with its input schema and handler.
func registerTool[T_IN any, T_OUT any](s *DcMcpServer, name string, bizLogic func(T_IN) (T_OUT, error)) {
	// Load description from file
	description := loadDescription(name)

	// Define tool using struct schema
	tool := mcp.NewTool(name,
		mcp.WithDescription(description),
		mcp.WithInputSchema[T_IN](),
	)

	// Register tool with handler created by the factory
	s.sdkServer.AddTool(tool, MakeHandler(bizLogic))
}

// StartKeepalives starts a background goroutine that broadcasts dummy JSON-RPC
// notifications to all connected clients at the specified interval.
// These keep alives help keep the connection alive through proxies, 
// thereby avoiding the benign but annoying "MCP Error" messages in MCP clients.
func (s *DcMcpServer) StartKeepalives(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.sdkServer.SendNotificationToAllClients(
					"notifications/message",
					map[string]any{"level": "debug", "message": "keepalive"},
				)
			}
		}
	}()
}
