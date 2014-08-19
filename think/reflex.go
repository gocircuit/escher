// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

// Reflex is a bundle of not yet attached sense endpoints (synapses).
type Reflex map[string]*Synapse

type Gate interface {
	Materialize() Reflex
}

// Ignore gates ignore their empty-string valve
type Ignore struct{}

func (Ignore) Materialize() Reflex {
	s, t := NewSynapse()
	go func() {
		s.Focus(DontCognize)
	}()
	return Reflex{"Subject": t}
}
