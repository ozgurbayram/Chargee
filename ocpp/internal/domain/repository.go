package domain

type ChargePointRepository interface {
	Upsert(cpId string, chargePoint *ChargePoint) error
	Get(cpId string) (*ChargePoint, error)
}
