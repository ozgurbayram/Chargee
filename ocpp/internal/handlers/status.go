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

	// Update charge point with connector status
	chargePoint, err := repository.Get(cpID)
	if err != nil {
		// If not found, create new
		chargePoint = &domain.ChargePoint{Id: cpID}
	}

	// Find or add connector
	connectorFound := false
	for i, conn := range chargePoint.Connectors {
		if conn.Id == request.ConnectorId {
			chargePoint.Connectors[i].Status = request.Status
			chargePoint.Connectors[i].ErrorCode = request.ErrorCode
			connectorFound = true
			break
		}
	}
	if !connectorFound {
		chargePoint.Connectors = append(chargePoint.Connectors, domain.Connector{
			Id:        request.ConnectorId,
			Status:    request.Status,
			ErrorCode: request.ErrorCode,
		})
	}

	if err := repository.Upsert(cpID, chargePoint); err != nil {
		return nil, err
	}

	response := domain.StatusNotificationResponse{}

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
