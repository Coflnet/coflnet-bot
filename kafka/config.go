package kafka

import (
	"os"
)

func KafkaHost() string {
	return os.Getenv("KAFKA_HOST")
}

func KafkaTopic() string {
	return os.Getenv("KAFKA_CONSUME_TOPIC")
}
