package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
)

var transactionCounter = 0

func HandleStartTransaction(message domain.OcppMessage, cpId string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.StartTransactionRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	transactionCounter++
	transactionId := transactionCounter

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
