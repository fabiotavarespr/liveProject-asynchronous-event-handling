# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
- [Milestone - 1](#milestone---1)
  - [Objective](#objective)
  - [Importance to project](#importance-to-project)
  - [Workflow](#workflow)
- [Deliverable](#deliverable)
  - [Structure](#structure)

# Milestone - 1
Kafka Basics Using the Command Line

## Objective

- Create schemas to represent specific types of events.
- Create and modify Kafka topics.
- Publish events to Kafka topics.
- Consume events from Kafka topics.

## Importance to project

- This milestone introduces working with Kafka topics at their most basic and fundamental level.
- Understanding how to interact with Kafka by using the available command-line tools provided by the platform will enable you to gain deeper insights into Kafka and will aid in testing and troubleshooting the milestones to come.
- Handling situations when a consumer of an event runs into trouble and can’t process the event is an important problem to solve. These events could require further investigation, and using the concept of a dead letter queue is a useful and common pattern used to achieve this goal. In Kafka, the implementation would just be another topic that would be the destination for events that can’t be processed successfully.

## Workflow

- 1 Download and unpack Kafka then set the environment variable KAFKA_HOME to the directory of the unpacked location. The easiest way to start is to use Steps 1 and 2 in the “Apache Kafka Quickstart” guide.

- 2 Define schemas that represent events that will be published and consumed in this e-commerce system. The context of “schema” in this liveProject only refers to a structured, textual representation of the data required to be communicated to the system as an event. Using a JSON document to represent the data is fairly standard. There is no requirement that you install or use the schema registry in Kafka or define an “Avro Schema” for this liveProject. (See the following reference to Kafka in Action, Chapter 11, for more information.)
  -  Define an event schema representing when an order has been received.
  -  Define an event schema representing when an order has been confirmed (is not a duplicate order and can be processed).
  -  Define an event schema representing when an order has been picked from within a warehouse and packed (is ready to be shipped).
  -  Define an event schema representing when an email notification needs to be sent out.
  -  Define an event schema representing when a consumer is unable to successfully process an event they have received. This “error” event should contain the event that could not be processed.

- 3 Create a Kafka topic that will contain order received events, and verify it exists. The easiest way to do this is to use Step 3 in the “Apache Kafka Quickstart” guide.

- 4 Modify the topic created in Step 3 by increasing its retention time to three days. For more information on retention policies, review the last bullet point in the notes below.

- 5 Create additional Kafka topics that use the same configuration as the topic created in Step 3.
  - Create a Kafka topic that will contain order confirmed events, and verify it exists.
  - Create a Kafka topic that will contain order picked and packed events, and verify it exists.
  - Create a Kafka topic that will contain notification events, and verify it exists.
  - Create a Kafka topic that will contain error events, and verify it exists.

- 6 Publish and consume events from the topics created. Create samples of each event schema created in Step 1 and test publishing/consuming them to and from each of their respective topics. The easiest way to do this is to use Steps 4 and 5 in the “Apache Kafka Quickstart” guide.

# Deliverable

The deliverable for this milestone is a set of startup scripts. It is recommended that scripts be created to start Zookeeper and Kafka as well as create all the topics that will be used throughout the project. Having these scripts will ensure consistency and eliminate potential human error that could lead to unnecessary troubleshooting efforts. Your scripts should be executable from the command line.

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
