package writer

import (
	"encoding/gob"
	"os"
	"github.com/originalgremlin/stream/structs/wire"
	"github.com/originalgremlin/stream/conf"
	"github.com/originalgremlin/stream/structs/message"
)

type FileSystem struct {
	In wire.Wire
	Err chan error
}

func NewFileSystem(err chan error) FileSystem {
	return FileSystem{
		wire.New(),
		make(chan error),
	}
}

func (writer FileSystem) Start(conf conf.Configuration) error {
	// TODO: buffer the writes?
	file, err := os.OpenFile(conf.String("path"), os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)

	for message := range writer.In {
		if err := encoder.Encode(message); err != nil {
			writer.Err <- err
		}
	}
	defer close(writer.Err)
	return nil
}

func (writer FileSystem) Write(message message.Message) {
	writer.In <- message
}
