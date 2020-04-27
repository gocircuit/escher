// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
)

type OneWayDoor struct { // :From :To :Door
	flow chan struct{}
}

func (w *OneWayDoor) Spark(eye *be.Eye, _ cir.Circuit, aux ...interface{}) cir.Value {
	w.flow = make(chan struct{})
	return nil
}

func (w *OneWayDoor) CognizeFrom(eye *be.Eye, value interface{}) {
	<-w.flow
	eye.Show("To", value)
}

func (w *OneWayDoor) CognizeTo(eye *be.Eye, value interface{}) {}

func (w *OneWayDoor) CognizeDoor(eye *be.Eye, value interface{}) {
	w.flow <- struct{}{}
}
