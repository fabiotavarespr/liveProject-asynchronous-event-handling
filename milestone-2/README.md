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
├── README.md
├── examples
│   ├── error.json
│   ├── notification.json
│   ├── order_confirmed.json
│   ├── order_picked_and_packed.json
│   └── order_received.json
└── scripts
    ├── docker-compose
    │   ├── Makefile
    │   ├── README.md
    │   └── docker-compose.yaml
    └── shell
        ├── README.md
        ├── create_topics.sh
        ├── start_kafka.sh
        └── start_zookeeper.sh

4 directories, 13 files
```
