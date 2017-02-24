package reader

import (
	"bytes"
	"github.com/originalgremlin/stream/structs"
	"io"
	"net"
)

type TCP struct{}

func (server *TCP) Serve(conf structs.Configuration, wire structs.Wire) error {
	listener, err := net.Listen("tcp", ":"+conf.String("port"))
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO: handle error
			return
		}
		defer conn.Close()

		go func(c net.Conn) {
			var buf bytes.Buffer
			io.Copy(&buf, c)
			wire <- structs.Message(buf)
		}(conn)
	}
	return nil
}
