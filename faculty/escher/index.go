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

type Index struct {
	be.Index
}

func (n Index) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	n.Index = matter.Index
	go func() {
		for vlv, _ := range matter.View.Gate {
			eye.Show(vlv, n.Index)
		}
	}()
	if matter.View.Len() == 0 {
		return n.Index
	}
	return nil
}

func (n Index) OverCognize(*be.Eye, Name, interface{}) {}
