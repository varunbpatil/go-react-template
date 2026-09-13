package users

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"

	"github.com/varunbpatil/go-react-template/domains/users"
	"github.com/varunbpatil/go-react-template/domains/users/ports"
	grpc "github.com/varunbpatil/go-react-template/inbound/grpc"
	usersv1 "github.com/varunbpatil/go-react-template/protos/gen/go_react_template/users/v1"
	"github.com/varunbpatil/go-react-template/protos/gen/go_react_template/users/v1/usersv1connect"
)

var _ usersv1connect.UserServiceHandler = (*Handler)(nil)

// Handler bridges the users domain to the Connect-Go UserServiceHandler interface.
type Handler struct {
	usersv1connect.UnimplementedUserServiceHandler

	svc ports.UserService
}

func New(svc ports.UserService) *Handler {
	return &Handler{svc: svc}
}

// Register registers the user service handler with the configured transport.
func Register(
	server grpc.Registrar,
	handler usersv1connect.UserServiceHandler,
) {
	path, httpHandler := usersv1connect.NewUserServiceHandler(handler, server.HandlerOptions()...)
	server.Handle(path, httpHandler)
}

// GetUser implements [usersv1connect.UserServiceHandler].
func (h *Handler) GetUser(
	ctx context.Context,
	req *connect.Request[usersv1.GetUserRequest],
) (*connect.Response[usersv1.GetUserResponse], error) {
	user, err := h.svc.GetUser(ctx, req.Msg.GetId())
	if errors.Is(err, users.ErrUserNotFound) {
		return nil, notFound(err)
	}
	if err != nil {
		return nil, serviceError(err)
	}
	return connect.NewResponse(&usersv1.GetUserResponse{
		User: userToProto(user),
	}), nil
}

// notFound reports that the requested user does not exist.
func notFound(err error) error {
	return connect.NewError(connect.CodeNotFound, err)
}

// serviceError reports a domain or adapter failure without conflating it with a malformed request.
func serviceError(err error) error {
	return connect.NewError(connect.CodeInternal, fmt.Errorf("user service: %w", err))
}
