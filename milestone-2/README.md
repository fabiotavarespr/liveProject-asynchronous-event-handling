# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
- [Milestone - 2](#milestone---2)
  - [Objective](#objective)
  - [Importance to project](#importance-to-project)
  - [Workflow](#workflow)
  - [Deliverable](#deliverable)

# Milestone - 2
Build a Basic Microservice and Kafka Event Publisher

## Objective

- Create a Go microservice using RESTful best practices.
- Create a publisher that can publish an event to a topic in Kafka.

## Importance to project

The goal of this milestone is to gain experience in creating a basic RESTful microservice and the code necessary to publish an event to a Kafka topic. You will continue to build on this service in the next milestone.

## Workflow

- 1 Create a new Order microservice in Go.
  - Create an HTTP GET endpoint to check the microservice’s health.
  - Test that the service works as expected by running the service and executing a call to the health endpoint. Ensure it responds accordingly.

- 2 Create a function that publishes an event to a topic in Kafka.
  - Test that this function works correctly by creating a main program in Go that will use this function to publish an event to the OrderReceived topic. Verify that it was received by using the appropriate Kafka command-line operation. The easiest way to verify that an event exists in a topic is to use the command illustrated in Step 5 of the “Apache Kafka Quickstart” guide.

## Deliverable

The deliverable for this milestone is a basic RESTful microservice containing a health endpoint and a separate event publisher that can publish an event to the OrderReceived topic in Kafka.