package users

import (
	"context"

	"github.com/varunbpatil/go-react-template/domains/users/models"
	"github.com/varunbpatil/go-react-template/domains/users/ports"
)

var _ ports.UserRepository = (*Repository)(nil)

type Repository struct{}

// New creates a postgres users repository.
func New() (*Repository, error) {
	return &Repository{}, nil
}

// Close implements [ports.UserRepository].
func (r *Repository) Close() error {
	return nil
}

// GetUser implements [ports.UserRepository].
func (r *Repository) GetUser(ctx context.Context, userID string) (*models.User, error) {
	return &models.User{}, nil
}
