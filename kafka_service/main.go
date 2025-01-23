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

	topic := "tasks-log-topic"

	hostName := os.Getenv("KAFKA_HOST")
	if hostName == "" {
		hostName = "localhost:9092"
	}

	worker, err := config.ConnectConsumer([]string{hostName})
	if err != nil {
		log.Fatalf("unable to connect to the topic: %v", err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("error while consuming: %v", err)
	}

	fmt.Println("Server kafka_service is running")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <- consumer.Errors():
				fmt.Println(err)
			case msg := <- consumer.Messages():
				utils.WriteToLog(msg.Value)

			case <- sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<- doneCh
	err = worker.Close(); if err != nil {
		log.Fatalf("error while consuming: %v", err)
	}

}