PROJECT_ROOT = $(shell git rev-parse --show-toplevel)
DOCKER_COMPOSE = docker compose -f $(PROJECT_ROOT)/deploy/docker-compose.yaml
BACKEND_DIR = $(PROJECT_ROOT)/backend
FRONTEND_DIR = $(PROJECT_ROOT)/frontend

.PHONY: help up down clean test

all: help

help:
	@echo "Available Commands:"
	@echo "  make up     Start all services using docker-compose"
	@echo "  make down   Stop and remove all containers"
	@echo "  make clean  Remove unused docker resources"
	@echo "  make test   Run backend tests"

up:
	$(DOCKER_COMPOSE) up --build

down:
	$(DOCKER_COMPOSE) down

clean:
	docker system prune -f

test:
	cd $(BACKEND_DIR) && go test ./...