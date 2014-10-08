// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "log"

	. "github.com/gocircuit/escher/circuit"
)

type Cell struct {
	show map[Name]*ReCognizer
	see map[Name]chan interface{}
	ping chan Name
}

func NewCell(r Reflex, _ Value) *Cell {
	x := &Cell{
		show: make(map[Name]*ReCognizer),
		see: make(map[Name]chan interface{}),
		ping: make(chan Name),
	}
	for vlv, syn := range r {
		v := vlv.(Name)
		x.show[v] = syn.Focus(
			func(w interface{}) {
				x.cognize(v, w)
			},
		)
		x.see[v] = make(chan interface{})
	}
	return x
}

// ReCognize
func (x *Cell) ReCognize(valve Name, value interface{}) {
	x.show[valve].ReCognize(value)
}

func (x *Cell) cognize(valve Name, value interface{}) {
	x.ping <- valve
	x.see[valve] <- value
}

func (x *Cell) Cognize() (valve Name, value interface{}) {
	valve = <- x.ping
	return valve, <-x.see[valve]
}
