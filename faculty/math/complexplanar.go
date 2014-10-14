// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package math

import (
	// "math/cmplx"

	// "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// ComplexPlanar
type ComplexPlanar struct{}

func (ComplexPlanar) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (ComplexPlanar) CognizeComplex(eye *be.Eye, v interface{}) {
	eye.Show("Planar", New().Grow("X", real(v.(complex128))).Grow("Y", imag(v.(complex128))))
}

func (ComplexPlanar) CognizePlanar(eye *be.Eye, v interface{}) {
	x := v.(Circuit)
	eye.Show("Complex", complex(x.FloatAt("X"),  x.FloatAt("Y")))
}
