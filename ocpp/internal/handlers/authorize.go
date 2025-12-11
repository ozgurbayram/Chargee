package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
)

func HandleAuthorize(message domain.OcppMessage, cpId string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.AuthorizeRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	response := &domain.AuthorizeResponse{
		IdTagInfo: domain.IdTagInfo{
			Status: domain.AuthorizationStatusAccepted,
		},
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return &domain.OcppMessage{
		Type:    domain.MessageTypeCallResult,
		Id:      message.Id,
		Action:  message.Action,
		Message: responseBytes,
	}, nil
}
