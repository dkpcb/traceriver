package table

import (
	"time"

	"github.com/dkpcb/pet/domain"
)

// User is the GORM database model for users.
// This is separate from the domain model to maintain clean architecture.
type User struct {
	ID            string  `gorm:"type:char(36);primaryKey"`
	LineUserID    string  `gorm:"type:varchar(255);uniqueIndex;not null"`
	DisplayName   string  `gorm:"type:varchar(255);not null"`
	WalletAddress *string `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TableName specifies the table name for GORM.
func (User) TableName() string {
	return "users"
}

// ToDomain converts the database model to a domain model.
func (u *User) ToDomain() *domain.User {
	return domain.NewUser(
		u.ID,
		u.LineUserID,
		u.DisplayName,
		u.WalletAddress,
	)
}

// FromDomainUser creates a database model from a domain model.
func FromDomainUser(d *domain.User) *User {
	return &User{
		ID:            d.ID,
		LineUserID:    d.LineUserID,
		DisplayName:   d.DisplayName,
		WalletAddress: d.WalletAddress,
	}
}
