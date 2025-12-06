package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
)

func StatusNotificationHandler(message domain.OcppMessage, cpID string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var request domain.StatusNotificationRequest

	if err := json.Unmarshal(message.Message, &request); err != nil {
		return nil, err
	}

	response := domain.StatusNotificationResponse{}

	var responseBytes, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return &domain.OcppMessage{
		Type:    message.Type,
		Id:      message.Id,
		Action:  message.Action,
		Message: responseBytes,
	}, nil
}
