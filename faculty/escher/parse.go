// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/kit/plumb"
	"github.com/gocircuit/escher/see"
)

type Parse struct{ be.Sparkless }

func (Parse) CognizeText(eye *be.Eye, v interface{}) {
	src := see.NewSrcString(plumb.AsString(v))
	for {
		v := see.SeeChamber(src)
		if v == nil {
			break
		}
		eye.Show("Value", v)
	}
}

func (Parse) Cognize(eye *be.Eye, v interface{}) {}
