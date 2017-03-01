package writer

import (
	"encoding/gob"
	"github.com/originalgremlin/stream/structs"
	"os"
)

type FileSystem struct {
	conf structs.Configuration
}

func FileSystem(conf structs.Configuration) FileSystem {
	return FileSystem{conf}
}

func (writer FileSystem) Write(messages chan structs.Message, errors chan error) {
	// TODO: buffer the writes?
	file, err := os.OpenFile(writer.conf.String("path"), os.O_APPEND, os.ModeAppend)
	if err != nil {
		errors <- err
		return messages, errors
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)

	go func() {
		for message := range messages {
			if err := encoder.Encode(message); err != nil {
				errors <- err
			} else {
				messages <- message
			}
		}
	}()
}
