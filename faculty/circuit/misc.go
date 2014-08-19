// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package circuit provides Escher gates for building dynamic cloud applications using the circuit runtime of http://gocircuit.org
package circuit

import (
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty/basic"
)

// ForkExit…
type ForkExit struct{}

func (ForkExit) Materialize() think.Reflex {
	return basic.MaterializeFork("Forked", "Spawn", "Exit")
}

// ForkIO…
type ForkIO struct{}

func (ForkIO) Materialize() think.Reflex {
	return basic.MaterializeFork("Forked", "Spawn", "Stdin", "Stdout", "Stderr")
}
