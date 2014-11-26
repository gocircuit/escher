// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register("index.Mirror", be.NewMaterializer(Mirror{}))
	faculty.Register("index.Generalize", be.NewMaterializer(Generalize{}))
	faculty.Register("index.Yield", be.NewMaterializer(Yield{}))
	faculty.Register("index.Lookup", be.NewMaterializer(Lookup{}))
}

/*
	To create an Index node within Escher, use the pattern:

	f Fork
	f:X = :X
	f:Y = :Y
	f: = :Index
*/
