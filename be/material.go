// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"os"
	"reflect"

	"github.com/hoijui/escher/circuit"
	"github.com/hoijui/escher/kit/runtime"
)

const cognizePrefix = "Cognize"
const cognizeEllipses = "OverCognize"

// NewMaterializer returns a materializer that generates copies of sample and sparks them with the aux data.
func NewMaterializer(sample Material, aux ...interface{}) Materializer {
	return (&materializer{sample, aux}).Materialize
}

type materializer struct {
	sample Material
	aux    []interface{}
}

func (x *materializer) Materialize(given Reflex, matter circuit.Circuit) interface{} {
	return Materialize(given, matter, x.sample, x.aux...)
}

func (x *materializer) String() string {
	if ns, ok := x.sample.(interface {
		MaterialString(...interface{}) string
	}); ok {
		return ns.MaterialString(x.aux...)
	}
	return fmt.Sprintf("Material(%T)", x.sample)
}

// Materialize materializes the native implementation v.
// It returns the resulting reflex and residue, but not the Go-facing instance.
func Materialize(given Reflex, matter circuit.Circuit, v Material, aux ...interface{}) interface{} {
	residue, _ := MaterializeInstance(given, matter, v, aux...)
	return residue
}

// MaterializeInstance materializes the native implementation v.
// It returns the resulting reflex and residue, as well as the Go-facing instance.
func MaterializeInstance(given Reflex, matter circuit.Circuit, v Material, aux ...interface{}) (residue, obj interface{}) {
	g, r := buildReflex(matter, v)
	verify(matter, g, given)

	eye := NewEye(given)
	residue = r.(Material).Spark(eye, matter, aux...)
	eye.Connect(given, g.Cognize) // live connection must happen after Spark completes
	return residue, r
}

func buildReflex(matter circuit.Circuit, v Material) (g gate, receiver interface{}) {

	// Build gate reflex
	u := makeReflex(v)
	t := u.Type()
	g = gate{
		Matter:   matter,
		Fixed:    make(map[circuit.Name]reflect.Value),
		Ellipses: u.MethodByName(cognizeEllipses),
	}

	// Build map of valves handled by dedicated methods.
	for i := 0; i < t.NumMethod(); i++ {
		n := t.Method(i).Name
		if len(n) >= len(cognizePrefix) && n[:len(cognizePrefix)] == cognizePrefix {
			g.Fixed[n[len(cognizePrefix):]] = u.MethodByName(n)
		}
	}

	return g, u.Interface()
}

// makeReflex creates a copy of like.
// Pointer types allocate the object pointed to and copy that object as well.
func makeReflex(like interface{}) reflect.Value {
	t := reflect.TypeOf(like)
	switch t.Kind() {
	case reflect.Ptr: // Pointer types are allocated
		return reflect.New(t.Elem()).Convert(t)
	default: // Value-based types are used as is
		return reflect.ValueOf(like)
	}
	panic(0)
}

// Verify all dedicated valves are connected and all connected valves are handled (by dedicated or ellipses).
func verify(matter circuit.Circuit, r gate, given Reflex) {

	// Verify all connected valves have dedicated handlers or there is a generic handler.
	ellipses := r.Ellipses.IsValid()
	for vlv := range given {
		if ellipses {
			continue
		}
		if _, ok := r.Fixed[vlv]; !ok {
			Panicf("gate does not handle connected valve (%v)", vlv)
		}
	}

	// Verify all dedicated valves are connected
	for vlv := range r.Fixed {
		if _, ok := given[vlv]; !ok {
			Panicf("gate valve (%v) must be connected", vlv)
		}
	}
}

// gate is a materialized native reflex.
type gate struct {
	Matter   circuit.Circuit
	Fixed    map[circuit.Name]reflect.Value // valve name -> dedicated handler
	Ellipses reflect.Value                  // ellipses handler
}

func (g *gate) Cognize(eye *Eye, valve circuit.Name, value interface{}) {

	// Catch panics during cognizing and report their context to the user
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%v\n\n%v\n", r, PrintableMatter(g.Matter))
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
