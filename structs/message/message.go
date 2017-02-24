package message

import (
	"time"
	"github.com/satori/go.uuid"
)

type Message struct {
	id        uuid.UUID
	timestamp time.Time
	Topic     string
	Contents  []byte
}

func New(topic, contents string) Message {
	return Message{
		uuid.NewV4(),
		time.Now(),
		topic,
		contents,
	}
}

func Parse(bytes []byte) Message {
	len := bytes[0]
	topic := string(bytes[1 : len+1])
	contents := bytes[len+1:]
	return New(topic, contents)
}
