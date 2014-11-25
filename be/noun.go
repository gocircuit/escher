// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"

	. "github.com/gocircuit/escher/circuit"
)

// Idle
type Idle struct{}

func (Idle) Spark(*Eye, Circuit, ...interface{}) Value {
	return nil
}

func (Idle) OverCognize(*Eye, Name, interface{}) {}

func NewIdleMaterializer() Materializer {
	return NewNativeMaterializer(Idle{})
}

// Noun

func MaterializeNoun(matter Circuit, v interface{}) (Reflex, Value) {
	return MaterializeNative(matter, &Noun{}, v)
}

func NewNoun(v interface{}) Materializer {
	return NewNativeMaterializer(&Noun{}, v)
}

type Noun struct {
	Value interface{}
}

func (n *Noun) Spark(eye *Eye, matter Circuit, aux ...interface{}) Value {
	n.Value = aux[0]
	go func() {
		for vlv, _ := range matter.CircuitAt("View").Gate {
			eye.Show(vlv, aux[0])
		}
	}()
	if matter.View.Len() == 0 {
		return aux[0]
	}
	return nil
}

func (n *Noun) OverCognize(*Eye, Name, interface{}) {}

func (n *Noun) NativeString(aux ...interface{}) string {
	return fmt.Sprintf("Noun(%v)", aux[0])
}

// Future
type Future struct {
	eye  *Eye
	view Circuit
}

func (f *Future) Spark(eye *Eye, matter Circuit, _ ...interface{}) Value {
	f.eye = eye
	f.view = matter.CircuitAt("View")
	return nil
}

func (f *Future) Charge(v Value) {
	go func() {
		for vlv, _ := range f.view.Gate {
			f.eye.Show(vlv, DeepCopy(v))
		}
	}()
}

func (f *Future) OverCognize(*Eye, Name, interface{}) {}
