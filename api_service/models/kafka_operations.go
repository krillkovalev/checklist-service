package models

import (
	"github.com/IBM/sarama"
	"api_service/config"
)



func PushMessageToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}

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