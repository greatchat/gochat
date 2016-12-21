package transport

import "log"

// Client defines consumer functionality
type Client interface {
	BasicClient
	Consumer(src string) (chan Message, error)
	Producer(dest string) (chan Message, error)
}

// BasicClient just sends and receives
type BasicClient interface {
	Receive(src string) (Message, error)
	Send(dest string, msg Message) error
}

type basicClientWrapper struct {
	chanBuff int
	BasicClient
}

// BasicClientWrapper wraps a basic client to meet the Client interface
func BasicClientWrapper(client BasicClient, chanBuff int) Client {
	return basicClientWrapper{
		chanBuff,
		client,
	}
}

// Consumer inplements consumer interface using basicClient
func (client basicClientWrapper) Consumer(src string) (chan Message, error) {
	ch := make(chan Message, client.chanBuff)
	go func() {
		for {
			// TODO only whil ch not close
			mes, err := client.Receive(src)
			if err != nil {
				// TODO: should surface
				log.Println(err)
				continue
			}
			ch <- mes
		}
	}()
	return ch, nil
}

// Producer inplements producer interface using basicClient
func (client basicClientWrapper) Producer(dest string) (chan Message, error) {
	ch := make(chan Message, client.chanBuff)
	go func() {
		for msg := range ch {
			err := client.Send(dest, msg)
			if err != nil {
				// TODO: should surface
				log.Println(err)
			}
		}
	}()
	return ch, nil
}
