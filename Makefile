# Makefile para API-Go

run:
	go run main.go

tidy:
	go mod tidy

build:
	go build -o api-go

migrate:
	go run scripts/migrate.go

.PHONY: run tidy build migrate