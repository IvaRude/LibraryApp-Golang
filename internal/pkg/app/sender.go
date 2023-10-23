//go:generate mockgen -source ./sender.go -destination=./mocks/sender/sender.go -package=mock_sender
package app

import (
	"homework-3/internal/pkg/models"
)

type Sender interface {
	SendAsyncMessage(message *models.HandlerMessage) error
	SendMessage(message *models.HandlerMessage) error
	SendMessages(messages []*models.HandlerMessage) error
}

func (a *App) SendMessage(mes *models.HandlerMessage) error {
	return a.HandlerSender.SendMessage(mes)
}
