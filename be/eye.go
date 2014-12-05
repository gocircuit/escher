// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
)

// Eye is a runtime facility that delivers messages by invoking gate methods and
// provides methods that the gate can use to send messages out.
//
// Eye is an implementation of Leslie Valiant's “Mind's Eye”, described in
//	http://www.probablyapproximatelycorrect.com/
// The mind's eye is a synchronization device which sees changes as ordered
// and thus introduces the illusory perception of time (and, eventually, of the
// higher-level concepts of cause and effect).
//
type Eye struct {
	show map[Name]nerve
}

type change struct {
	Valve Name
	Value interface{}
}

type EyeCognizer func(eye *Eye, valve Name, value interface{})

func NewEye(given Reflex, cog EyeCognizer) (eye *Eye) {
	eye = &Eye{show: make(map[Name]nerve)}
	for vlv_, syn_ := range given {
		vlv, syn := vlv_, syn_
		n := make(nerve, 1)
		eye.show[vlv] = n
		go func() {
			n <- syn.Connect(
				func(w interface{}) {
					cog(eye, vlv, w)
				},
			)
		}()
	}
	return
}

type nerve chan *ReCognizer

func (eye *Eye) Show(valve Name, v interface{}) {
	n := eye.show[valve]
	r := <-n
	defer func() {
		n <- r
	}()
	r.ReCognize(v)
}
