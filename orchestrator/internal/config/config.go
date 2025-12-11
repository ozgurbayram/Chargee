package config

type Config struct {
	HttpPort  string
	KafkaUrl  string
	KafkaPort string
}

func NewConfig() *Config {
	return &Config{
		HttpPort:  "",
		KafkaUrl:  "",
		KafkaPort: "",
	}
}
