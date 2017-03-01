package structs

import (
	"github.com/originalgremlin/stream/configuration"
)

type SignalHandler interface {
	// Reload configuration.
	Reload(configuration.Configuration) error

	// Shutdown gracefully shuts down the server without interrupting any active connections.
	Shutdown() error

	// Close immediately closes all active connections and exits the server.
	Close() error
}

type Reader interface {
	Read() Pipeline
	SignalHandler
}

type Writer interface {
	Write(Pipeline) Pipeline
	SignalHandler
}
