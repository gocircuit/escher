// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package cmplx

import (
	// "math/cmplx"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(Planar{}), "cmplx", "Planar")
	faculty.Register(be.NewMaterializer(Polar{}), "cmplx", "Polar")
}

// Planar
type Planar struct{}

func (Planar) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Planar) CognizeComplex(eye *be.Eye, v interface{}) {
	eye.Show("Planar", New().Grow("X", real(v.(complex128))).Grow("Y", imag(v.(complex128))))
}

func (Planar) CognizePlanar(eye *be.Eye, v interface{}) {
	x := v.(Circuit)
	eye.Show("Complex", complex(x.FloatAt("X"), x.FloatAt("Y")))
}
