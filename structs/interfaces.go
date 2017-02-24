package structs

import (
	"github.com/originalgremlin/stream/structs/message"
	"github.com/originalgremlin/stream/structs/wire"
	"github.com/originalgremlin/stream/conf"
)

type SignalHandler interface {
	// Reload configuration.
	Reload() error

	// Shutdown gracefully shuts down the server without interrupting any active connections.
	Shutdown() error

	// Close immediately closes all active connections and exits the server.
	Close() error
}

type Reader interface {
	Start(conf.Configuration, wire.Wire) error
	SignalHandler
}

type Transformer interface {
	Transform(message.Message) message.Message
}

type Writer interface {
	Start(conf.Configuration) error
	Write(message.Message)
	SignalHandler
}
