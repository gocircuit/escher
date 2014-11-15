// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("index.Mirror", be.NewNativeMaterializer(Mirror{}))
	faculty.Register("index.Generalize", be.NewNativeMaterializer(Generalize{}))
	faculty.Register("index.Yield", be.NewNativeMaterializer(Yield{}))
}

/*
	To create an Index node within Escher, use the pattern:

	f Fork
	f:? = "Index"
	f:X = :X
	f:Y = :Y
	f: = :Index
*/