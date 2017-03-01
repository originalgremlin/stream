package main

import (
	"github.com/originalgremlin/stream/configuration"
	"github.com/originalgremlin/stream/filter"
	"github.com/originalgremlin/stream/reader"
	"github.com/originalgremlin/stream/signals"
	"github.com/originalgremlin/stream/structs"
	"github.com/originalgremlin/stream/transform"
	"github.com/originalgremlin/stream/writer"
	"os"
	"syscall"
)

func main() {
	conf := configuration.NewConfiguration()

	readers := reader.Readers(
		reader.NewHTTP(conf),
	)

	transforms := transform.Transforms(
		transform.Identity,
	)

	filters := filter.Filters(
		filter.Empty,
	)

	writers := writer.Writers(
		writer.FileSystem(conf),
	)

	structs.MergePipeline(readers.Read, transforms, filters, writers.Write)

	// handle signals
	// SIGHUP: reload configuration
	signals.Handle(syscall.SIGHUP, func(sig os.Signal) error {
		if ok, err := configuration.Validate(); ok {
			conf := configuration.NewConfiguration()
			errs := structs.Errors()
			errs.Append(readers.Reload(conf), writers.Reload(conf))
			return errs.Error()
		} else {
			return err
		}
	})

	// SIGTERM: exit gracefully
	signals.Handle(syscall.SIGTERM, func(sig os.Signal) error {
		errs := structs.Errors()
		// TODO: enforce a timeout?
		errs.Append(readers.Shutdown(), writers.Shutdown())
		if errs.Error() == nil {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return errs.Error()
	})

	// SIGINT: exit forcefully
	signals.Handle(syscall.SIGINT, func(sig os.Signal) error {
		errs := structs.Errors()
		// TODO: enforce a timeout?
		errs.Append(readers.Close(), writers.Close())
		if errs.Error() == nil {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return errs.Error()
	})
}
