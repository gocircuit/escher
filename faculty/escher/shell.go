// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"io"
	"log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/shell"
)

// Shell reflexes expose their temporal valve input in the form of an interactive circuit navigation and manipulation REPL.
type Shell struct{
	view chan interface{}
	shell *shell.Shell
}

func (h *Shell) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	h.view = make(chan interface{}, 1)
	return nil
}

// In: { Name string, In io.Reader, Out io.WriteCloser, Err io.WriteCloser }
func (h *Shell) CognizeUser(eye *be.Eye, v interface{}) {
	go func() {
		x := v.(Circuit)
		sh := shell.NewShell(
			x.StringAt("Name"),
			x.At("In").(io.Reader),
			x.At("Out").(io.WriteCloser),
			x.At("Err").(io.WriteCloser),
		)
		u := <-h.view
		if _, ok := u.(Circuit); !ok {
			log.Fatalf("Shell gate received non-circuit value: %v", u)
		}
		sh.Start(u.(Circuit)) // shell attaches to first value on default valve
	}()
}

func (h *Shell) Cognize(eye *be.Eye, v interface{}) {
	h.view <- v
}
