//go:generate mockgen -source ./sender.go -destination=./mocks/sender/sender.go -package=mock_sender
package infrastructure

import (
	"homework-3/internal/pkg/models"
)

type Sender interface {
	SendMessage(message *models.HandlerMessage) error
}
