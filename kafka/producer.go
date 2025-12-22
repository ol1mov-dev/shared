package kafka

import "github.com/segmentio/kafka-go"

func ProduceMessage(kafkaAddress string, kafkaTopic string) (*kafka.Writer, error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaAddress), // адрес вашего Kafka брокера
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{}, // балансировщик для распределения по партициям
	}

	return writer, nil
}
