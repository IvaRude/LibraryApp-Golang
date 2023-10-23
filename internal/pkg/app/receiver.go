//go:generate mockgen -source ./receiver.go -destination=./mocks/receiver/receiver.go -package=mock_receiver
package app

type Receiver interface {
	Subscribe(topic string) error
}
