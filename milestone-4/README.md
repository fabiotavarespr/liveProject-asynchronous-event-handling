# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
- [Milestone - 3](#milestone---3)
  - [Objective](#objective)
  - [Importance to project](#importance-to-project)
  - [Workflow](#workflow)
  - [Deliverable](#deliverable)

# Milestone - 3
Build an Inventory Consumer That Handles Order Received Events

## Objective

- Programmatically consume events from a topic in Kafka while handling duplicates.
- Programmatically publish an error event to a topic in Kafka.

## Importance to project

- The goal of this milestone is to gain experience programmatically subscribing to a topic in Kafka and handling duplicate events in an idempotent manner. This consumer is the first of many consumers that will be built in this project. It will also be the recipient of the first event published in this project (from Milestone 1) and one of three consumers that will receive the event asynchronously. The other two consumers will be built in the next milestone.
- This will also be the first time that errors will be published as events to the DeadLetterQueue topic in Kafka. This pattern is useful when creating a common and consistent process for error handling in the case an event can’t be processed successfully.

## Workflow

- 1 Create a new Inventory consumer in Go.
  - Create a long-lived subscription to the OrderReceived topic in Kafka.
  - Create functionality that extracts (and logs) the order information from the relevant event schema.
    - Treat received events in an idempotent manner, meaning any duplicates that are received must not create any side effects within the system. This can be achieved by two potential solutions: configuring the topic to enable idempotence or handling the duplicate checking in the event consumer by tracking the events processed to detect and discard duplicates.
    - Consider carefully which configuration option the consumer should use to handle Kafka message offset resets. Please refer to the notes below about offsets for more information.
  - Publish an event to the OrderConfirmed Kafka topic when an order has been verified not to be a duplicate.
  - Create functionality that publishes an error event containing the received OrderReceived event to the DeadLetterQueue topic in Kafka when the event can’t be processed successfully. This is the first time you are being asked to publish an error event. You created an error event schema in Milestone 1. If any errors occur when processing the OrderReceived event, create and publish an error event to the DeadLetterQueue topic representing the error received.
- 2 Test that the Inventory consumer works as expected by posting an order payload to a running Order service. Verify that the correct order was received and logged in the Inventory consumer. You should also be able to verify that an event was published to the OrderConfirmed Kafka topic. The easiest way to verify that an event exists in a topic is to use the command illustrated in Step 5 of the “Apache Kafka Quickstart” guide. If any errors occurred while processing the OrderReceived event, you should be able to confirm that an event was published to the DeadLetterQueue Kafka topic.

## Deliverable

The deliverable for this milestone is an Inventory consumer that will be subscribed to the OrderReceived topic in Kafka and can publish events to both the OrderConfirmed and DeadLetterQueue topics in Kafka.