package main

import (
	"fmt"
	"kafka_service/config"
	"kafka_service/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/IBM/sarama"
)

func main() {

	file, err := os.OpenFile("KafkaServiceLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "tasks-log-topic"
	}

	hostName := os.Getenv("KAFKA_HOST")
	if hostName == "" {
		hostName = "localhost:9092"
	}

	worker, err := config.ConnectConsumer([]string{hostName})
	if err != nil {
		log.Fatalf("unable to connect to the topic: %v", err)
	}

	fmt.Println("Server kafka_service is running")


	partitions, err := worker.Partitions(topic)
	if err != nil {
		log.Fatalf("error getting partitions: %v", err)
	}

	for _, partition := range partitions {
		consumer, err := worker.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatalf("error consuming partitions: %v", err)
		}

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		doneCh := make(chan struct{})
		go func(consumer sarama.PartitionConsumer) {
			for {
				select {
				case err := <- consumer.Errors():
					fmt.Println(err)
				case msg := <- consumer.Messages():
					utils.WriteToLog(file, msg.Value)
					if err != nil {
						fmt.Println(err)
					}
				case <- sigchan:
					fmt.Println("Interrupt is detected")
					doneCh <- struct{}{}
				}
			}
		}(consumer)

		<- doneCh
		err = worker.Close(); if err != nil {
			log.Fatalf("error while consuming: %v", err)
		}
	}

}