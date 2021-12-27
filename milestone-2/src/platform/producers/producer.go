package producers

import (
	"encoding/json"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/src/app/events"
)

func ProducerEvent(event events.Event, topic string) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("BROKER_ADDRESS")})
	if err != nil {
		return err
	}

	defer producer.Close()

	message, _ := json.Marshal(event)

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	// Wait for message deliveries before shutting down
	producer.Flush(600)

	return nil
}
