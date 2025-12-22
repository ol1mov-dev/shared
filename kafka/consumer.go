package kafka

import (
	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(topic string, groupId string) (*kafka.Reader, error) {
	kafkaConf, err := LoadKafkaConfigs()

	if err != nil {
		return nil, err
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaConf.Address},
		Topic:   topic,
		GroupID: groupId,
	})

	return reader, nil
}
