package receiverhandlers

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/models"

	"github.com/IBM/sarama"
)

type StdoutKafkaReceiverHandler struct {
}

func NewStdoutKafkaReceiverHandler() *StdoutKafkaReceiverHandler {
	return &StdoutKafkaReceiverHandler{}
}

func (skrc *StdoutKafkaReceiverHandler) Handle(message *sarama.ConsumerMessage) error {
	mes := models.HandlerMessage{}
	err := json.Unmarshal(message.Value, &mes)
	if err != nil {
		return err
	}
	fmt.Println("Received Key: ", string(message.Key), " Value: ", mes)
	return nil
}
