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
	"github.com/gocircuit/escher/plumb"
)

type Form struct{
	pos, neg plumb.Given
}

func (h *Form) Spark() {
	h.pos.Init()
	h.neg.Init()
}

func (h *Form) CognizePositive(eye *be.Eye, v interface{}) {
	h.pos.Fix(v)
	eye.Show(
		"_", 
		combine(
			v.(Circuit), 
			h.neg.Use().(Circuit),
		),
	)
}

func (h *Form) CognizeNegative(eye *be.Eye, v interface{}) {
	h.neg.Fix(v)
	eye.Show(
		"_", 
		combine(
			h.pos.Use().(Circuit),
			v.(Circuit), 
		),
	)
}

func (h *Form) Cognize_(*be.Eye, interface{}) {}

// combine substitutes dot-prefix addresses in the pos with corresponding data from the neg.
func combine(pos, neg Circuit) Circuit {
	pos = pos.Clone()
	for gname, gvalue := range pos.Gates() {
		switch t := gvalue.(type) {
		case Address:
			if len(t.Path()) > 0 && t.Path()[0] == "" { // If the gate meaning is a substitution address
				pos.ReGrow(gname, neg.Lookup(t.Path()[1:]...))
			}
		case Circuit:
			pos.ReGrow(gname, combine(t, neg))
		}
		if s, ok := gname.(Address); ok && len(s.Path()) > 0 && s.Path()[0] == "" { // If the gate name is a substitution address
			rename(pos, gname, neg.Lookup(s.Path()[1:]...))
		}
	}
	return pos
}

func rename(u Circuit, g, h Name) Circuit {
	// gate
	x, ok := u.Exclude(g)
	if !ok {
		panic(2)
	}
	u.Include(h, x)
	// links
	for vlv, vec := range u.Valves(g) {
		u.Unlink(NewVector(g, vlv), vec)
		u.Link(NewVector(h, vlv), vec)
	}
	return u
}
