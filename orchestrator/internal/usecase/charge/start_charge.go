package charge

import (
	"context"
	"errors"
)

type StartChargeRequest struct{}

type StartChargeResult struct{}

type StartChargeUseCase interface {
	Execute(ctx context.Context, req StartChargeRequest) (StartChargeResult, error)
}

var ErrNotImplemented = errors.New("not implemented")

type StartChargeService struct{}

func NewStartChargeService() StartChargeUseCase {
	return StartChargeService{}
}

func (StartChargeService) Execute(context.Context, StartChargeRequest) (StartChargeResult, error) {
	return StartChargeResult{}, ErrNotImplemented
}
