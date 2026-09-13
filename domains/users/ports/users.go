// Package ports defines the users domain ports including the public API for the users domain.
package ports

import (
	"context"

	"github.com/varunbpatil/go-react-template/domains/users/models"
	"github.com/varunbpatil/go-react-template/shared"
)

// UserService is the public API for the users domain.
type UserService interface {
	// Lifecycle hooks to start and stop this service
	shared.StartStopper

	GetUser(ctx context.Context, userID string) (*models.User, error)
}

// UserRepository is interface that users repositories implement.
type UserRepository interface {
	// Lifecycle hooks to close the repository while shutting down
	shared.Closer

	GetUser(ctx context.Context, userID string) (*models.User, error)
}
