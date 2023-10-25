package fixtures

import (
	"homework-3/internal/pkg/models"
	"time"
)

type HandlerMessageBuilder struct {
	instance *models.HandlerMessage
}

func NewHandlerMessage() *HandlerMessageBuilder {
	return &HandlerMessageBuilder{instance: &models.HandlerMessage{}}
}

func (b *HandlerMessageBuilder) Timestamp(timestamp time.Time) *HandlerMessageBuilder {
	b.instance.Timestamp = timestamp
	return b
}
func (b *HandlerMessageBuilder) EventType(v string) *HandlerMessageBuilder {
	b.instance.EventType = v
	return b
}

func (b *HandlerMessageBuilder) Body(body string) *HandlerMessageBuilder {
	b.instance.Req.Body = body
	return b
}

func (b *HandlerMessageBuilder) Method(method string) *HandlerMessageBuilder {
	b.instance.Req.Method = method
	return b
}

func (b *HandlerMessageBuilder) P() *models.HandlerMessage {
	return b.instance
}

func (b *HandlerMessageBuilder) V() models.HandlerMessage {
	return *b.instance
}

func (b *HandlerMessageBuilder) Valid() *HandlerMessageBuilder {
	return NewHandlerMessage().Timestamp(time.Now()).EventType("Author").Method("GET").Body("")
}
