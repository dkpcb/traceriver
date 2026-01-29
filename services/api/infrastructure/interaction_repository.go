package infrastructure

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/dkpcb/traceriver/services/api/domain"
	"github.com/dkpcb/traceriver/services/api/infrastructure/table"
	"github.com/dkpcb/traceriver/services/api/repository"
)

// InteractionRepository is the GORM implementation of repository.InteractionRepository.
type InteractionRepository struct {
	db *gorm.DB
}

// NewInteractionRepository creates a new InteractionRepository.
func NewInteractionRepository(db *gorm.DB) repository.InteractionRepository {
	return &InteractionRepository{db: db}
}

// Save persists a new interaction to the database.
func (r *InteractionRepository) Save(ctx context.Context, interaction *domain.Interaction) error {
	row := table.FromDomainInteraction(interaction)
	if err := r.db.WithContext(ctx).Create(row).Error; err != nil {
		return fmt.Errorf("failed to save interaction: %w", err)
	}
	return nil
}

// FindByID retrieves an interaction by its ID.
func (r *InteractionRepository) FindByID(ctx context.Context, id string) (*domain.Interaction, error) {
	var row table.Interaction
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find interaction by ID: %w", err)
	}
	return row.ToDomain(), nil
}

// FindByRequesterID retrieves all interactions requested by a specific user.
func (r *InteractionRepository) FindByRequesterID(ctx context.Context, requesterID string) ([]*domain.Interaction, error) {
	var rows []table.Interaction
	if err := r.db.WithContext(ctx).Where("requester_id = ?", requesterID).Find(&rows).Error; err != nil {
		return nil, fmt.Errorf("failed to find interactions by requester ID: %w", err)
	}

	result := make([]*domain.Interaction, len(rows))
	for i, row := range rows {
		result[i] = row.ToDomain()
	}
	return result, nil
}

// FindByApproverID retrieves all interactions where a specific user is the approver.
func (r *InteractionRepository) FindByApproverID(ctx context.Context, approverID string) ([]*domain.Interaction, error) {
	var rows []table.Interaction
	if err := r.db.WithContext(ctx).Where("approver_id = ?", approverID).Find(&rows).Error; err != nil {
		return nil, fmt.Errorf("failed to find interactions by approver ID: %w", err)
	}

	result := make([]*domain.Interaction, len(rows))
	for i, row := range rows {
		result[i] = row.ToDomain()
	}
	return result, nil
}

// Update updates an existing interaction.
func (r *InteractionRepository) Update(ctx context.Context, interaction *domain.Interaction) error {
	row := table.FromDomainInteraction(interaction)
	if err := r.db.WithContext(ctx).Save(row).Error; err != nil {
		return fmt.Errorf("failed to update interaction: %w", err)
	}
	return nil
}
