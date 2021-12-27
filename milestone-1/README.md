# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
- [Milestone - 1](#milestone---1)
  - [Objective](#objective)
  - [Structure](#structure)

# Milestone - 1
Kafka Basics Using the Command Line

## Objective

- Create schemas to represent specific types of events.
- Create and modify Kafka topics.
- Publish events to Kafka topics.
- Consume events from Kafka topics.

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
