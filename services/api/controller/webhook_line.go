package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dkpcb/traceriver/services/api/usecase"
)

// WebhookController handles LINE webhook requests.
type WebhookController struct {
	requestInteractionUsecase *usecase.RequestInteractionUsecase
}

// NewWebhookController creates a new WebhookController.
func NewWebhookController(
	requestInteractionUsecase *usecase.RequestInteractionUsecase,
) *WebhookController {
	return &WebhookController{
		requestInteractionUsecase: requestInteractionUsecase,
	}
}

// LineWebhookRequest represents the LINE webhook request structure.
// This mirrors the OpenAPI schema definition.
type LineWebhookRequest struct {
	Destination string      `json:"destination"`
	Events      []LineEvent `json:"events"`
}

// LineEvent represents a LINE event.
type LineEvent struct {
	Type      string        `json:"type"`
	Timestamp int64         `json:"timestamp"`
	Source    LineSource    `json:"source"`
	Mode      string        `json:"mode"`
	Message   *LineMessage  `json:"message,omitempty"`
	Postback  *LinePostback `json:"postback,omitempty"`
}

// LineSource represents the source of a LINE event.
type LineSource struct {
	Type    string  `json:"type"`
	UserID  string  `json:"userId"`
	GroupID *string `json:"groupId,omitempty"`
	RoomID  *string `json:"roomId,omitempty"`
}

// LineMessage represents a LINE message.
type LineMessage struct {
	ID   string  `json:"id"`
	Type string  `json:"type"`
	Text *string `json:"text,omitempty"`
}

// LinePostback represents LINE postback data.
type LinePostback struct {
	Data string `json:"data"`
}

// PostWebhookLine handles POST /webhook/line requests.
// This implements the operationId: postWebhookLine from the OpenAPI spec.
func (c *WebhookController) PostWebhookLine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse request body
	var req LineWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.sendError(w, http.StatusBadRequest, fmt.Sprintf("invalid request body: %v", err))
		return
	}

	// Process each event
	// In a production system, you might want to process these asynchronously
	for _, event := range req.Events {
		if err := c.handleEvent(ctx, event); err != nil {
			// Log the error but continue processing other events
			fmt.Printf("Error handling event: %v\n", err)
			// Don't return error to LINE as it might retry the same webhook
		}
	}

	// Respond with success
	c.sendSuccess(w)
}

// handleEvent processes a single LINE event.
func (c *WebhookController) handleEvent(ctx context.Context, event LineEvent) error {
	// Only process message events with text
	if event.Type != "message" {
		return nil
	}

	if event.Message == nil || event.Message.Type != "text" || event.Message.Text == nil {
		return nil
	}

	// Execute the request interaction usecase
	input := &usecase.RequestInteractionInput{
		RequesterLineUserID: event.Source.UserID,
		MessageText:         *event.Message.Text,
	}

	_, err := c.requestInteractionUsecase.Execute(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to request interaction: %w", err)
	}

	return nil
}

// sendSuccess sends a successful response.
func (c *WebhookController) sendSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// sendError sends an error response.
func (c *WebhookController) sendError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
