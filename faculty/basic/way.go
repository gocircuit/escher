// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("J3", Junction3{})
}

// Junction3
type Junction3 struct{}

func (Junction3) Materialize() think.Reflex {
	a0Endo, a0Exo := think.NewSynapse()
	a1Endo, a1Exo := think.NewSynapse()
	a2Endo, a2Exo := think.NewSynapse()
	go func() {
		h := &junction3{
			ready: make(chan struct{}),
		}
		h.re[0] = a0Endo.Focus(func(v interface{}) { h.Cognize(0, v) })
		h.re[1] = a1Endo.Focus(func(v interface{}) { h.Cognize(1, v) })
		h.re[2] = a2Endo.Focus(func(v interface{}) { h.Cognize(2, v) })
		close(h.ready)
	}()
	return think.Reflex{
		"X": a0Exo, 
		"Y": a1Exo, 
		"Z": a2Exo,
	}
}

type junction3 struct {
	ready chan struct{}
	re [3]*think.ReCognizer
}

func (h *junction3) Cognize(way int, v interface{}) {
	<-h.ready
	ch := make(chan struct{})
	for i, re := range h.re {
		if i == way {
			continue
		}
		go func() {
			re.ReCognize(v)
		}()
	}
	for i, _ := range h.re {
		if i == way {
			continue
		}
		<-ch
	}
}
