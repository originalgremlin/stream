package signals

import (
	"fmt"
	"os"
	"os/signal"
)

func Handle(sig os.Signal, handler func(os.Signal) error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig)

	go func() {
		for s := range c {
			fmt.Printf("Handling signal %s\n", s)
			if err := handler(s); err != nil {
				fmt.Printf("Error handling signal %s: %s\n", s, err)
			}
		}
	}()
}
