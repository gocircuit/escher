// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
	"github.com/gocircuit/escher/memory"
)

type Range struct{
	mem plumb.Given
}

func (h *Range) Spark() {
	h.mem.Init()
}

func (h *Range) CognizeMemory(eye *be.Eye, v interface{}) {
	h.mem.Fix(v.(*memory.Memory))
}

func (h *Range) CognizeIn(eye *be.Eye, v interface{}) {
	eye.Show(
		"_", 
		rangeOverWith(
			h.mem.Use().(*memory.Memory),
			v.(Circuit).CircuitAt("Over"), // Circuit
			v.(Circuit).At("With"), // Materializable design
		),
	)
}

func (h *Range) Cognize_(*be.Eye, interface{}) {}

func rangeOverWith(mem *memory.Memory, over Circuit, with Meaning) Circuit {
	gates := over.Gates()
	ch := make(chan Circuit, len(gates))
	for gname_, gvalue_ := range gates {
		gname, gvalue := gname_, gvalue_
		go func() {
			x := be.NewCell(be.Materialize(mem, with))
			x.ReCognize("_", New().Grow(gname, gvalue))
			vlv, val := x.Cognize()
			if vlv != "_" {
				panic(4)
			}
			ch <- val.(Circuit)
		}()
	}
	r := New()
	for _ = range gates {
		for n, v := range (<-ch).Gates() {
			r.ReGrow(n, v)
		}
	}
	return r
}
