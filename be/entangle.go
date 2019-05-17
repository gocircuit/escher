// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"sync"

	cir "github.com/gocircuit/escher/circuit"
)

// NewEntanglement returns two materializers that each materialize once, to
// a gate with a default valve which is connected to the other.
func NewEntanglement() (p, n *Entanglement) {
	x, y := NewSynapse()
	return &Entanglement{synapse: x}, &Entanglement{synapse: y}
}

type Entanglement struct {
	sync.Mutex
	synapse *Synapse
}

func (em *Entanglement) Materialize(given Reflex, _ cir.Circuit) cir.Value {
	em.Lock()
	defer em.Unlock()
	if len(given) != 1 {
		panic(2)
	}
	y := em.synapse
	em.synapse = nil
	go Link(given[cir.DefaultValve], y)
	return nil
}

func (em *Entanglement) Synapse() *Synapse {
	em.Lock()
	defer em.Unlock()
	y := em.synapse
	em.synapse = nil
	return y
}
