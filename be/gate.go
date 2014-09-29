// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "log"
	. "reflect"

	"github.com/gocircuit/escher/circuit"
)

// type RetinaCognizer func(eye *Eye, value interface{})

const prefix = "Cognize"

func MaterializeInterface(v Gate, matter *Matter) (Reflex, circuit.Value) {
	w := makeGate(v)
	spark := w.Interface().(Gate).Spark(matter) // Initialize
	r := gate{w}
	var valve []string
	t := r.Value.Type()
	for i := 0; i < t.NumMethod(); i++ {
		n := t.Method(i).Name
		if len(n) >= len(prefix) && n[:len(prefix)] == prefix {
			valve = append(valve, n[len(prefix):])
		}
	}
	x, _ := NewEyeCognizer(r.Cognize, valve...)
	return x, spark
}

type gate struct {
	Value
}

func (r *gate) Cognize(eye *Eye, valve string, value interface{}) {
	m := r.Value.MethodByName(prefix + valve)
	m.Call(
		[]Value{
			ValueOf(eye), 
			ValueOf(value),
		},
	)
}

func makeGate(like interface{}) Value {
	t := TypeOf(like)
	switch t.Kind() {
	case Ptr: // Pointer types are allocated
		return New(t.Elem()).Convert(t)
	default: // Value-based types are used as is
		return ValueOf(like)
	}
	panic(0)
}
