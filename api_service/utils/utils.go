package utils

import (
	"api_service/config"
	pb "api_service/generated/tasks"
	"api_service/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func LogAction(action string) {
	record := models.Messsage{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Action:    action,
	}

	msg, err := json.Marshal(record)
	if err != nil {
		log.Print("Error marshalling log message: ", err)
		return
	}

	if err := models.PushMessageToQueue("tasks-log-topic", msg); err != nil {
		log.Print("Error pushing message to queue: ", err)
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request, action string, request any, grpcCall func(pb.DBServiceClient, context.Context, any) (any, error)) {
	client, conn, err := config.InitClient()
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer conn.Close()

	if action != "active" && action != "list" {
		if err = json.NewDecoder(r.Body).Decode(request); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}

	res, err := grpcCall(client, r.Context(), request)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	Protores, ok := res.(proto.Message)
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	m := protojson.MarshalOptions{EmitUnpopulated: true}
	responseBody, err := m.Marshal(Protores)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	LogAction(action)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
