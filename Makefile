all: run lint

run:
	docker-compose up -d 

lint:
	-cd api_service && golangci-lint run ./...
	-cd db_service && golangci-lint run ./...
	-cd kafka_service && golangci-lint run ./...

gen: 
	@protoc \
		--proto_path=protobuf "protobuf/tasks.proto" \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative
