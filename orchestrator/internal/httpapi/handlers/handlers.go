package handlers

import "orchestrator/internal/usecase/charge"

type Deps struct {
	StartCharge charge.StartChargeUseCase
	StopCharge  charge.StopChargeUseCase
}

type Handler struct {
	startCharge charge.StartChargeUseCase
	stopCharge  charge.StopChargeUseCase
}

func New(deps Deps) *Handler {
	return &Handler{
		startCharge: deps.StartCharge,
		stopCharge:  deps.StopCharge,
	}
}
