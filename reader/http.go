package reader

import (
	"net/http"
	"strings"
	"github.com/originalgremlin/stream/conf"
	"github.com/originalgremlin/stream/structs/wire"
	"github.com/originalgremlin/stream/structs/message"
)

type HTTP struct{}

func NewHTTP() HTTP {
	return HTTP{}
}

func (server *HTTP) Serve(c conf.Configuration, w wire.Wire) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			// TODO: error handling
			message := message.Message{
				Topic:    strings.Split(req.URL.Path, "/")[1],
				Contents: make([]byte, req.ContentLength),
			}
			// TODO: error handling
			req.Body.Read(message.Contents)
			w <- message
		}
	})
	return http.ListenAndServe(c.String("port"), nil)
}

func (server *HTTP) Close() error {
	return nil
}

func (server *HTTP) Shutdown() error {
	return nil
}
