package kafka

import (
	"homework-3/tests/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	var (
		mes = fixtures.NewHandlerMessage().Valid()
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		kafkaSender := setUp(t)
		defer kafkaSender.tearDown()

		kafkaMes, err := kafkaSender.ks.buildMessage(mes.P())
		require.NoError(t, err)

		kafkaSender.producer.EXPECT().SendSyncMessage(kafkaMes).Return(int32(0), int64(0), nil)
		// act
		err = kafkaSender.ks.SendMessage(mes.P())
		// assert

		assert.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("kafka send message error", func(t *testing.T) {
			t.Parallel()
			// arrange
			kafkaSender := setUp(t)
			defer kafkaSender.tearDown()

			kafkaMes, err := kafkaSender.ks.buildMessage(mes.P())
			require.NoError(t, err)

			kafkaSender.producer.EXPECT().SendSyncMessage(kafkaMes).Return(int32(0), int64(0), KafkaSendMessageError)
			// act
			err = kafkaSender.ks.SendMessage(mes.P())
			// assert

			assert.Error(t, KafkaSendMessageError, err)
		})
	})
}
