ifndef VERBOSE
MAKEFLAGS += --no-print-directory
endif

include .env
export

DOCKER_COMPOSE_COMMAND = docker compose

init:
	${DOCKER_COMPOSE_COMMAND} up -d

run:
	go run ./cmd/api/main.go