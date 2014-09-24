// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"
	"log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
	"github.com/gocircuit/escher/memory"
)

type Range struct{
	mem plumb.Given
	aux plumb.Given
}

func (h *Range) Spark() {
	h.mem.Init()
	h.aux.Init()
}

func (h *Range) CognizeMemory(eye *be.Eye, v interface{}) {
	h.mem.Fix(v)
}

func (h *Range) CognizeAux(eye *be.Eye, v interface{}) {
	h.aux.Fix(v)
}

func (h *Range) CognizeOverWith(eye *be.Eye, v interface{}) {
	eye.Show(
		DefaultValve, 
		rangeOverWith(
			h.mem.Use().(memory.Memory),
			h.aux.Use(),
			v.(Circuit).CircuitAt("Over"), // Circuit
			v.(Circuit).At("With"), // Materializable design (circuit or address)
		),
	)
}

func (h *Range) Cognize(*be.Eye, interface{}) {}

func rangeOverWith(mem memory.Memory, aux Value, over Circuit, with Value) Circuit {
	gates := over.Gates()
	ch := make(chan Circuit, len(gates))
	var i int
	for gname_, gvalue_ := range gates {
		gname, gvalue := gname_, gvalue_
		index := i
		i++
		go func() { // For each gate
			var x *be.Cell
			switch t := with.(type) {
			case Circuit:
				x = be.NewCell(be.Materialize(mem, t))
			case string:
				x = be.NewCell(be.Materialize(mem, NewAddressParse(t)))
			default:
				log.Fatalf("Unknown type at Range:With (%T)", with)
			}
			x.ReCognize(
				DefaultValve,
				New().
					Grow("Aux", aux).
					Grow("Name", gname).
					Grow("Value", gvalue).
					Grow("Count", len(gates)).
					Grow("Index", index),
			)
			vlv, val := x.Cognize() // Read the first output from the default gate of the cell
			if vlv != DefaultValve {
				panic(4)
			}
			ch <- val.(Circuit)
		}()
	}
	r := New()
	for _ = range gates { // merge all output circuits into one
		for n, v := range (<-ch).Gates() {
			r.ReGrow(n, v)
		}
	}
	return r
}
