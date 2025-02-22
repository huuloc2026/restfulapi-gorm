build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run:build
	@./bin/ecom


MIGRATIONS_DIR=cmd/migrate/migrations

migration:
	@command -v migrate >/dev/null 2>&1 || { echo "migrate CLI not found! Install it first."; exit 1; }
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) "$(filter-out $@,$(MAKECMDGOALS))"

migrate-up:
	@migrate -path $(MIGRATIONS_DIR) -database $$DATABASE_URL up

migrate-down:
	@migrate -path $(MIGRATIONS_DIR) -database $$DATABASE_URL down

