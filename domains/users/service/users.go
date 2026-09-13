// Package service implements the users service.
package service

import (
	"context"

	"github.com/varunbpatil/go-react-template/domains/users/models"
	"github.com/varunbpatil/go-react-template/domains/users/ports"
)

var _ ports.UserService = (*Service)(nil)

type Service struct {
	repo ports.UserRepository
}

// New creates a users service.
func New(repo ports.UserRepository) (ports.UserService, error) {
	return &Service{repo: repo}, nil
}

// Start implements [ports.UserService].
func (s *Service) Start(ctx context.Context) error {
	return nil
}

// Stop implements [ports.UserService].
func (s *Service) Stop(ctx context.Context) error {
	return nil
}

// GetUser implements [ports.UserService].
func (s *Service) GetUser(ctx context.Context, userID string) (*models.User, error) {
	return s.repo.GetUser(ctx, userID)
}
