package table

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/dkpcb/traceriver/services/api/domain"
)

// Interaction is the GORM database model for interactions.
// This is separate from the domain model to maintain clean architecture.
type Interaction struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	RequesterID string    `gorm:"type:char(36);not null;index"`
	ApproverID  string    `gorm:"type:char(36);not null;index"`
	Status      string    `gorm:"type:varchar(20);not null"`
	Metadata    Metadata  `gorm:"type:json"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

// TableName specifies the table name for GORM.
func (Interaction) TableName() string {
	return "interactions"
}

// Metadata is a custom type for JSON handling in GORM.
type Metadata map[string]interface{}

// Scan implements the sql.Scanner interface for GORM.
func (m *Metadata) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*m = result
	return nil
}

// Value implements the driver.Valuer interface for GORM.
func (m Metadata) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// ToDomain converts the database model to a domain model.
func (i *Interaction) ToDomain() *domain.Interaction {
	var metadata map[string]interface{}
	if i.Metadata != nil {
		metadata = i.Metadata
	}

	return domain.NewInteraction(
		i.ID,
		i.RequesterID,
		i.ApproverID,
		domain.InteractionStatus(i.Status),
		metadata,
		i.CreatedAt,
	)
}

// FromDomain creates a database model from a domain model.
func FromDomainInteraction(d *domain.Interaction) *Interaction {
	now := time.Now()
	return &Interaction{
		ID:          d.ID,
		RequesterID: d.RequesterID,
		ApproverID:  d.ApproverID,
		Status:      string(d.Status),
		Metadata:    d.Metadata,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   now,
	}
}
