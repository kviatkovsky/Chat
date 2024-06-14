ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

run:
	go run cmd/app/main.go --config=./config/local.yaml

migrate-db:
	docker run -v /Users/vitaliykviatkovsky/Developer/golang/ChatWS/server/db/migrations:/migrations --network host migrate/migrate -path=/migrations -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable up

docker-rebuild:
	docker-compose build && docker-compose up

up: migrate-db run

test:
	go test ./internal/grpc/...

.PHONY: run generate migrate-db docker-rebuildinitial test