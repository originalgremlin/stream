package transformer

import "github.com/originalgremlin/stream/structs/message"

type Identity bool

func (t *Identity) Transform(m message.Message) message.Message {
	return m
}
