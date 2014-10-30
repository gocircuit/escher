// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"log"
	"os"
	. "reflect"

	"github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/kit/runtime"
)

const cognizePrefix = "Cognize"
const cognizeEllipses = "OverCognize"

// NewNativeMaterializer returns a materializer that generates copies of sample and sparks them with the aux data.
func NewNativeMaterializer(sample Native, aux ...interface{}) Materializer {
	return &NativeMaterializer{sample, aux}
}

type NativeMaterializer struct {
	sample Native
	aux []interface{}
}

func (x *NativeMaterializer) Materialize(matter *Matter) (Reflex, circuit.Value) {
	return MaterializeNative(matter, x.sample, x.aux...)
}

func (x *NativeMaterializer) String() string {
	return fmt.Sprintf("Native(%T)", x.sample)
}

// MaterializeNative materializes the native implementation v.
// It returns the resulting reflex and residue, but not the Go-facing instance.
func MaterializeNative(matter *Matter, v Native, aux ...interface{}) (reflex Reflex, residue circuit.Value) {
	reflex, residue, _ = MaterializeNativeInstance(matter, v, aux...)
	return
}

// MaterializeNativeInstance materializes the native implementation v.
// It returns the resulting reflex and residue, as well as the Go-facing instance.
func MaterializeNativeInstance(matter *Matter, v Native, aux ...interface{}) (Reflex, circuit.Value, interface{}) {
	w := makeNative(v)
	r := gate{matter, w, w.Type()}
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
	reflex, eye := NewEyeCognizer(r.Cognize, valve...)
	return reflex, w.Interface().(Native).Spark(eye, matter, aux...), w.Interface()
}

// gate is a materialized native reflex.
type gate struct {
	*Matter
	Value
	Type
}

func (g *gate) Cognize(eye *Eye, valve circuit.Name, value interface{}) {
	// Catch panics during cognizing and report their context to the user
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic:\n%v\nRecovered: %v\n\n", g.Matter, r)
			runtime.PrintStack()
			os.Exit(1)
		}
	}()

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
		if _, ok := g.Type.MethodByName(cognizePrefix + alias); ok {
			m := g.Value.MethodByName(cognizePrefix + alias)
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
	m := g.Value.MethodByName(cognizeEllipses)
	m.Call(
		[]Value{
			ValueOf(eye),
			ValueOf(valve),
			ValueOf(value),
		},
	)
}

// makeNative creates a copy of like.
// Pointer types allocate the object pointed to and copy that object as well.
func makeNative(like interface{}) Value {
	t := TypeOf(like)
	switch t.Kind() {
	case Ptr: // Pointer types are allocated
		return New(t.Elem()).Convert(t)
	default: // Value-based types are used as is
		return ValueOf(like)
	}
	panic(0)
}
