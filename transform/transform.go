package transform

import (
	"fmt"
	"github.com/originalgremlin/stream/structs"
)

type Transform func(*structs.Message) (*structs.Message, error)

func Transforms(transforms ...Transform) func(structs.Pipeline) structs.Pipeline {
	return func(in structs.Pipeline) structs.Pipeline {
		out := structs.NewPipeline()
		go func() {
			defer close(out)
			for message := range in.Messages {
				passed := true

				// serial execution of all transforms
				for _, transform := range transforms {
					if message, err := transform(message); err != nil {
						passed = false
						in.Errors <- fmt.Errorf("Error '%s' transforming message '%s': %s\n", transform, message, err)
						break
					}
				}

				// forward message if it was successfully transformed
				if passed {
					out.Messages <- message
				}
			}
		}()
		return out
	}
}

func Identity(m *structs.Message) (*structs.Message, error) {
	return m, nil
}

func Nil(m *structs.Message) (*structs.Message, error) {
	m.Clear()
	return m, nil
}
