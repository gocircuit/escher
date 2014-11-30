// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package index provides gates for manipulating circuits interpreted as hierarchical indices.
package index

import (
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(Lookup{}), "index", "Lookup")
	faculty.Register(be.NewMaterializer(Mirror{}), "index", "Mirror")
	faculty.Register(be.NewMaterializer(Yield{}), "index", "Yield")
}

/*
	To create an Index node within Escher, use the pattern:

	f Fork
	f:Alice = :X
	f:Bob = :Y
	f: = :ResultIndex
*/
