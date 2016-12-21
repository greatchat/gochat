package kafka_test

import (
	"testing"
	"time"

	"github.com/Shopify/sarama"
	chat "github.com/greatchat/gochat/transport"
	"github.com/greatchat/gochat/transport/kafka"
)

func TestSendReceiveKafkaMessage(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true

	client, err := kafka.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		t.Fatal(err)
	}

	message := chat.Message{
		Body:      "this is a test",
		Author:    "go devs",
		Timestamp: time.Now(),
	}

	done := make(chan interface{})
	go func() {
		defer close(done)
		received, err := client.Receive("dev")
		if err != nil {
			t.Error(err)
		}

		if received.Body != message.Body {
			t.Errorf("expected Body: %s, got: %s", message.Body, received.Body)
		}

		if received.Author != message.Author {
			t.Errorf("expected Author: %s, got: %s", message.Author, received.Author)
		}

		if received.Timestamp.Unix() != message.Timestamp.Unix() {
			t.Errorf("expected Timestamp: %s, got: %s", message.Timestamp, received.Timestamp)
		}
	}()
	err = client.Send("dev", message)
	if err != nil {
		t.Error(err)
	}

	<-done
}
