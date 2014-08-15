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
	faculty.Root.AddTerminal("3Way", Way3{})
	faculty.Root.AddTerminal("4Way", Way4{})
	faculty.Root.AddTerminal("5Way", Way5{})
}

// Way3
type Way3 struct{}

func (Way3) Materialize() think.Reflex {
	w0Endo, w0Exo := think.NewSynapse()
	w1Endo, w1Exo := think.NewSynapse()
	w2Endo, w2Exo := think.NewSynapse()
	go func() {
		h := &way3{}
		h.re[0] = w0Endo.Focus(func(v interface{}) { h.Cognize(0, v) })
		h.re[1] = w1Endo.Focus(func(v interface{}) { h.Cognize(1, v) })
		h.re[2] = w2Endo.Focus(func(v interface{}) { h.Cognize(2, v) })
	}()
	return think.Reflex{
		"0": w0Exo, 
		"1": w1Exo, 
		"2": w2Exo,
	}
}

type way3 struct {
	re [3]*think.ReCognizer
}

func (h *way3) Cognize(way int, v interface{}) {
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

// Way4
type Way4 struct{}

func (Way4) Materialize() think.Reflex {
	w0Endo, w0Exo := think.NewSynapse()
	w1Endo, w1Exo := think.NewSynapse()
	w2Endo, w2Exo := think.NewSynapse()
	w3Endo, w3Exo := think.NewSynapse()
	go func() {
		h := &way4{}
		h.re[0] = w0Endo.Focus(func(v interface{}) { h.Cognize(0, v) })
		h.re[1] = w1Endo.Focus(func(v interface{}) { h.Cognize(1, v) })
		h.re[2] = w2Endo.Focus(func(v interface{}) { h.Cognize(2, v) })
		h.re[3] = w3Endo.Focus(func(v interface{}) { h.Cognize(3, v) })
	}()
	return think.Reflex{
		"0": w0Exo, 
		"1": w1Exo, 
		"2": w2Exo,
		"3": w3Exo,
	}
}

type way4 struct {
	re [4]*think.ReCognizer
}

func (h *way4) Cognize(way int, v interface{}) {
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

// Way5
type Way5 struct{}

func (Way5) Materialize() think.Reflex {
	w0Endo, w0Exo := think.NewSynapse()
	w1Endo, w1Exo := think.NewSynapse()
	w2Endo, w2Exo := think.NewSynapse()
	w3Endo, w3Exo := think.NewSynapse()
	w4Endo, w4Exo := think.NewSynapse()
	go func() {
		h := &way5{}
		h.re[0] = w0Endo.Focus(func(v interface{}) { h.Cognize(0, v) })
		h.re[1] = w1Endo.Focus(func(v interface{}) { h.Cognize(1, v) })
		h.re[2] = w2Endo.Focus(func(v interface{}) { h.Cognize(2, v) })
		h.re[3] = w3Endo.Focus(func(v interface{}) { h.Cognize(3, v) })
		h.re[4] = w4Endo.Focus(func(v interface{}) { h.Cognize(4, v) })
	}()
	return think.Reflex{
		"0": w0Exo, 
		"1": w1Exo, 
		"2": w2Exo,
		"3": w3Exo,
		"4": w4Exo,
	}
}

type way5 struct {
	re [5]*think.ReCognizer
}

func (h *way5) Cognize(way int, v interface{}) {
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
