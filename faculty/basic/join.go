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
	. "github.com/gocircuit/escher/image"
)

func MaterializeFork(name string, parts ...string) think.Reflex {
	reflex, eye := faculty.NewEye(append(parts, name)...)
	go func() {
		h := &join{
			name: name,
			parts: parts,
			ready: make(chan struct{}),
		}
		h.reply = eye.Focus(h.ShortCognize)
		close(h.ready)
	}()
	return reflex
}

type join struct {
	name string
	parts []string
	ready chan struct{}
	reply *faculty.EyeReCognizer
}

func (h *join) ShortCognize(mem faculty.Impression) {
	<-h.ready
	recent := mem.Index(0) // most-recently perceived change
	g := faculty.MakeImpression()
	if recent.Valve() == h.name { // if most recently updated valve is the joined image
		joined := recent.Value().(Image)
		for _, part := range h.parts {
			g.Show(0, part, joined[part])
		}
	} else { // if the most-recently updated valve is one of the parts, update the joined image
		x := Make()
		for _, part := range h.parts {
			x.Grow(part, mem.Valve(part).Value())
		}
		g.Show(0, h.name, x)
	}
	h.reply.ReCognize(g)
}
