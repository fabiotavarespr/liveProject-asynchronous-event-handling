# Live Project - Asynchronous Event Handling Using Microservices and Kafka
[Manning Live Project - Asynchronous Event Handling Using Microservices and Kafka](https://www.manning.com/liveproject/asynchronous-event-handling-using-microservices-and-kafka)

- [Live Project - Asynchronous Event Handling Using Microservices and Kafka](#live-project---asynchronous-event-handling-using-microservices-and-kafka)
  - [About this liveProject](#about-this-liveproject)
  - [Techniques employed](#techniques-employed)
  - [Project outline](#project-outline)
    
## About this liveProject
There are many facets to building an e-commerce system: from the customer shopping experience all the way to the fulfillment and delivery of the products to the customer. The simplicity of the architecture and the speed at which an order can be processed, beginning to end, can have a large impact on the market share and market valuation a company has. Leveraging an asynchronous, event-driven architecture can lead to decoupled and single-focused applications and services that can scale and evolve independently.

PPE 4 Peeps is a small, online, consumer-based retailer of healthcare-related personal protective equipment (PPE). Due to a dramatic rise in demand for PPE, the company is looking to optimize their e-commerce system in order to deliver their products to consumers faster. The current e-commerce system is a monolith that has grown increasingly hard to maintain over the years. It operates in a synchronous manner, which means that each step in the process relies on the completion of a previous step. This leads to unnecessary bottlenecks and waste in the system.

In this liveProject, you are being asked to break apart the order fulfillment part of the monolith into loosely coupled and highly cohesive microservices as well as incorporate asynchronous event handling as the main communication mechanism between them. This approach will allow the system to operate more in parallel and enable each part of the order fulfillment process to evolve and scale independently. You will build microservices in the Go programming language with Kafka as the communication mechanism between them, leveraging its ability to send events between applications, processes, and servers. You will also handle situations in which an event can’t be processed successfully as well as define metrics that will help evaluate the performance of the system that you build.

## Techniques employed
The following are some of the techniques you’ll employ throughout this project. Don’t worry if you haven’t mastered any of these areas—we’ll give you the necessary resources to learn more about each.

Go is a versatile and open-source programming language that was first developed by Google and released in 2012. Since then, the language has grown in popularity and is considered one of the fastest-growing programming languages in the software industry. Go performs well, has fast compilation, and is easy to use.

In this liveProject, we will use Go to do the following:

- Build a microservice.
- Build an event publisher.
- Build multiple event consumers.

Kafka is an increasingly popular, open-source, distributed streaming data platform. It offers three main capabilities: publication and subscription of streaming data, fault-tolerant and durable data storage, and real-time data processing. These capabilities will be key in enabling our system to handle and process events asynchronously. A topic is essentially a named container of data. All data sent to Kafka is stored in topics. Because Kafka durably persists the data it is sent, a consumer can retrieve data as far back as the retention period is configured for the topic. This means that if an application that consumes data from a Kafka topic goes down, when the application comes back up, it could resume processing data where it left off. The opportunity for data loss is significantly low to nonexistent, which is one of the main reasons Kafka was chosen for this project.

In this liveProject, you will use Kafka to do the following:

- Use the command-line tools provided to create, alter, and test topics.
- Publish events to multiple topics.
- Consume events from multiple topics.

## Project outline
The project covers five key concepts: event schema creation, topic creation, publishing an event to a topic programmatically, consuming an event from a topic programmatically, and defining and publishing metrics that will aid in evaluating performance of the system. The learning and implementation of these concepts will take place over six milestones that will each build on its predecessor. These milestones are broadly titled as follows:

1. Kafka Basics Using the Command Line

2. Build a Basic Microservice and Kafka Event Publisher

3. Build an Order Service and Publish a First Event

4. Build an Inventory Consumer That Handles Order Received Events

5. Build Notification, Warehouse, and Shipper Consumers

6. Define and Use KPIs To Evaluate System Performance

By the end of this liveProject you will have created five Kafka topics, a Go microservice, a Go event publisher, and four Go event consumers. These all add up to a system that communicates asynchronously, can process multiple events in parallel, and can be scaled independently. Throughout the project, there will be many opportunities to practice and hone your craft.

