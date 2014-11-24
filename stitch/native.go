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
	"reflect"

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
	aux    []interface{}
}

func (x *NativeMaterializer) Materialize(matter *Matter) (Reflex, circuit.Value) {
	return MaterializeNative(matter, x.sample, x.aux...)
}

func (x *NativeMaterializer) String() string {
	if ns, ok := x.sample.(interface {
		NativeString(...interface{}) string
	}); ok {
		return ns.NativeString(x.aux...)
	}
	return fmt.Sprintf("Native(%T)", x.sample)
}

// MaterializeNative materializes the native implementation v.
// It returns the resulting reflex and residue, but not the Go-facing instance.
func MaterializeNative(matter *Matter, v Native, aux ...interface{}) (reflex Reflex, residue circuit.Value) {
	reflex, residue, _ = MaterializeNativeInstance(matter, v, aux...)
	return
}

// StitchNativeInstance materializes the native implementation v.
// It returns the resulting reflex and residue, as well as the Go-facing instance.
func StitchNativeInstance(given Reflex, matter Circuit, v Native, aux ...interface{}) (expected Reflex, residue, obj interface{}) {

	// Build gate reflex
	u := makeNative(v)
	t := u.Type()
	r := gate{
		Matter:   matter,
		Fixed:    make(map[circuit.Name]reflect.Value),
		Ellipses: u.MethodByName(cognizeEllipses),
	}

	// Build map of valves handled by dedicated methods.
	for i := 0; i < t.NumMethod(); i++ {
		n := t.Method(i).Name
		if len(n) >= len(cognizePrefix) && n[:len(cognizePrefix)] == cognizePrefix {
			r.Fixed[n[len(cognizePrefix):]] = u.MethodByName(n)
		}
	}

	// All dedicated valves need to be connected. All connected valves need to be handled (by dedicated or ellipses).

	// Verify all connected valves have dedicated handlers or there is a generic handler.
	var connected []circuit.Name
	ellipses := r.Ellipses.IsValid()
	for vlv, _ := range given {
		connected = append(connected, vlv)
		if ellipses {
			continue
		}
		if _, ok := r.Fixed[vlv]; !ok {
			log.Fatalf("%v gate does not handle connected valve (%v):\n%v\n", ellipses, vlv, matter)
		}
	}

	// Verify all dedicated valves are connected
	for vlv, _ := range r.Fixed {
		if _, ok := given[vlv]; !ok {
			log.Fatalf("gate valve (%v) must be connected:\n%v\n", vlv, matter)
		}
	}

	reflex, eye := NewEyeCognizer(r.Cognize, connected...)
	for vlv, antisyn := range given {
	}

	expected = nil
	obj = u.Interface()
	residue = obj.(Native).Spark(eye, matter, aux...)

	return
}

// gate is a materialized native reflex.
type gate struct {
	Matter   Circuit
	Fixed    map[circuit.Name]reflect.Value // valve name -> dedicated handler
	Ellipses reflect.Value                  // ellipses handler
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

	// Resolve handler
	handler, ell := g.Ellipses, true
	if _, ok := valve.(string); ok {
		if h, ok := g.Fixed[valve]; ok {
			handler, ell = h, false
		}
	}

	// Invoke handler
	if ell {
		handler.Call(
			[]reflect.Value{
				reflect.ValueOf(eye),
				reflect.ValueOf(valve),
				reflect.ValueOf(value),
			},
		)
	} else {
		handler.Call(
			[]reflect.Value{
				reflect.ValueOf(eye),
				reflect.ValueOf(value),
			},
		)
	}
}

// makeNative creates a copy of like.
// Pointer types allocate the object pointed to and copy that object as well.
func makeNative(like interface{}) reflect.Value {
	t := reflect.TypeOf(like)
	switch t.Kind() {
	case reflect.Ptr: // Pointer types are allocated
		return reflect.New(t.Elem()).Convert(t)
	default: // Value-based types are used as is
		return reflect.ValueOf(like)
	}
	panic(0)
}
