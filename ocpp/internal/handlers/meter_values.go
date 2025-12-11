package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
)

func HandleMeterValues(message domain.OcppMessage, cpId string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.MeterValuesRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	chargePoint, err := repository.Get(cpId)
	if err != nil {
		chargePoint = &domain.ChargePoint{Id: cpId}
	}

	connectorFound := false
	for i, conn := range chargePoint.Connectors {
		if conn.Id == req.ConnectorId {
			if len(req.MeterValue) > 0 {
				chargePoint.Connectors[i].LastMeterValue = &req.MeterValue[len(req.MeterValue)-1]
			}
			connectorFound = true
			break
		}
	}
	if !connectorFound {
		connector := domain.Connector{Id: req.ConnectorId}
		if len(req.MeterValue) > 0 {
			connector.LastMeterValue = &req.MeterValue[len(req.MeterValue)-1]
		}
		chargePoint.Connectors = append(chargePoint.Connectors, connector)
	}

	if err := repository.Upsert(cpId, chargePoint); err != nil {
		return nil, err
	}

	response := &domain.MeterValuesResponse{}

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
