package config

type CouchbaseConfig struct {
	ConnectionString string
	Username         string
	Password         string
	BucketName       string
}

type Config struct {
	Port            string
	CouchbaseConfig CouchbaseConfig
}

func NewConfig() *Config {
	return &Config{
		Port: ":8083",
		CouchbaseConfig: CouchbaseConfig{
			ConnectionString: "couchbase://localhost",
			Username:         "admin",
			Password:         "password",
			BucketName:       "ocpp",
		},
	}
}
