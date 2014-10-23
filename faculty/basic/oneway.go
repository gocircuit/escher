// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// OneWay
type OneWay struct{}

func (OneWay) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (OneWay) CognizeFrom(eye *be.Eye, value interface{}) {
	eye.Show("To", value)
}

func (OneWay) CognizeTo(eye *be.Eye, value interface{}) {}

// OneWayDoor
type OneWayDoor struct {
	ch chan struct{}
}

func (x *OneWayDoor) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	x.ch = make(chan struct{}, 1)
	return nil
}

func (x *OneWayDoor) CognizeDoor(eye *be.Eye, value interface{}) {
	x.ch <- struct{}{}
}

func (x *OneWayDoor) CognizeFrom(eye *be.Eye, value interface{}) {
	<-x.ch
	eye.Show("To", value)
}

func (x *OneWayDoor) CognizeTo(eye *be.Eye, value interface{}) {}
