# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
- [Milestone - 3](#milestone---3)
  - [Objective](#objective)
  - [Importance to project](#importance-to-project)
  - [Workflow](#workflow)
  - [Deliverable](#deliverable)

# Milestone - 3
Build an Order Service and Publish a First Event

## Objective

- Add functionality to the existing Order microservice to accept an order.
- Programmatically publish an event to a topic in Kafka.

## Importance to project

The goal of this milestone is to gain experience creating a RESTful microservice and programmatically publishing an event to a topic in Kafka. This event is the catalyst of execution for all services and components that will be built throughout this project. Completion of this milestone is essential to the success of the next milestone and beyond.

## Workflow

- 1 Create a new Order microservice in Go.
  - Create an HTTP POST endpoint that is responsible for receiving a payload that represents an order.
  - Create a function that will verify that the order payload is valid.
  - Create a function that will translate the order payload into the relevant event schema.
  - Publish an event to the OrderReceived topic in Kafka indicating that an order has been received.
  - Compose the above functions to create an end-to-end solution.
- 2 Test that the service works as expected by posting an order payload to a running service. Verify that it was received by the OrderReceived topic by using the appropriate Kafka command-line operation. The easiest way to verify that an event exists in a topic is to use the command illustrated in Step 5 of the “Apache Kafka Quickstart” guide.

## Deliverable

The deliverable for this milestone is an Order microservice that, when executed, will publish an event to the OrderReceived topic in Kafka.