package kafka

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/models"

	"github.com/IBM/sarama"
)

type KafkaSender struct {
	producer KafkaProducer
	topic    string
}

func NewKafkaSender(producer KafkaProducer, topic string) *KafkaSender {
	return &KafkaSender{
		producer,
		topic,
	}
}

func (s *KafkaSender) SendMessage(message *models.HandlerMessage) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return err
	}

	partition, offset, err := s.producer.SendSyncMessage(kafkaMsg)

	if err != nil {
		fmt.Println("Send message connector error", KafkaSendMessageError)
		return KafkaSendMessageError
	}

	fmt.Println("Partition: ", partition, " Offset: ", offset)
	return nil
}

func (s *KafkaSender) buildMessage(message *models.HandlerMessage) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Send message marshal error", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprint(message.Timestamp)),
		Headers: []sarama.RecordHeader{ // например, в хедер можно записать версию релиза
			{
				Key:   []byte("test-header"),
				Value: []byte("test-value"),
			},
		},
	}, nil
}
