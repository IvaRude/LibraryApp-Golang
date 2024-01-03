package kafka

import "errors"

var (
	KafkaSendMessageError    = errors.New("Kafka Send Message Error")
	KafkaReceiveMessageError = errors.New("Kafka Receive Message Error")
)
