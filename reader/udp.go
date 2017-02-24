package reader

import (
	"github.com/originalgremlin/stream/structs"
	"net"
)

type UDP struct{}

func (server *UDP) Serve(conf structs.Configuration, wire structs.Wire) error {
	conn, err := net.ListenPacket("udp", ":"+conf.String("port"))
	if err != nil {
		return err
	}
	for {
		bytes := make([]byte, 65507)
		n, _, err := conn.ReadFrom(bytes)
		if err != nil {
			// TODO: handle error
		}
		wire <- structs.Message(bytes[:n])
	}
	return nil
}
