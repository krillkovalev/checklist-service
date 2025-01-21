package utils

import (
	"fmt"
	"kafka_service/models"
	"log"
	"encoding/json"
	"os"
)

func WriteToLog(data []byte) error {
	msg := new(models.Messsage)
	err := json.Unmarshal(data, &msg)
    if err != nil {
        return err
    }

	file, err := os.OpenFile("KafkaServiceLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	str := fmt.Sprintf("INFO: %s %s", msg.Timestamp, msg.Action)
	_, err = fmt.Fprintln(file, str)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}