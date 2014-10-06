// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package model provides a basis of gates for circuit transformations.
package model

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("model.Hamiltonian", be.NewGateMaterializer(&Hamiltonian{}, nil))
	faculty.Register("model.Eulerian", be.NewGateMaterializer(&Eulerian{}, nil))
	faculty.Register("model.Reservoir", be.NewGateMaterializer(&Reservoir{}, nil))
	faculty.Register("model.Mix", be.NewGateMaterializer(&Mix{}, nil))
	faculty.Register("model.Range_", be.NewGateMaterializer(&Range{}, nil))
	faculty.Register("model.IO", be.NewGateMaterializer(IO{}, nil))
}
