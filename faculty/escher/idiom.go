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
)

type Idiom struct {
	be.Idiom
}

func (n Idiom) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	n.Idiom = matter.Idiom
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

func (n Idiom) OverCognize(*be.Eye, Name, interface{}) {}
