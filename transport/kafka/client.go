package kafka

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	chat "github.com/greatchat/gochat/transport"
)

type client struct {
	client sarama.Client
}

// NewClient implements basic client using kafka
func NewClient(addrs []string, conf *sarama.Config) (chat.BasicClient, error) {
	sarClient, err := sarama.NewClient(addrs, conf)
	if err != nil {
		return nil, err
	}

	kafka := client{
		client: sarClient,
	}

	return kafka, nil
}

func (kc client) Send(dest string, msg chat.Message) error {
	producer, err := sarama.NewSyncProducerFromClient(kc.client)
	if err != nil {
		return err
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	pMsg := &sarama.ProducerMessage{
		Topic:     dest,
		Partition: 0,
		Value:     sarama.ByteEncoder(msgBytes),
	}

	_, _, err = producer.SendMessage(pMsg)
	if err != nil {
		return err
	}

	return nil
}

func (kc client) Receive(src string) (chat.Message, error) {
	var message chat.Message

	consumer, err := sarama.NewConsumerFromClient(kc.client)
	if err != nil {
		return message, err
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	pc, err := consumer.ConsumePartition(src, 0, sarama.OffsetNewest)
	if err != nil {
		return message, nil
	}

	defer func() {
		if err := pc.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-pc.Messages()
	err = json.Unmarshal(msg.Value, &message)
	return message, err
}
