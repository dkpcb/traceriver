package repository

import (
	"context"

	"github.com/dkpcb/pet/domain"
)

// UserRepository defines the persistence interface for User domain objects.
type UserRepository interface {
	// Save persists a new user to the database.
	// Returns an error if the user cannot be saved.
	Save(ctx context.Context, user *domain.User) error

	// FindByID retrieves a user by their ID.
	// Returns nil if the user is not found.
	FindByID(ctx context.Context, id string) (*domain.User, error)

	// FindByLineUserID retrieves a user by their LINE user ID.
	// Returns nil if the user is not found.
	FindByLineUserID(ctx context.Context, lineUserID string) (*domain.User, error)

	// Update updates an existing user.
	// Returns an error if the user cannot be updated.
	Update(ctx context.Context, user *domain.User) error
}
