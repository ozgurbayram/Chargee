package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
	"time"
)

func HandleBootNotification(message domain.OcppMessage, cpId string) (*domain.OcppMessage, error) {
	var req domain.BootNotificationRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	response := &domain.BootNotificationResponse{
		Status:      domain.RegistrationStatusAccepted,
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
		Interval:    300,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return &domain.OcppMessage{
		Type:    domain.MessageTypeCallResult,
		Id:      message.Id,
		Action:  message.Action,
		Message: json.RawMessage(responseBytes),
	}, nil
}
