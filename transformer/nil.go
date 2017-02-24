package transformer

import "github.com/originalgremlin/stream/structs/message"

type Nil bool

func (t *Nil) Transform(m message.Message) message.Message {
	m.Contents = nil
	return m
}
