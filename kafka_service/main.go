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
	"github.com/joho/godotenv"
)

func main() {

	topic := "tasks-log-topic"

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	hostName := os.Getenv("KAFKA_HOST")

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