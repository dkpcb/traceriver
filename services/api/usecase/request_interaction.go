package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/dkpcb/traceriver/services/api/domain"
	"github.com/dkpcb/traceriver/services/api/repository"
)

// RequestInteractionInput represents the input for requesting an interaction.
type RequestInteractionInput struct {
	RequesterLineUserID string
	MessageText         string
}

// RequestInteractionOutput represents the output of requesting an interaction.
type RequestInteractionOutput struct {
	InteractionID string
	ApproverID    string
}

// RequestInteractionUsecase handles the business logic for creating interaction requests.
type RequestInteractionUsecase struct {
	interactionRepo repository.InteractionRepository
	userRepo        repository.UserRepository
	lineService     repository.LineService
}

// NewRequestInteractionUsecase creates a new RequestInteractionUsecase.
func NewRequestInteractionUsecase(
	interactionRepo repository.InteractionRepository,
	userRepo repository.UserRepository,
	lineService repository.LineService,
) *RequestInteractionUsecase {
	return &RequestInteractionUsecase{
		interactionRepo: interactionRepo,
		userRepo:        userRepo,
		lineService:     lineService,
	}
}

// Execute processes an interaction request from a LINE message.
// It parses the message text (expected format: "meet_{UUID}"),
// validates the request, creates the interaction, and sends a notification.
func (u *RequestInteractionUsecase) Execute(ctx context.Context, input *RequestInteractionInput) (*RequestInteractionOutput, error) {
	// 1. Parse the message text to extract approver UUID
	approverUUID, err := u.parseMessageText(input.MessageText)
	if err != nil {
		return nil, fmt.Errorf("invalid message format: %w", err)
	}

	// 2. Get or create the requester user
	requester, err := u.userRepo.FindByLineUserID(ctx, input.RequesterLineUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to find requester: %w", err)
	}
	if requester == nil {
		return nil, fmt.Errorf("requester user not found: %s", input.RequesterLineUserID)
	}

	// 3. Validate the approver exists
	approver, err := u.userRepo.FindByID(ctx, approverUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find approver: %w", err)
	}
	if approver == nil {
		return nil, fmt.Errorf("approver user not found: %s", approverUUID)
	}

	// 4. Validate not requesting to themselves
	if requester.ID == approver.ID {
		return nil, fmt.Errorf("cannot request interaction with yourself")
	}

	// 5. Create the interaction domain model
	interactionID := uuid.New().String()
	interaction := domain.NewInteraction(
		interactionID,
		requester.ID,
		approver.ID,
		domain.InteractionStatusPending,
		nil, // metadata can be added later if needed
		time.Now(),
	)

	// 6. Save the interaction
	if err := u.interactionRepo.Save(ctx, interaction); err != nil {
		return nil, fmt.Errorf("failed to save interaction: %w", err)
	}

	// 7. Send notification to the approver via LINE
	notificationMessage := fmt.Sprintf(
		"%s さんから交流申請が届きました。",
		requester.DisplayName,
	)
	if err := u.lineService.SendMessage(ctx, approver.LineUserID, notificationMessage); err != nil {
		// Log the error but don't fail the entire operation
		// The interaction is already saved
		fmt.Printf("Warning: failed to send LINE notification: %v\n", err)
	}

	return &RequestInteractionOutput{
		InteractionID: interactionID,
		ApproverID:    approver.ID,
	}, nil
}

// parseMessageText extracts the UUID from the message text.
// Expected format: "meet_{UUID}"
func (u *RequestInteractionUsecase) parseMessageText(text string) (string, error) {
	const prefix = "meet_"

	text = strings.TrimSpace(text)
	if !strings.HasPrefix(text, prefix) {
		return "", fmt.Errorf("message must start with '%s'", prefix)
	}

	uuidStr := strings.TrimPrefix(text, prefix)

	// Validate UUID format
	if _, err := uuid.Parse(uuidStr); err != nil {
		return "", fmt.Errorf("invalid UUID format: %w", err)
	}

	return uuidStr, nil
}
