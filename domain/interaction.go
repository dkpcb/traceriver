package domain

import "time"

// InteractionStatus represents the status of an interaction.
type InteractionStatus string

const (
	InteractionStatusPending  InteractionStatus = "pending"
	InteractionStatusApproved InteractionStatus = "approved"
	InteractionStatusRejected InteractionStatus = "rejected"
)

// Interaction represents an interaction between two users.
// This is a pure domain model without any infrastructure concerns.
type Interaction struct {
	ID          string
	RequesterID string
	ApproverID  string
	Status      InteractionStatus
	Metadata    map[string]interface{}
	CreatedAt   time.Time
}

// NewInteraction creates a new Interaction with required fields.
// Metadata is optional and can be nil.
func NewInteraction(
	id, requesterID, approverID string,
	status InteractionStatus,
	metadata map[string]interface{},
	createdAt time.Time,
) *Interaction {
	return &Interaction{
		ID:          id,
		RequesterID: requesterID,
		ApproverID:  approverID,
		Status:      status,
		Metadata:    metadata,
		CreatedAt:   createdAt,
	}
}

// Approve marks the interaction as approved.
func (i *Interaction) Approve() {
	i.Status = InteractionStatusApproved
}

// Reject marks the interaction as rejected.
func (i *Interaction) Reject() {
	i.Status = InteractionStatusRejected
}

// IsPending returns true if the interaction is in pending status.
func (i *Interaction) IsPending() bool {
	return i.Status == InteractionStatusPending
}
