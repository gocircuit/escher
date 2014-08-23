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
	"github.com/gocircuit/escher/think"
)

// Materialize
type Materialize struct{}

func (Materialize) Materialize() think.Reflex {
	sourceEndo, sourceExo := think.NewSynapse()
	seenEndo, seenExo := think.NewSynapse()
	go func() {
		h := &see{}
		h.seen = seenEndo.Focus(think.DontCognize)
		sourceEndo.Focus(h.CognizeSource)
	}()
	return think.Reflex{
		"Source": sourceExo,
		"Materializen":   seenExo,
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
	h.seen.ReCognize(es.MaterializeCircuit(es.NewSrcString(src)))
}
