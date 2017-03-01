package reader

import (
	"github.com/originalgremlin/stream/configuration"
	"github.com/originalgremlin/stream/structs"
	"net/http"
	"strings"
)

type HTTP struct {
	conf     configuration.Configuration
	pipeline structs.Pipeline
}

func NewHTTP(conf configuration.Configuration) HTTP {
	return HTTP{
		conf,
		structs.NewPipeline(),
	}
}

func (r *HTTP) Read() structs.Pipeline {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			switch req.Method {
			case "POST":
				// TODO: error handling
				topic := strings.Split(req.URL.Path, "/")[1]
				contents := make([]byte, req.ContentLength)
				req.Body.Read(contents)
				r.pipeline.Messages <- structs.NewMessage(topic, contents)
			}
		})
		r.pipeline.Errors <- http.ListenAndServe(r.conf.String("port"), nil)
	}()
	return r.pipeline
}

func (r *HTTP) Reload() error {
	return nil
}

func (r *HTTP) Shutdown() error {
	return nil
}

func (r *HTTP) Close() error {
	return nil
}
