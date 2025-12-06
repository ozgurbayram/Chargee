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
