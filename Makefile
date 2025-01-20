.PHONY: run api db stop

api:
	cd api_service && go run main.go & echo $$! > api.pid

db:
	cd db_service && go run main.go & echo $$! > db.pid

run: api db
	@echo "Servers are running. Use 'make stop' to terminate them."

stop:
	@if [ -f api.pid ]; then kill $$(cat api.pid) && rm api.pid; echo "Stopped API server"; fi
	@if [ -f db.pid ]; then kill $$(cat db.pid) && rm db.pid; echo "Stopped DB server"; fi
