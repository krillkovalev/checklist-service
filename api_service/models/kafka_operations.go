package models

import (
	"github.com/IBM/sarama"
	"api_service/config"
	"os"

)

func PushMessageToQueue(topic string, message []byte) error {
	kafkaHost := os.Getenv("KAFKA_HOST")
    if kafkaHost == "" {
        kafkaHost = "localhost:9092"
    }
    brokers := []string{kafkaHost}

	producer, err := config.ConnectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err =  producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}