package users

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/varunbpatil/go-react-template/domains/users/ports"
)

// Register registers the MCP tools and resources for the users domain.
func Register(server *mcp.Server, svc ports.UserService) {
	h := &handler{svc: svc}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_user",
		Description: "Get a user by ID",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
	}, h.getUser)
}

type handler struct {
	svc ports.UserService
}

type getUserInput struct {
	ID string `json:"id" jsonschema:"The ID of the user to retrieve"`
}

type getUserOutput struct {
	Name string `json:"name" jsonschema:"The user's display name"`
}

func (h *handler) getUser(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input getUserInput,
) (*mcp.CallToolResult, getUserOutput, error) {
	user, err := h.svc.GetUser(ctx, input.ID)
	if err != nil {
		return nil, getUserOutput{}, err
	}
	return nil, userToMCP(user), nil
}
