// Package mcp provides the Model Context Protocol inbound adapter.
package mcp

import (
	"net/http"

	gomcp "github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/varunbpatil/go-react-template/version"
)

// New creates a MCP server. Domain packages register their tools with the returned server.
func New() *gomcp.Server {
	return gomcp.NewServer(&gomcp.Implementation{
		Name:        "go-react-template",
		Title:       "Go React Template",
		Description: "",
		Version:     version.Version,
	}, nil)
}

// Handler exposes server through the Streamable HTTP MCP transport.
func Handler(server *gomcp.Server) http.Handler {
	return gomcp.NewStreamableHTTPHandler(func(*http.Request) *gomcp.Server {
		return server
	}, nil)
}
