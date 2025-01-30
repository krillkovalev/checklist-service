PROTO_DIR=protobuf
OUT_DIR=./generated/tasks

all: run lint

run:
	docker-compose up -d 

lint:
	-cd api_service && golangci-lint run ./...
	-cd db_service && golangci-lint run ./...
	-cd kafka_service && golangci-lint run ./...

gen:
	protoc \
	--go_out=$(OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(OUT_DIR) \
	--proto_path=$(PROTO_DIR) tasks.proto \
	--go-grpc_opt=paths=source_relative 
                
