// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("Junction", Junction3{})
}

// Junction3
type Junction3 struct{}

func (Junction3) Materialize() think.Reflex {
	a0Endo, a0Exo := think.NewSynapse()
	a1Endo, a1Exo := think.NewSynapse()
	a2Endo, a2Exo := think.NewSynapse()
	go func() {
		h := &junction3{
			connected: make(chan struct{}),
		}
		h.re[0] = a0Endo.Focus(func(v interface{}) { h.Cognize(0, v) })
		h.re[1] = a1Endo.Focus(func(v interface{}) { h.Cognize(1, v) })
		h.re[2] = a2Endo.Focus(func(v interface{}) { h.Cognize(2, v) })
		close(h.connected)
	}()
	return think.Reflex{
		"X": a0Exo, 
		"Y": a1Exo, 
		"Z": a2Exo,
	}
}

type junction3 struct {
	connected chan struct{}
	re [3]*think.ReCognizer
}

func (h *junction3) Cognize(way int, v interface{}) {
	<-h.connected
	println("Junction <—", way)
	ch := make(chan struct{})
	for i, re_ := range h.re {
		// println(fmt.Sprintf("Junction *** %T vs %T", i, way))
		if i == way {
			continue
		}
		re := re_
		go func() {
			re.ReCognize(v)
			ch <- struct{}{}
		}()
	}
	for i, _ := range h.re {
		// println(fmt.Sprintf("Junction <…> %#T vs %#T ", i, way))
		if i == way {
			continue
		}
		<-ch
	}
	println("¡spark!")
}
