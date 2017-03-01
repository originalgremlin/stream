package structs

import (
	"sync"
)

type Pipeline struct {
	Messages chan Message
	Errors   chan error
	Done     chan bool
}

func NewPipeline() Pipeline {
	return Pipeline{
		make(chan Message),
		make(chan error),
		make(chan bool),
	}
}

func (p *Pipeline) Close() {
	close(p.Messages)
	close(p.Errors)
}

func ConnectPipeline(source func() Pipeline, sinks ...func(Pipeline) Pipeline) Pipeline {
	out := source
	for _, sink := range sinks {
		out = sink(out)
	}
	return out
}

func MergePipeline(ps ...Pipeline) Pipeline {
	var wg sync.WaitGroup
	out := NewPipeline()

	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed, then calls wg.Done.
	wg.Add(2 * len(ps))
	for _, p := range ps {
		go func(p Pipeline) {
			for message := range p.Messages {
				out.Messages <- message
			}
			wg.Done()
		}(p)
		go func(p Pipeline) {
			for err := range p.Errors {
				out.Errors <- err
			}
			wg.Done()
		}(p)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		out.Close()
	}()
	return out

}
