package repository

import "context"

// LineService defines the interface for LINE messaging operations.
// This is placed in the repository package as it's an external service abstraction.
type LineService interface {
	// SendMessage sends a text message to a LINE user.
	// userID is the LINE user ID of the recipient.
	// message is the text content to send.
	// Returns an error if the message cannot be sent.
	SendMessage(ctx context.Context, userID string, message string) error

	// SendFlexMessage sends a Flex Message to a LINE user.
	// userID is the LINE user ID of the recipient.
	// flexMessage is the JSON representation of the Flex Message.
	// Returns an error if the message cannot be sent.
	SendFlexMessage(ctx context.Context, userID string, flexMessage string) error
}
