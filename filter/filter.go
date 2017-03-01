package filter

import (
	"github.com/originalgremlin/stream/structs"
	"sync"
)

type Filter func(*structs.Message) bool

func Filters(filters ...Filter) func(structs.Pipeline) structs.Pipeline {
	return func(in structs.Pipeline) structs.Pipeline {
		out := structs.NewPipeline()
		go func() {
			defer close(out)
			for message := range in.Messages {
				var wg sync.WaitGroup
				passed := true

				// parallel execution of all filters
				for _, filter := range filters {
					wg.Add(1)
					go func(filter Filter) {
						defer wg.Done()
						passed = passed && filter(message)
					}(filter)
				}
				wg.Wait()

				// forward message if all filters passed
				if passed {
					out.Messages <- message
				}
			}
		}()
		return out
	}
}

func True(m *structs.Message) bool {
	return true
}

func False(m *structs.Message) bool {
	return false
}

func Nil(m *structs.Message) bool {
	return m.Contents() != nil
}

func Empty(m *structs.Message) bool {
	contents := m.Contents()
	return contents != nil && len(contents) > 0
}
