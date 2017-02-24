package main

import (
	"github.com/originalgremlin/stream/conf"
	"github.com/originalgremlin/stream/errors"
	"github.com/originalgremlin/stream/reader"
	"github.com/originalgremlin/stream/signals"
	"github.com/originalgremlin/stream/structs"
	"github.com/originalgremlin/stream/structs/wire"
	"github.com/originalgremlin/stream/writer"
	"os"
	"syscall"
)

func main() {
	// setup
	conf := conf.New()
	wire := wire.New()
	errch := make(chan error)

	// readers
	readers := []structs.Reader{
		reader.NewHTTP(),
	}

	// transformers
	transformers := []structs.Transformer{

	}

	// writers
	writers := []structs.Writer{
		writer.NewFileSystem(errch),
	}

	// forward all messages from servers to writers
	for _, r := range readers {
		go r.Start(conf, wire)
	}
	for _, w := range writers {
		go w.Start(conf)
	}
	for message := range wire {
		for _, t := range transformers {
			message = t.Transform(message)
		}
		for _, w := range writers {
			w.Write(message)
		}
	}

	// handle signals
	signals.Handle(syscall.SIGHUP, func(sig os.Signal) error {
		c := make(chan error)
		errs := errors.Errors(len(readers) + len(writers))

		// reload configuration
		for _, r := range readers {
			go func() {
				c <- r.Reload()
			}()
		}
		for _, w := range writers {
			go func() {
				c <- w.Reload()
			}()
		}

		for err := range c {
			errs = errs.Append(err)
		}
		return errs.Error()
	})

	signals.Handle(syscall.SIGTERM, func(sig os.Signal) error {
		c := make(chan error)
		errs := errors.Errors(len(readers) + len(writers))

		// exit gracefully
		for _, r := range readers {
			go r.Shutdown()
		}
		for _, w := range writers {
			go w.Shutdown()
		}

		for err := range c {
			errs = errs.Append(err)
		}
		if !errs.IsNil() {
			// TODO: print error
		}
		// TODO: exit?
		return errs.Error()
	})

	signals.Handle(syscall.SIGINT, func(sig os.Signal) error {
		// exit immediately
		for _, r := range readers {
			go r.Close()
		}
		for _, w := range writers {
			go w.Close()
		}
		return nil
	})

}
