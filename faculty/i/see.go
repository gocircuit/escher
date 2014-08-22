// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package i provides reflection primitives.
package i

import (
	es "github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/faculty/basic"
)

func init() {
	ns := faculty.Root.Refine("i")
	ns.AddTerminal("See", See{})
	ns.AddTerminal("Understand", Understand{})
}

// See
type See struct{}

func (See) Materialize() think.Reflex {
	sourceEndo, sourceExo := think.NewSynapse()
	seenEndo, seenExo := think.NewSynapse()
	go func() {
		h := &see{}
		h.seen = seenEndo.Focus(think.DontCognize)
		sourceEndo.Focus(h.CognizeSource)
	}()
	return think.Reflex{
		"Source": sourceExo,
		"Seen": seenExo,
	}
}

type see struct {
	seen *think.ReCognizer
}

func (h *see) CognizeSource(v interface{}) {
	src, ok := basic.AsString(v)
	if !ok {
		panic("non-string name perceived by os.see")
	}
	h.seen.ReCognize(es.SeeCircuit(es.NewSrcString(src)))
}
