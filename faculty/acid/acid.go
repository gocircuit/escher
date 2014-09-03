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
	// . "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/kit/plumb"
)

func Init(x, y string) {
	ns := faculty.Root.Refine("acid")
	ns.AddTerminal("XDir", Dir{x})
	ns.AddTerminal("YDir", Dir{y})
}

// Dir
type Dir struct{
	dir string
}

func (d Dir) Materialize(*be.Matter) be.Reflex {
	x := dir(d.dir)
	reflex, _ := plumb.NewEyeCognizer(x.Cognize, "Path", "_")
	return reflex
}

type dir string

func (d dir) Cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	if dvalve != "Path" {
		return
	}
	// img := value.(Image)
	panic("not ready")
	// eye.Show("_", ??)
}
