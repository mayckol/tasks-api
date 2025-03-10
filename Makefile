createmigration:
	@echo "Enter migration name:"
	@read name; \
	migrate create -ext sql -dir internal/infra/database/sql/migrations -seq $$name

migrateup:
	@MYSQL_DATABASE=$(shell grep -E '^MYSQL_DATABASE=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_USER=$(shell grep -E '^MYSQL_USER=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_PASSWORD=$(shell grep -E '^MYSQL_PASSWORD=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_HOST=$(shell grep -E '^MYSQL_HOST=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_PORT=$(shell grep -E '^MYSQL_PORT=' .env | cut -d'=' -f2 | xargs) && \
	DATABASE_URL="mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@tcp($${MYSQL_HOST}:$${MYSQL_PORT})/$${MYSQL_DATABASE}" && \
	migrate -path ./internal/infra/database/sql/migrations -database "$${DATABASE_URL}" -verbose up

migratedown:
	@MYSQL_DATABASE=$(shell grep -E '^MYSQL_DATABASE=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_USER=$(shell grep -E '^MYSQL_USER=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_PASSWORD=$(shell grep -E '^MYSQL_PASSWORD=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_HOST=$(shell grep -E '^MYSQL_HOST=' .env | cut -d'=' -f2 | xargs) && \
	MYSQL_PORT=$(shell grep -E '^MYSQL_PORT=' .env | cut -d'=' -f2 | xargs) && \
	DATABASE_URL="mysql://$${MYSQL_USER}:$${MYSQL_PASSWORD}@tcp($${MYSQL_HOST}:$${MYSQL_PORT})/$${MYSQL_DATABASE}" && \
	migrate -path ./internal/infra/database/sql/migrations -database "$${DATABASE_URL}" -verbose down

tests:
	go test -v ./...

seed:
	go run cmd/seed/main.go

sqlgen:
	@sqlc generate

gendocs:
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/tasks_api/main.go