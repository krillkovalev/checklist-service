.PHONY: run api db

api:
	cd api_service && go run main.go &

db:
	cd db_service && go run main.go &

run: api db
