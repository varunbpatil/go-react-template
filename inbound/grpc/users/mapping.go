package users

import (
	"github.com/varunbpatil/go-react-template/domains/users/models"
	usersv1 "github.com/varunbpatil/go-react-template/protos/gen/go_react_template/users/v1"
)

// userToProto converts the user domain type to proto.
func userToProto(user *models.User) *usersv1.User {
	return &usersv1.User{Name: user.Name}
}
