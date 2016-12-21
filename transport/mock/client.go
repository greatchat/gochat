package mock

import chat "github.com/greatchat/gochat/transport"

// Client is a mock implementation chat Client
type Client struct {
	ReceiveFunc  func(src string) (chat.Message, error)
	SendFunc     func(dest string, msg chat.Message) error
	ConsumerFunc func(src string) (chan chat.Message, error)
	ProducerFunc func(dest string) (chan chat.Message, error)
}

// Receive is a mock implementation
func (mock *Client) Receive(src string) (chat.Message, error) {
	return mock.ReceiveFunc(src)
}

// Send is a mock implementation
func (mock *Client) Send(dest string, msg chat.Message) error {
	return mock.SendFunc(dest, msg)
}

// Consumer is a mock implementation
func (mock *Client) Consumer(src string) (chan chat.Message, error) {
	return mock.ConsumerFunc(src)
}

// Producer is a mock implementation
func (mock *Client) Producer(dest string) (chan chat.Message, error) {
	return mock.ProducerFunc(dest)
}
