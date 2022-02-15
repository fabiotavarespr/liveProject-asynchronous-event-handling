package consumers

// Consumer represents the subscription to a specified Kafka topic
type Consumer struct {
	Broker string
	Group  string
	Topic  string
}
