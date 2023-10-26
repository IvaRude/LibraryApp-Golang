package kafka

import (
	"testing"

	mock_producer "homework-3/internal/infrastructure/kafka/mocks/producer"

	"github.com/golang/mock/gomock"
)

type kafkaSenderFixture struct {
	ctrl     *gomock.Controller
	producer *mock_producer.MockKafkaProducer
	ks       *KafkaSender
}

func setUp(t *testing.T) kafkaSenderFixture {
	ctrl := gomock.NewController(t)
	mp := mock_producer.NewMockKafkaProducer(ctrl)
	return kafkaSenderFixture{
		ctrl:     ctrl,
		producer: mp,
		ks:       NewKafkaSender(mp, "library"),
	}
}

func (ks *kafkaSenderFixture) tearDown() {
	ks.ctrl.Finish()
}
