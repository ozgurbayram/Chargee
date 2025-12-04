package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
	"time"
)

func HandleBootNotification(message domain.OcppMessage) (*domain.BootNotificationResponse, error) {
	var req domain.BootNotificationRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	response := &domain.BootNotificationResponse{
		Status:      domain.RegistrationStatusAccepted,
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
	}

	return response, nil
}
