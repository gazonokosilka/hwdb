include .env
export

service-run:
	go run main.go

postgres-up:
	docker run -e POSTGRES_PASSWORD=pass -p 5432:5432 -v postgres_data:/var/lib/postgresql -d postgres:18-bookworm
migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

postgres-up:
	docker compose up postgres -d

app-up:
	docker compose up application -d