package kafka

import "os"

type KafkaConfig struct {
	Host string
	Port string
}

func LoadKafkaConfigs() (*KafkaConfig, error) {
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")

	return &KafkaConfig{Host: kafkaHost, Port: kafkaPort}, nil
}
