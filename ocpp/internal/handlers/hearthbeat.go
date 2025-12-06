package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
	"time"
)

func HandleHeartbeat(message domain.OcppMessage, cpID string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.HeartbeatRequest

	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	chargePoint := &domain.ChargePoint{
		Id:            cpID,
		LastHeartbeat: time.Now().UTC(),
	}

	if err := repository.Upsert(cpID, chargePoint); err != nil {
		return nil, err
	}

	response := &domain.HeartbeatResponse{
		CurrentTime: time.Now().UTC().Format(time.RFC3339),
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
