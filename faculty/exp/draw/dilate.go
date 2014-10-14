// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package draw

import (
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

// Dilate
type Dilate struct {
	factor plumb.Given // dilation factor
}

func (f *Dilate) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	f.factor.Init()
	return nil
}

func (f *Dilate) CognizeFactor(eye *be.Eye, v interface{}) {
	f.factor.Fix(v)
}

func (f *Dilate) CognizeView(eye *be.Eye, v interface{}) {
	w := v.(Circuit)
	w.Include("Orientation", w.ComplexAt("Orientation") * f.factor.Use().(complex128))
	eye.Show(DefaultValve, w)
}

func (f *Dilate) Cognize(eye *be.Eye, v interface{}) {}
