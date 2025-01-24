package utils

import (
	"fmt"
	"kafka_service/models"
	"encoding/json"
	"os"
)

func WriteToLog(file *os.File, data []byte) error {
	msg := new(models.Messsage)
	err := json.Unmarshal(data, &msg)
    if err != nil {
        return err
    }

	
	str := fmt.Sprintf("INFO: %s %s", msg.Timestamp, msg.Action)
	_, err = fmt.Fprintln(file, str)
	if err != nil {
		return err
	}

	return nil
}