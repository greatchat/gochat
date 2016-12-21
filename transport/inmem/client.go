package inmem

import (
	chat "github.com/greatchat/gochat/transport"
)

type client struct {
	chanBuff int
	channels map[string]chan chat.Message
}

// NewClient initalises new in memory client
func NewClient(chanBuffer int) chat.BasicClient {
	return &client{
		chanBuff: chanBuffer,
		channels: make(map[string]chan chat.Message),
	}
}

func (c *client) channel(name string) chan chat.Message {
	ch, found := c.channels[name]
	if !found {
		ch = make(chan chat.Message, c.chanBuff)
		c.channels[name] = ch
	}
	return ch
}

// Send message to dest channel
func (c *client) Send(dest string, msg chat.Message) error {
	// TODO timeout
	c.channel(dest) <- msg
	return nil
}

// Receive gets message from src chan
func (c *client) Receive(src string) (chat.Message, error) {
	// TODO timeout
	msg := <-c.channel(src)
	return msg, nil
}
