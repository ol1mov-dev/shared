package kafka

import (
	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(kafkaAddress string, topic string, groupId string) (*kafka.Reader, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaAddress}, // адрес вашего Kafka брокера
		Topic:    topic,
		GroupID:  groupId, // ID группы консьюмеров
		MinBytes: 10e3,    // 10KB
		MaxBytes: 10e6,    // 10MB
	})

	return reader, nil
}
