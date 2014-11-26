// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"io"
	"io/ioutil"

	. "github.com/gocircuit/escher/circuit"
)

// Sink

func NewSink() Materializer {
	return NewMaterializer(sink{})
}

type sink struct{}

func (sink) Spark(*Eye, Circuit, ...interface{}) Value {
	return nil
}

func (sink) OverCognize(_ *Eye, _ Name, v interface{}) {
	SinkValue(v)
}

func SinkValue(v interface{}) {
	switch t := v.(type) {
	case Circuit:
		for _, g := range t.Gate {
			SinkValue(g)
		}
	case io.Closer:
		t.Close()
	case io.Reader:
		io.Copy(ioutil.Discard, t)
	}
}

// Source

func NewSource(v interface{}) Materializer {
	return NewMaterializer(&source{}, v)
}

func MaterializeSource(given Reflex, matter Circuit, v interface{}) (Reflex, Value) {
	return Materialize(given, matter, &source{}, v)
}

type source struct {
	Value interface{}
}

func (n *source) Spark(eye *Eye, matter Circuit, aux ...interface{}) Value {
	n.Value = aux[0]
	go func() {
		for vlv, _ := range matter.CircuitAt("View").Gate {
			eye.Show(vlv, aux[0])
		}
	}()
	if matter.CircuitAt("View").Len() == 0 {
		return aux[0]
	}
	return nil
}

func (n *source) OverCognize(*Eye, Name, interface{}) {}

func (n *source) NativeString(aux ...interface{}) string {
	return fmt.Sprintf("Source(%v)", aux[0])
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
