package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
	"time"
)

var transactionCounter = 0

func HandleStartTransaction(message domain.OcppMessage, cpId string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.StartTransactionRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	transactionCounter++
	transactionId := transactionCounter

	// Get charge point
	chargePoint, err := repository.Get(cpId)
	if err != nil {
		chargePoint = &domain.ChargePoint{Id: cpId}
	}

	// Add transaction
	transaction := domain.Transaction{
		Id:          transactionId,
		ConnectorId: req.ConnectorId,
		IdTag:       req.IdTag,
		StartTime:   time.Now(),
		StartMeter:  req.MeterStart,
	}
	chargePoint.Transactions = append(chargePoint.Transactions, transaction)

	if err := repository.Upsert(cpId, chargePoint); err != nil {
		return nil, err
	}

	response := &domain.StartTransactionResponse{
		IdTagInfo: domain.IdTagInfo{
			Status: domain.AuthorizationStatusAccepted,
		},
		TransactionId: transactionId,
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
