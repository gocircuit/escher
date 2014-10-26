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
	"github.com/gocircuit/escher/kit/plumb"
)

type Idiom struct {
	Circuit
}

func (n Idiom) Spark(eye *Eye, matter *Matter, aux ...interface{}) Value {
	n.Idiom = matter.Idiom.DeepCopy()
	go func() {
		for vlv, _ := range matter.View.Gate {
			eye.Show(vlv, n.Idiom)
		}
	}()
	if matter.View.Len() == 0 {
		return n.Idiom
	}
	return nil
}

func (n Idiom) OverCognize(*Eye, Name, interface{}) {}
