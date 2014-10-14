// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package view

import (
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

// Focus
type Focus struct {
	name plumb.Given // gate name to focus on
}

func (f *Focus) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (f *Focus) CognizeName(eye *be.Eye, v interface{}) {
	f.name.Fix(v)
}

func (f *Focus) CognizeView(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, v.(Circuit).At(f.name.Use()))
}

func (f *Focus) Cognize(eye *be.Eye, v interface{}) {}
