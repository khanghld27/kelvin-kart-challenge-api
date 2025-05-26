gen:
	@go generate -v ./...

env:
	@export $(grep -v '^#' ./.env | xargs)
mod:
	@go mod tidy && go mod vendor

install-migrate-tool:
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migration:
	@migrate create -ext sql -dir databases/postgres/migrations -seq $(name)
	@echo "Migration file created in databases/postgres/migrations"

migrateup:
	@migrate -database ${DB_URI} -path databases/postgres/migrations up 1

migratedown:
	@migrate -source databases/postgres/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}" down 1

run:
	@go run ./cmd/services/core/...

docker-up:
	@docker-compose -f ./deployments/local/docker-compose.yml build --no-cache
	@docker-compose -f ./deployments/local/docker-compose.yml up -d

docker-down:
	@docker-compose -f ./deployments/local/docker-compose.yml down