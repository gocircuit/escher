// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package testing

import (
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
)

func Init(srcdir string) {
	srcDir = srcdir
	faculty.Register("testing.Match", be.NewNativeMaterializer(&Match{}))
	faculty.Register("testing.FilterAll", be.NewNativeMaterializer(FilterAll{}))
	faculty.Register("testing.Exec", be.NewNativeMaterializer(Exec{}))
}

var srcDir string
