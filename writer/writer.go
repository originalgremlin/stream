package writer

import (
	"github.com/originalgremlin/stream/configuration"
	"github.com/originalgremlin/stream/structs"
	"sync"
)

type Writers struct {
	writers []structs.Writer
}

func Writers(writers ...structs.Writer) structs.Writer {
	return Writers{writers}
}

func (r *Writers) Write(in structs.Pipeline) structs.Pipeline {
	for _, writer := range r.writers {
		go writer.Write(in)
	}
	return in
}

func (r *Writers) Reload(conf configuration.Configuration) error {
	var wg sync.WaitGroup
	errs := structs.Errors()
	for _, writer := range r.writers {
		wg.Add(1)
		go func(writer structs.Writer) {
			defer wg.Done()
			errs.Append(writer.Reload(conf))
		}(writer)
	}
	wg.Wait()
	return errs.Error()
}

func (r *Writers) Shutdown() error {
	var wg sync.WaitGroup
	errs := structs.Errors()
	for _, writer := range r.writers {
		wg.Add(1)
		go func(writer structs.Writer) {
			defer wg.Done()
			errs.Append(writer.Shutdown())
		}(writer)
	}
	wg.Wait()
	return errs.Error()

}

func (r *Writers) Close() error {
	var wg sync.WaitGroup
	errs := structs.Errors()
	for _, writer := range r.writers {
		wg.Add(1)
		go func(writer structs.Writer) {
			defer wg.Done()
			errs.Append(writer.Close())
		}(writer)
	}
	wg.Wait()
	return errs.Error()
}
