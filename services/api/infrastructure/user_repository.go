package infrastructure

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/dkpcb/traceriver/services/api/domain"
	"github.com/dkpcb/traceriver/services/api/infrastructure/table"
	"github.com/dkpcb/traceriver/services/api/repository"
)

// UserRepository is the GORM implementation of repository.UserRepository.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

// Save persists a new user to the database.
func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	row := table.FromDomainUser(user)
	if err := r.db.WithContext(ctx).Create(row).Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

// FindByID retrieves a user by their ID.
func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var row table.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return row.ToDomain(), nil
}

// FindByLineUserID retrieves a user by their LINE user ID.
func (r *UserRepository) FindByLineUserID(ctx context.Context, lineUserID string) (*domain.User, error) {
	var row table.User
	if err := r.db.WithContext(ctx).Where("line_user_id = ?", lineUserID).First(&row).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find user by LINE user ID: %w", err)
	}
	return row.ToDomain(), nil
}

// Update updates an existing user.
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	row := table.FromDomainUser(user)
	if err := r.db.WithContext(ctx).Save(row).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
