package config

type Config struct {
	HttpAddress string
	HttpPort    string
	KafkaUrl    string
	KafkaPort   string
}

func NewConfig() *Config {
	return &Config{
		HttpAddress: "localhost:",
		HttpPort:    "8084",
		KafkaUrl:    "",
		KafkaPort:   "",
	}
}
