// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package testing

import (
	// "log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// 
type Exec struct {}

func (Exec) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Exec) CognizeIn(eye *be.Eye, v interface{}) {
	panic(0)
}

func (Exec) CognizeOut(eye *be.Eye, v interface{}) {}
