package repository

import (
	"context"

	"github.com/dkpcb/pet/domain"
)

// InteractionRepository defines the persistence interface for Interaction domain objects.
// Implementations should handle the conversion between domain models and database models.
type InteractionRepository interface {
	// Save persists a new interaction to the database.
	// Returns an error if the interaction cannot be saved.
	Save(ctx context.Context, interaction *domain.Interaction) error

	// FindByID retrieves an interaction by its ID.
	// Returns nil if the interaction is not found.
	FindByID(ctx context.Context, id string) (*domain.Interaction, error)

	// FindByRequesterID retrieves all interactions requested by a specific user.
	FindByRequesterID(ctx context.Context, requesterID string) ([]*domain.Interaction, error)

	// FindByApproverID retrieves all interactions where a specific user is the approver.
	FindByApproverID(ctx context.Context, approverID string) ([]*domain.Interaction, error)

	// Update updates an existing interaction.
	// Returns an error if the interaction cannot be updated.
	Update(ctx context.Context, interaction *domain.Interaction) error
}
