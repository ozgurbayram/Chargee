package domain

type ChargePointRepository interface {
	Upsert(cpId string, chargePoint *ChargePoint) error
}
