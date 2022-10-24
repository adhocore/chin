package chin

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

var chars = []string{"+", "\\", "|", "!", "/", "-", "x", "+", "\\", "|", "!", "/", "-", "x"}

// Chin is the spinner struct
type Chin struct {
	stch chan bool
	stop bool
	wg   *sync.WaitGroup
}

// New gets a new spinner
func New() *Chin {
	return &Chin{}
}

// WithWait attaches a wait group
func (s *Chin) WithWait(wg *sync.WaitGroup) *Chin {
	wg.Add(1)
	s.wg = wg
	return s
}

// Start starts the spinner
func (s *Chin) Start() {
	tput("civis")
	s.doSpin()
}

// Stop stops the spinner
func (s *Chin) Stop() {
	if s.wg != nil {
		defer s.wg.Done()
	}
	s.stop = true
	tput("cvvis")
}

func (s *Chin) doSpin() {
	for !s.stop {
	outer:
		select {
		case _, ok := <-s.stch:
			if ok {
				fmt.Print("\010")
				break outer
			}
		default:
			for _, c := range chars {
				if s.stop {
					s.stch <- true
				}
				fmt.Print(c, "\010")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}

// https://rosettacode.org/wiki/Terminal_control/Hiding_the_cursor
func tput(arg string) error {
	cmd := exec.Command("tput", arg)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
