// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

type Mix struct{}

func (h *Mix) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return &Mix{}
}

func (h *Mix) CognizePN(eye *be.Eye, v interface{}) {
	eye.Show(
		DefaultValve, 
		combine(
			v.(Circuit).CircuitAt("Positive"),
			v.(Circuit).CircuitAt("Negative"),
		),
	)
}

func (h *Mix) Cognize(*be.Eye, interface{}) {}

// combine substitutes dot-prefix addresses in the pos with corresponding data from the neg.
func combine(pos, neg Circuit) Circuit {
	pos = pos.Clone()
	for gname, gvalue := range pos.Gates() {
		switch t := gvalue.(type) {
		case Address:
			if len(t.Path) > 0 && t.Path[0] == "" { // If the gate meaning is a substitution address
				pos.ReGrow(gname, neg.Goto(t.Path[1:]...))
			}
		case Circuit:
			pos.ReGrow(gname, combine(t, neg))
		}
		if s, ok := gname.(Address); ok && len(s.Path) > 0 && s.Path[0] == "" { // If the gate name is a substitution address
			rename(pos, gname, neg.Goto(s.Path[1:]...))
		}
	}
	return pos
}

func rename(u Circuit, g, h Name) Circuit {
	// gate
	x := u.Exclude(g)
	if x == nil {
		panic(2)
	}
	u.Include(h, x)
	// links
	for vlv, vec := range u.Valves(g) {
		u.Unlink(Vector{g, vlv}, vec)
		u.Link(Vector{h, vlv}, vec)
	}
	return u
}
