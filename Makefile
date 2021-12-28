PWD = $(shell pwd -L)
GO_CMD=go
DOCKER_CMD=docker
DOCKER_COMPOSE_CMD=docker-compose
GO_TEST=$(GO_CMD) test
CONTAINER_NAME=kafka
PATH_DOCKER_COMPOSE_FILE=resources/scripts/docker-compose/docker-compose.yaml

.PHONY: all test vendor

all: help

help: ## Display help screen
	@echo "Usage:"
	@echo "	make [COMMAND]"
	@echo "	make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docker-compose-up: ## Run docker-compose services of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) up -d

docker-compose-down: ## Stop docker-compose services of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) down --remove-orphans

docker-compose-restart: docker-compose-down docker-compose-up ## Restart docker-compose services of project

docker-compose-logs: ## Logs docker-compose containers of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) logs -f

docker-compose-ps: ## List docker-compose containers of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) ps

topic-list: ## List Topic
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --list

topic-describe: ## Describe Topic
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --describe

topic-create-OrderReceived: ## Create a Topic with name OrderReceived
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderReceived

topic-create-OrderConfirmed: ## Create a Topic with name OrderConfirmed
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderConfirmed

topic-create-OrderPickedAndPacked: ## Create a Topic with name OrderPickedAndPacked
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderPickedAndPacked

topic-create-Notification: ## Create a Topic with name Notification
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic Notification

topic-create-DeadLetterQueue: ## Create a Topic with name DeadLetterQueue
	docker exec -it $(CONTAINER_NAME) kafka-topics --bootstrap-server localhost:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic DeadLetterQueue

topic-creates: topic-create-OrderReceived topic-create-OrderConfirmed topic-create-OrderPickedAndPacked topic-create-Notification topic-create-DeadLetterQueue ##Create all topics

topic-consumer-OrderReceived: ## Consumer Topic OrderReceived
	docker exec -it $(CONTAINER_NAME) kafka-console-consumer --bootstrap-server localhost:9092 --topic OrderReceived --from-beginning

topic-producer-OrderReceived: ## Consumer Topic OrderReceived
	docker exec -it $(CONTAINER_NAME) kafka-console-producer --bootstrap-server localhost:9092 --topic OrderReceived

topic-consumer-OrderConfirmed: ## Consumer Topic OrderConfirmed
	docker exec -it $(CONTAINER_NAME) kafka-console-consumer --bootstrap-server localhost:9092 --topic OrderConfirmed --from-beginning

topic-producer-OrderConfirmed: ## Consumer Topic OrderConfirmed
	docker exec -it $(CONTAINER_NAME) kafka-console-producer --bootstrap-server localhost:9092 --topic OrderConfirmed

topic-consumer-OrderPickedAndPacked: ## Consumer Topic OrderPickedAndPacked
	docker exec -it $(CONTAINER_NAME) kafka-console-consumer --bootstrap-server localhost:9092 --topic OrderPickedAndPacked --from-beginning

topic-producer-OrderPickedAndPacked: ## Consumer Topic OrderPickedAndPacked
	docker exec -it $(CONTAINER_NAME) kafka-console-producer --bootstrap-server localhost:9092 --topic OrderPickedAndPacked

topic-consumer-Notification: ## Consumer Topic Notification
	docker exec -it $(CONTAINER_NAME) kafka-console-consumer --bootstrap-server localhost:9092 --topic Notification --from-beginning

topic-producer-Notification: ## Consumer Topic Notification
	docker exec -it $(CONTAINER_NAME) kafka-console-producer --bootstrap-server localhost:9092 --topic Notification

topic-consumer-DeadLetterQueue: ## Consumer Topic DeadLetterQueue
	docker exec -it $(CONTAINER_NAME) kafka-console-consumer --bootstrap-server localhost:9092 --topic DeadLetterQueue --from-beginning

topic-producer-DeadLetterQueue: ## Consumer Topic DeadLetterQueue
	docker exec -it $(CONTAINER_NAME) kafka-console-producer --bootstrap-server localhost:9092 --topic DeadLetterQueue

fmt: tidy ## Go Format
	$(GO_CMD) fmt ./...

tidy: ## Go Mod tidy
	$(GO_CMD) mod tidy

run-go: fmt ## Run Project 
	$(GO_CMD) run cmd/api/main.go

download-dependencies: ## Download Dependencies
	go get -d github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest

swag-run: ## Run swag
	~/go/bin/swag init -g cmd/api/main.go