package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
	"time"
)

func HandleStopTransaction(message domain.OcppMessage, cpId string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.StopTransactionRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	chargePoint, err := repository.Get(cpId)
	if err != nil {
		chargePoint = &domain.ChargePoint{Id: cpId}
	}

	now := time.Now()
	for i, tx := range chargePoint.Transactions {
		if tx.Id == req.TransactionId {
			endTime := now
			chargePoint.Transactions[i].EndTime = &endTime
			chargePoint.Transactions[i].EndMeter = &req.MeterStop
			break
		}
	}

	if err := repository.Upsert(cpId, chargePoint); err != nil {
		return nil, err
	}

	response := &domain.StopTransactionResponse{
		IdTagInfo: &domain.IdTagInfo{
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
