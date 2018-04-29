package spinner

import (
	"fmt"
	"time"
)

type Spinner struct {
	done chan string
}

func Start() *Spinner {
	s := new(Spinner)
	s.done = make(chan string)
	go func() {
		for {
			select {
			case <-s.done:
				return
			default:
			}
			for _, e := range `—\|/` {
				fmt.Printf("\r%c ", e)
				time.Sleep(time.Millisecond * 75)
			}
		}
	}()
	return s
}

func (s *Spinner) Stop() {
	s.done <- ""
}
