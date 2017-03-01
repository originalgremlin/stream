package reader

import (
	"github.com/originalgremlin/stream/configuration"
	"github.com/originalgremlin/stream/structs"
	"sync"
)

type Readers struct {
	readers  []structs.Reader
	pipeline structs.Pipeline
}

func Readers(readers ...structs.Reader) structs.Reader {
	return Readers{
		readers,
		structs.NewPipeline(),
	}
}

func (r *Readers) Read() structs.Pipeline {
	pipelines := make([]structs.Pipeline, len(r.readers))
	var wg sync.WaitGroup
	for i, reader := range r.readers {
		wg.Add(1)
		go func(reader structs.Reader) {
			defer wg.Done()
			pipelines[i] = reader.Read()
		}(reader)
	}
	wg.Wait()
	r.pipeline = structs.MergePipeline(pipelines...)
	return r.pipeline
}

func (r *Readers) Reload(conf configuration.Configuration) error {
	var wg sync.WaitGroup
	errs := structs.Errors()
	for _, reader := range r.readers {
		wg.Add(1)
		go func(reader structs.Reader) {
			defer wg.Done()
			errs.Append(reader.Reload(conf))
		}(reader)
	}
	wg.Wait()
	return errs.Error()
}

func (r *Readers) Shutdown() error {
	var wg sync.WaitGroup
	errs := structs.Errors()
	for _, reader := range r.readers {
		wg.Add(1)
		go func(reader structs.Reader) {
			defer wg.Done()
			errs.Append(reader.Shutdown())
		}(reader)
	}
	wg.Wait()
	return errs.Error()

}

func (r *Readers) Close() error {
	var wg sync.WaitGroup
	errs := structs.Errors()
	for _, reader := range r.readers {
		wg.Add(1)
		go func(reader structs.Reader) {
			defer wg.Done()
			errs.Append(reader.Close())
		}(reader)
	}
	wg.Wait()
	return errs.Error()
}
