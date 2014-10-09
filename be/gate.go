// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"log"
	. "reflect"

	"github.com/gocircuit/escher/circuit"
)

const cognizePrefix = "Cognize"
const cognizeEllipses = "OverCognize"

// NewGateMaterializer returns a materializer that generates copies of sample and sparks them with the aux data.
func NewGateMaterializer(sample Gate, aux ...interface{}) MaterializerWithMatterFunc {
	return func(matter *Matter) (Reflex, circuit.Value) {
		return materializeGate(matter, sample, aux...)
	}
}

func materializeGate(matter *Matter, v Gate, aux ...interface{}) (Reflex, circuit.Value) {
	w := makeGate(v)
	spark := w.Interface().(Gate).Spark(matter, aux...) // Initialize
	r := gate{w, w.Type()}
	// Enumerate the valves handled by dedicated methods.
	dedicated := make(map[circuit.Name]struct{})
	for i := 0; i < r.Type.NumMethod(); i++ {
		n := r.Type.Method(i).Name
		if len(n) >= len(cognizePrefix) && n[:len(cognizePrefix)] == cognizePrefix {
			dedicated[n[len(cognizePrefix):]] = struct{}{}
		}
	}
	// Verify that all connected valves in matter have handlers or that there is a generic cognizer method.
	var valve []circuit.Name
	_, over := r.Type.MethodByName(cognizeEllipses)
	for vlv, _ := range matter.View.Gate {
		valve = append(valve, vlv)
		if over {
			continue
		}
		if _, ok := dedicated[vlv]; !ok {
			log.Fatalf("gate %T does not have methods to handle the connected valve %v", v, vlv)
		}
	}
	// Not all handled valves need to be connected. But all connected valves need to be handled by a gate method.
	x, _ := NewEyeCognizer(r.Cognize, valve...)
	return x, spark
}

type gate struct {
	Value
	Type
}

func (r *gate) Cognize(eye *Eye, valve circuit.Name, value interface{}) {
	// Compute valve string
	var alias string
	var letter bool
	switch valve.(type) {
	case string, int:
		alias = fmt.Sprintf("%v", valve)
		letter = true
	default:
		letter = false
	}

	// If there is a dedicated method for valve, use that.
	if letter {
		if _, ok := r.Type.MethodByName(cognizePrefix + alias); ok {
			m := r.Value.MethodByName(cognizePrefix + alias)
			m.Call(
				[]Value{
					ValueOf(eye), 
					ValueOf(value),
				},
			)
			return
		}
	}
	// Otherwise call the generic cognizer
	m := r.Value.MethodByName(cognizeEllipses)
	m.Call(
		[]Value{
			ValueOf(eye),
			ValueOf(valve),
			ValueOf(value),
		},
	)
}

// makeGate creates a new value of the same type as like. Pointer types allocate the object pointed to.
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
