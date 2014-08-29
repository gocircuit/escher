// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/be"
)

// ForkCommand…
type ForkCommand struct{}

func (ForkCommand) Materialize() be.Reflex {
	return basic.MaterializeConjunction("_", "Path", "Dir", "Args", "Env")
}

// ForkIO…
type ForkIO struct{}

func (ForkIO) Materialize() be.Reflex {
	return basic.MaterializeConjunction("_", "Spawn", "Stdin", "Stdout", "Stderr")
}

// ForkExit…
type ForkExit struct{}

func (ForkExit) Materialize() be.Reflex {
	return basic.MaterializeConjunction("_", "Spawn", "Exit")
}