include .env

APP_NAME=inventory-reservation-system-backend

DB_URL=postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}

run:
	@go run ${CURDIR}/cmd/api/main.go serve

dev:
	@air

build:
	@go build -o ${CURDIR}/build/${APP_NAME} ${CURDIR}/cmd/main.go

start:
	@${CURDIR}/build/${APP_NAME} serve

test:
	@go test -v ./...

test-cover:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out
	@rm -f coverage.out

coverage:
	@go test ./... -coverprofile=cover.out
	@go tool cover -html=cover.out
	@rm -f cover.out

mockery:
	@mockery --all --case underscore

migrate-create:
	@migrate create \
	-ext sql \
	-dir ${CURDIR}/db/migration \
	-seq ${name}

migrate-up:
	@migrate \
	-path ${CURDIR}/db/migration \
	-database "${DB_URL}" \
	-verbose up

migrate-down:
	@migrate \
	-path ${CURDIR}/db/migration \
	-database "${DB_URL}" \
	-verbose down

migrate-force:
	@migrate \
	-path ${CURDIR}/db/migration \
	-database "${DB_URL}" \
	force ${version}

migrate-version:
	@migrate \
	-path ${CURDIR}/db/migration \
	-database "${DB_URL}" \
	version

migrate-reset:
	@migrate \
	-path ${CURDIR}/db/migration \
	-database "${DB_URL}" \
	-verbose down -all

	@migrate \
	-path ${CURDIR}/db/migration \
	-database "${DB_URL}" \
	-verbose up

seed:
	@set -e; \
	for file in $(CURDIR)/db/seed/*.sql; do \
		psql "${DB_URL}" -v ON_ERROR_STOP=1 -f "$$file"; \
	done

docker-build:
	@docker build -t ${APP_NAME} .

docker-run:
	@docker run --env-file ${CURDIR}/configs/envs/.env ${APP_NAME}

docker-compose-up:
	@docker compose up --build

docker-compose-down:
	@docker compose down

fmt:
	@gofmt -w .

tidy:
	@go mod tidy

lint:
	@golangci-lint run