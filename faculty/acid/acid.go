// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package acid provides gates for accessing files from the X and Y (re)source directories of the Escher program.
package acid

import (
	// "path"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

func Init(x, y, z string) {
	faculty.Register("acid.XDir", Dir{x})
	faculty.Register("acid.YDir", Dir{y})
	faculty.Register("acid.ZDir", Dir{z})
}

// Dir
type Dir struct{
	dir string
}

func (d Dir) Materialize() (be.Reflex, Value) {
	x := dir(d.dir)
	reflex, _ := be.NewEyeCognizer(x.Cognize, "Path", DefaultValve)
	return reflex, d.dir
}

type dir string

func (d dir) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	if dvalve != "Path" {
		return
	}
	eye.Show(DefaultValve, string(d))
}
