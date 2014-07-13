// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package think


// Reflex is a bundle of un-attached sense endpoints
type Reflex map[string]*Synapse

type Gate interface {
	Materialize() Reflex
}

// Ignore gates ignore their empty-string valve
type Ignore struct{}

func (Ignore) Materialize() Reflex {
	s, t := NewSynapse()
	go func() {
		s.Attach(DontCognize)
	}()
	return Reflex{"": t}
}
