package infra

import (
	"ocpp/internal/domain"

	"github.com/couchbase/gocb/v2"
)

type CouchbaseChargePointRepository struct {
	bucket *gocb.Bucket
}

func NewCouchbaseChargePointRepository(bucket *gocb.Bucket) domain.ChargePointRepository {
	return &CouchbaseChargePointRepository{
		bucket: bucket,
	}
}

func (r *CouchbaseChargePointRepository) Upsert(cpId string, chargePoint *domain.ChargePoint) error {
	collection := r.bucket.DefaultCollection()
	_, err := collection.Upsert(cpId, chargePoint, nil)
	return err
}

func (r *CouchbaseChargePointRepository) Get(cpId string) (*domain.ChargePoint, error) {
	collection := r.bucket.DefaultCollection()
	result, err := collection.Get(cpId, nil)
	if err != nil {
		return nil, err
	}
	var chargePoint domain.ChargePoint
	err = result.Content(&chargePoint)
	return &chargePoint, err
}
