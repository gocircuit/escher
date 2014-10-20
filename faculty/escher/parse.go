// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/see"
)

type Parse struct{}

func (Parse) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Parse) CognizeChunk(eye *be.Eye, v interface{}) {
	src := see.NewSrcString(string(v.([]byte)))
	for {
		_, v := see.Parse(src)
		if v == nil {
			break
		}
		eye.Show("Value", v)
	}
}

func (Parse) CognizeValue(eye *be.Eye, v interface{}) {}
