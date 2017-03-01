package statistics

import (
	"fmt"
	"github.com/originalgremlin/stream/structs"
)

type Statistics struct {
	topics map[string]uint64
	errors map[string]uint64
}

func Statistics() Statistics {
	return Statistics{}
}

func (s *Statistics) Topics() map[string]uint64 {
	return s.topics
}

func (s *Statistics) Errors() map[string]uint64 {
	return s.errors
}

func (s *Statistics) Message(message structs.Message) {
	// TODO: stuff with counts per topic and timestamp statistics
	fmt.Println(message)
}

func (s *Statistics) Error(err error) {
	// TODO: stuff with counts per error type and timestamp statistics
	fmt.Println(err)
}
