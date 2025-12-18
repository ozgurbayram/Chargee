package infra

import (
	"ocpp/internal/config"
	"time"

	"github.com/couchbase/gocb/v2"
)

func CouchbaseInitialization(cfg *config.Config) (*gocb.Cluster, *gocb.Bucket, error) {
	cluster, err := gocb.Connect(cfg.CouchbaseConfig.ConnectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.CouchbaseConfig.Username,
			Password: cfg.CouchbaseConfig.Password,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	bucket := cluster.Bucket(cfg.CouchbaseConfig.BucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, nil, err
	}

	return cluster, bucket, nil
}
