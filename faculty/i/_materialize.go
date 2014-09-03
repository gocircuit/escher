// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package i

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/faculty/basic"
	es "github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/be"
)

// Materialize
type Materialize struct{}

func (Materialize) Materialize() be.Reflex {
	sourceEndo, sourceExo := be.NewSynapse()
	seenEndo, seenExo := be.NewSynapse()
	go func() {
		h := &see{}
		h.seen = seenEndo.Focus(be.DontCognize)
		sourceEndo.Focus(h.CognizeSource)
	}()
	return be.Reflex{
		"Source": sourceExo,
		"Materializen":   seenExo,
	}
}

type see struct {
	seen *be.ReCognizer
}

func (h *see) CognizeSource(v interface{}) {
	src, ok := basic.AsString(v)
	if !ok {
		panic("non-string name perceived by os.see")
	}
	h.seen.ReCognize(es.MaterializeCircuit(es.NewSrcString(src)))
}
