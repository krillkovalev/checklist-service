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

	worker, err := config.ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("unable to connect to the topic: %v", err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("error while consuming: %v", err)
	}

	fmt.Println("Consumer started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	msgCnt := 0
	go func(msgCnt int) {
		for {
			select {
			case err := <- consumer.Errors():
				fmt.Println(err)
			case msg := <- consumer.Messages():
				msgCnt++
				utils.WriteToLog(msg.Value)

			case <- sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}(msgCnt)

	<- doneCh
	fmt.Println("Processed", msgCnt, "messages")

	err = worker.Close(); if err != nil {
		log.Fatalf("error while consuming: %v", err)
	}

}