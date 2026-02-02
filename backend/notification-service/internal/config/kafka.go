package config

import "os"

type Config struct {
	KafkaBrokers []string
}

func LoadKafkaConfig() *Config {
	kafkaAddr := os.Getenv("KAFKA_BROKERS")
	if kafkaAddr == "" {
		kafkaAddr = "localhost:19092"
	}

	return &Config{
		KafkaBrokers: []string{kafkaAddr},
	}
}
