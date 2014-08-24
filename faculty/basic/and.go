// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
)

func MaterializeAnd(name string, parts ...string) think.Reflex {
	reflex, eye := faculty.NewEye(append(parts, name)...)
	go func() {
		h := &and{
			name:      name,
			parts:     parts,
			connected: make(chan struct{}),
		}
		h.reply = eye.Focus(h.ShortCognize)
		close(h.connected)
	}()
	return reflex
}

type and struct {
	name      string
	parts     []string
	connected chan struct{}
	reply     *faculty.EyeNerve
}

func (h *and) ShortCognize(mem faculty.Impression) {
	<-h.connected
	recent := mem.Index(0) // most-recently perceived change
	g := faculty.MakeImpression()
	if recent.Valve() == h.name { // if most recently updated valve is the anded image
		anded := recent.Value().(Image)
		for i, part := range h.parts {
			g.Show(i, part, anded[part])
		}
	} else { // if the most-recently updated valve is one of the parts, update the anded image
		x := Make()
		for _, part := range h.parts {
			x.Grow(part, mem.Valve(part).Value())
		}
		g.Show(0, h.name, x)
	}
	h.reply.ReCognize(g)
}
