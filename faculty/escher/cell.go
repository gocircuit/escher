// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
	"github.com/gocircuit/escher/memory"
)

/*
	Memory *Memory
*/
type Embody struct{
	mem plumb.Given
	design plumb.Given
}

func (h *Embody) Spark() {
	h.mem.Init()
	h.design.Init()
}

func (h *Embody) CognizeMemory(_ *be.Eye, v interface{}) {
	h.mem.Fix(v)
}

func (h *Embody) Cognize(*be.Eye, interface{}) {} // DefaultValve

func (h *Embody) CognizeDesign(_ *be.Eye, v interface{}) {
	h.design.Fix(v)
}

func (h *Embody) CognizeWhen(eye *be.Eye, w interface{}) {
	cell := be.NewCell(
		be.Materialize(
			h.mem.Use().(*memory.Memory), 
			h.design.Use().(Meaning),
		),
	)
	eye.Show(DefaultValve, New().Grow("When", w).Grow("Cell", cell))
}

// Connect
type Connect struct{
	cell plumb.Given
}

func (h *Connect) Spark() {
	h.cell.Init()
}

func (h *Connect) CognizeCell(_ *be.Eye, v interface{}) {
	h.cell.Fix(v)
}

func (h *Connect) CognizePush(_ *be.Eye, v interface{}) {
	w := v.(Circuit)
	cell := h.cell.Use().(*be.Cell)
	cell.ReCognize(w.StringAt("Valve"), w.At("Value"))
}

func (h *Connect) CognizePull(eye *be.Eye, v interface{}) {
	cell := h.cell.Use().(*be.Cell)
	vlv, val := cell.Cognize()
	eye.Show("_", New().Grow("Valve", vlv).Grow("Value", val).Grow("Pull", v))
}

func (h *Connect) Cognize_(*be.Eye, interface{}) {}
