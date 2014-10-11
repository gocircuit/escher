// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"
	"io"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/shell"
)

// Shell reflexes expose their temporal valve input in the form of an interactive circuit navigation and manipulation REPL.
type Shell struct{
	view chan Circuit
}

func (h *Shell) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	h.view = make(chan Circuit, 1)
	return &Shell{}
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
		for {
			view := <-h.view
			sh.Loop(view)
			eye.Show("Out", v)
		}
	}()
}

// In: Circuit
func (h *Shell) CognizeIn(eye *be.Eye, v interface{}) {
	h.view <- v.(Circuit)
}

// Out: Circuit
func (h *Shell) CognizeOut(*be.Eye, interface{}) {}
