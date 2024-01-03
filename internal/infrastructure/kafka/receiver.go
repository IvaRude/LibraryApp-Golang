package kafka

import (
	"errors"
	"log"

	"github.com/IBM/sarama"
)

type HandleFunc func(message *sarama.ConsumerMessage)

type KafkaReceiver struct {
	consumer *Consumer
	handlers map[string]KafkaReceiverHandler
}

func NewReceiver(consumer *Consumer, handlers map[string]KafkaReceiverHandler) *KafkaReceiver {
	return &KafkaReceiver{
		consumer: consumer,
		handlers: handlers,
	}
}

func (r *KafkaReceiver) Subscribe(topic string) error {
	receiverHandler, ok := r.handlers[topic]

	if !ok {
		return errors.New("can not find handler")
	}

	// получаем все партиции топика
	partitionList, err := r.consumer.SingleConsumer.Partitions(topic)

	if err != nil {
		return err
	}

	/*
	   sarama.OffsetOldest - перечитываем каждый раз все
	   sarama.OffsetNewest - перечитываем только новые

	   Можем задавать отдельно на каждую партицию
	   Также можем сходить в отдельное хранилище и взять оттуда сохраненный offset
	*/
	initialOffset := sarama.OffsetOldest

	for _, partition := range partitionList {
		pc, err := r.consumer.SingleConsumer.ConsumePartition(topic, partition, initialOffset)

		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			for message := range pc.Messages() {
				err = receiverHandler.Handle(message)
				if err != nil {
					log.Print(err)
				}
			}
		}(pc, partition)
	}

	return nil
}
