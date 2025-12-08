.PHONY: build up down logs clean help

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Docker image
	docker-compose build

up: ## Start the application
	docker-compose up -d

down: ## Stop the application
	docker-compose down

logs: ## Show application logs
	docker-compose logs -f

restart: ## Restart the application
	docker-compose restart

clean: ## Remove containers and images
	docker-compose down -v --rmi all

rebuild: ## Rebuild and restart the application
	docker-compose up -d --build

