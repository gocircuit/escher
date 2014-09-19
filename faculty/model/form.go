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
	form plumb.Given
	weave plumb.Given
}

func (h *Form) Spark() {
	h.form.Init()
	h.weave.Init()
}

func (h *Form) CognizeForm(eye *be.Eye, v interface{}) {
	h.form.Fix(v)
	eye.Show(
		"_", 
		combine(
			v.(Circuit), 
			h.weave.Use().(Circuit),
		),
	)
}

func (h *Form) CognizeWeave(eye *be.Eye, v interface{}) {
	h.weave.Fix(v)
	eye.Show(
		"_", 
		combine(
			h.form.Use().(Circuit),
			v.(Circuit), 
		),
	)
}

func (h *Form) Cognize_(*be.Eye, interface{}) {}

// combine substitutes dot-prefix addresses in the form with corresponding data from the weave.
func combine(form, weave Circuit) Circuit {
	form = form.Clone()
	for gname, gvalue := range form.Gates() {
		switch t := gvalue.(type) {
		case Address:
			if len(t.Path()) > 0 && t.Path()[0] == "" { // If the gate meaning is a substitution address
				form.ReGrow(gname, weave.Lookup(t.Path()[1:]...))
			}
		case Circuit:
			form.ReGrow(gname, combine(t, weave))
		}
		if s, ok := gname.(Address); ok && len(s.Path()) > 0 && s.Path()[0] == "" { // If the gate name is a substitution address
			rename(form, gname, weave.Lookup(s.Path()[1:]...))
		}
	}
	return form
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
