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

// Die
type Die struct {
	at plumb.Given // die at this age only or after
}

func (d *Die) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	d.at.Init()
	return nil
}

func (d *Die) CognizeAt(eye *be.Eye, v interface{}) {
	d.at.Fix(v)
}

func (d *Die) CognizeView(eye *be.Eye, v interface{}) {
	if v.(Circuit).IntAt("Time") >= d.at.Use().(int) {
		return
	}
	eye.Show(DefaultValve, v)
}

func (d *Die) Cognize(eye *be.Eye, v interface{}) {}
