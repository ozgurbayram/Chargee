package handlers

import (
	"encoding/json"
	"ocpp/internal/domain"
	"strconv"
)

func HandleMeterValues(message domain.OcppMessage, cpId string, repository domain.ChargePointRepository) (*domain.OcppMessage, error) {
	var req domain.MeterValuesRequest
	if err := json.Unmarshal(message.Message, &req); err != nil {
		return nil, err
	}

	for _, mv := range req.MeterValue {
		for _, sv := range mv.SampledValue {
			if sv.Measurand == domain.MeasurandEnergyActiveImportRegister {
				if _, err := strconv.ParseFloat(sv.Value, 64); err != nil {
				}
			}
		}
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
