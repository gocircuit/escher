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
	eu "github.com/gocircuit/escher/understand"
)

// Understand
type Understand struct{}

func (Understand) Materialize() think.Reflex {
	seenEndo, seenExo := think.NewSynapse()
	understoodEndo, understoodExo := think.NewSynapse()
	go func() {
		h := &understand{}
		h.understood = understoodEndo.Focus(think.DontCognize)
		seenEndo.Focus(h.CognizeSeen)
	}()
	return think.Reflex{
		"Seen":       seenExo,
		"Understood": understoodExo,
	}
}

type understand struct {
	understood *think.ReCognizer
}

func (h *understand) CognizeSeen(v interface{}) {
	switch t := v.(type) {
	case *es.Circuit:
		h.understood.ReCognize(eu.Understand(t))
	case nil:
		h.understood.ReCognize(nil)
	}
	panic("seen incomprehensible")
}
