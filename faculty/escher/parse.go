// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "log"

	"github.com/hoijui/escher/a"
	"github.com/hoijui/escher/be"
	"github.com/hoijui/escher/kit/plumb"
	"github.com/hoijui/escher/see"
)

type Parse struct{ be.Sparkless }

func (Parse) CognizeText(eye *be.Eye, v interface{}) {
	src := a.NewSrcString(plumb.AsString(v))
	for {
		v := see.SeeChamber(src)
		if v == nil {
			break
		}
		eye.Show("Value", v)
	}
}

func (Parse) Cognize(eye *be.Eye, v interface{}) {}
