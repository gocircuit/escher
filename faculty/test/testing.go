// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package test

import (
	// . "github.com/hoijui/escher/circuit"
	"github.com/hoijui/escher/be"
	"github.com/hoijui/escher/faculty"
)

func Init(srcdir string) {
	srcDir = srcdir
	faculty.Register(be.NewMaterializer(&Match{}), "test", "Match")
	faculty.Register(be.NewMaterializer(Filter{}), "test", "Filter")
	faculty.Register(be.NewMaterializer(Exec{}), "test", "Exec")
}

var srcDir string
