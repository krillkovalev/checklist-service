package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"kafka_service/models"
)

func WriteToLog(file io.Writer, data []byte) error {
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
