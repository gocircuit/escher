// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"
	// "log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Flatten
//
// On Before, Flatten receives:
//	{
//		Idiom Circuit	// namespace of values
//		Value Circuit	// value to materialize
//	}
//
// On After, Flatten sends:
//	{
//		Idiom Circuit
//		Root Circuit
//		Plane Circuit
//		Residue Circuit
//	}
//
type Flatten struct{}

func (Flatten) Spark(eye *be.Eye, _ *be.Matter, _ ...interface{}) Value {
	return nil
}

func (Flatten) CognizeBefore(eye *be.Eye, value interface{}) {
	v := value.(Circuit)
	idiom := v.CircuitAt("Idiom")
	value := v.CircuitAt("Value")

	x := newFlattenTransformer(idiom)
	_, residue := be.MaterializeTransform(idiom, value, x)

	??

	eye.Show(
		"After", 
		New().
			Grow("Idiom", x.idiom). // original idiom
			Grow("Value", x.value). // original value
			Grow("Root", x.root). // materialization path -> plane gate name
			Grow("Residue", residue), // materialization path -> source
			Grow("Plane", x.plane). // gate name = index, gate value = source
	)
}

func (Flatten) CognizeAfter(eye *be.Eye, v interface{}) {}


// flattenTransformer
type flattenTransformer struct {
	idiom Circuit
	sync.Mutex
	root, plane Circuit
}

func newFlattenTransformer(idiom Circuit) *flattenTransformer {
	return &flattenTransformer{
		idiom: idiom,
		root: New(),
		plane: New(),
	}
}

func (x *flattenTransformer) Transform(addr Address, value Value) Value {
	return flattenMaterializer{x, addr}
}

func (x *flattenTransformer) Link() ?? {
}

// flattenMaterializer
type flattenMaterializer struct {
	transformer *flattenTransformer
	source Value
}

func (m flattenMaterializer) Materialize(matter *Matter) (Reflex, Value) {
	??
	reflex, _ := NewIdleMaterializer(matter, m.source)
	return reflex, m.source
}

// 