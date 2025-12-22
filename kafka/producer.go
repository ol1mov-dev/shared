package kafka

import "github.com/segmentio/kafka-go"

func ProduceMessage(kafkaTopic string) (*kafka.Writer, error) {
	kafkaConf, err := LoadKafkaConfigs()
	if err != nil {
		return nil, err
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaConf.Address), // адрес вашего Kafka брокера
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{}, // балансировщик для распределения по партициям
	}

	return writer, nil
}
