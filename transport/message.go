package transport

import "time"

// Message is the message sent between clients
type Message struct {
	Body      string
	Author    string
	Timestamp time.Time
}
