package config

type Config struct {
	Port         string `mapstructure:"port"`
	CouchbaseURL string `mapstructure:"couchbase_url"`
}

func NewConfig() *Config {
	return &Config{
		Port:         ":8083",
		CouchbaseURL: "couchbase://localhost:8091",
	}
}
