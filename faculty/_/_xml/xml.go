// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"sync"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/be"
)

type Reservoir struct{
	once sync.Once
	drain chan io.WriteCloser
	x Circuit
}

func (r *Reservoir) Is() {
	r.drain = make(chan io.WriteCloser, 1)
	r.x = New()
}

func (r *Reservoir) CognizeY(eye *be.Eye, v interface{}) {
	r.once.Do(func() {r.drain <- v.(io.WriteCloser) })
}

/*
	{
		Command string // Up, Down, End, Include(Name, Value), Exclude(Name), Link([]{Image, Valve})
		… // arguments
	}
*/
func (r *Reservoir) CognizeX(eye *be.Eye, v interface{}) {
	u := v.(Circuit)
	switch u.StringAt("Command") {
	case "Open":
	case "Close":
	case "Include":
	case "Exclude":
	case "Tie":
	case "Stop":
		?
		w := <-r.drain
	}
}
