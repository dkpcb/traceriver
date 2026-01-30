package domain

// User represents a user entity in the TraceRiver system.
// This is a pure domain model without any infrastructure concerns.
type User struct {
	ID            string
	LineUserID    string
	DisplayName   string
	WalletAddress *string
}

// NewUser creates a new User with required fields.
// walletAddress is optional and can be nil.
func NewUser(id, lineUserID, displayName string, walletAddress *string) *User {
	return &User{
		ID:            id,
		LineUserID:    lineUserID,
		DisplayName:   displayName,
		WalletAddress: walletAddress,
	}
}
