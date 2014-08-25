// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package i

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
	es "github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/be"
)

func init() {
	ns := faculty.Root.Refine("i")
	ns.AddTerminal("See", See{})
	ns.AddTerminal("Understand", Understand{})
	ns.AddTerminal("Memory", Memory{})
	// ns.AddTerminal("Materialize", Materialize{})
}

// See
type See struct{}

func (See) Materialize() be.Reflex {
	sourceEndo, sourceExo := be.NewSynapse()
	seenEndo, seenExo := be.NewSynapse()
	go func() {
		h := &see{}
		h.seen = seenEndo.Focus(be.DontCognize)
		sourceEndo.Focus(h.CognizeSource)
	}()
	return be.Reflex{
		"Source": sourceExo,
		"Seen":   seenExo,
	}
}

type see struct {
	seen *be.ReCognizer
}

func (h *see) CognizeSource(v interface{}) {
	h.seen.ReCognize(es.SeeCircuit(es.NewSrcString(plumb.AsString(v))))
}
