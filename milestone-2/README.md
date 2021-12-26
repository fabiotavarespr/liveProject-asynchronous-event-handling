# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
- [Milestone - 2](#milestone---2)
  - [Objective](#objective)
  - [Structure](#structure)

# Milestone - 2
Build a Basic Microservice and Kafka Event Publisher

## Objective

- Create a Go microservice using RESTful best practices.
- Create a publisher that can publish an event to a topic in Kafka.

## Structure

```shell
.
├── Makefile
├── README.md
├── docs
│   ├── API_DOCS.md
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── resources
│   ├── examples
│   │   ├── error.json
│   │   ├── notification.json
│   │   ├── order_confirmed.json
│   │   ├── order_picked_and_packed.json
│   │   └── order_received.json
│   └── scripts
│       ├── docker-compose
│       │   ├── Makefile
│       │   ├── README.md
│       │   └── docker-compose.yaml
│       └── shell
│           ├── README.md
│           ├── create_topics.sh
│           ├── start_kafka.sh
│           └── start_zookeeper.sh
└── src
    ├── app
    │   ├── BUSINESS_LOGIC.md
    │   ├── controllers
    │   │   ├── health_controller.go
    │   │   └── order_controller.go
    │   └── models
    │       └── health_model.go
    ├── pkg
    │   ├── PROJECT_SPECIFIC.md
    │   ├── configs
    │   │   └── fiber_config.go
    │   ├── middleware
    │   │   └── fiber_middleware.go
    │   ├── routes
    │   │   ├── health_route.go
    │   │   ├── not_found_route.go
    │   │   ├── public_routes.go
    │   │   └── swagger_route.go
    │   └── utils
    │       └── start_server.go
    └── platform
        └── PLATFORM_LEVEL.md

16 directories, 34 files
```
