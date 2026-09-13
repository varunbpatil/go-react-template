package users

import "github.com/varunbpatil/go-react-template/domains/users/models"

// userToProto converts the user domain type to mcp.
func userToMCP(user *models.User) getUserOutput {
	return getUserOutput{Name: user.Name}
}
