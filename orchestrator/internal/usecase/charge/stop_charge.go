package charge

import "context"

type StopChargeRequest struct{}

type StopChargeResult struct{}

type StopChargeUseCase interface {
	Execute(ctx context.Context, req StopChargeRequest) (StopChargeResult, error)
}

type StopChargeService struct{}

func NewStopChargeService() StopChargeUseCase {
	return StopChargeService{}
}

func (StopChargeService) Execute(context.Context, StopChargeRequest) (StopChargeResult, error) {
	return StopChargeResult{}, ErrNotImplemented
}
