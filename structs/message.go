package structs

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type Message struct {
	id        uuid.UUID
	timestamp time.Time
	topic     string
	contents  []byte
}

func NewMessage(topic string, contents []byte) Message {
	return Message{
		uuid.NewV4(),
		time.Now(),
		topic,
		contents,
	}
}

func FromBytes(bytes []byte) Message {
	n := bytes[0]
	topic := string(bytes[1 : n+1])
	contents := bytes[n+1:]
	return NewMessage(topic, contents)
}

func (m *Message) Id() uuid.UUID {
	return m.id
}

func (m *Message) Timestamp() time.Time {
	return m.timestamp
}

func (m *Message) Topic() string {
	return m.topic
}

func (m *Message) Contents() []byte {
	return m.contents
}

func (m *Message) String() string {
	return fmt.Sprintf("%s [%s]", m.topic, m.id)
}

func (m *Message) Clear() {
	m.Update([]byte{})
}

func (m *Message) Update(contents []byte) {
	m.contents = contents
}
