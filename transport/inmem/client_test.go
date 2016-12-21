package inmem_test

import (
	"testing"
	"time"

	chat "github.com/greatchat/gochat/transport"
	"github.com/greatchat/gochat/transport/inmem"
)

func TestSendReceiveInMemMessage(t *testing.T) {
	client := inmem.NewClient(100)
	message := chat.Message{
		Body:      "this is a test",
		Author:    "go devs",
		Timestamp: time.Now(),
	}

	t.Run("TestSuccess", func(t *testing.T) {
		err := client.Send("dev", message)
		if err != nil {
			t.Error(err)
		}

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

		if received.Timestamp != message.Timestamp {
			t.Errorf("expected Timestamp: %s, got: %s", message.Timestamp, received.Timestamp)
		}
	})
}
