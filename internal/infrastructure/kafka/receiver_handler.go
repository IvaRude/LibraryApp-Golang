//go:generate mockgen -source ./receiver_handler.go -destination=./mocks/receiver_handler/receiver_handler.go -package=mock_receiver_handler

package kafka

import "github.com/IBM/sarama"

type KafkaReceiverHandler interface {
	Handle(message *sarama.ConsumerMessage) error
}
