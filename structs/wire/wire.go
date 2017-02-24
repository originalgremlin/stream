package wire

import "github.com/originalgremlin/stream/structs/message"

type Wire chan message.Message

func New() Wire {
	return make(Wire)
}