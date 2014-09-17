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

func (h *Embody) Cognize_(*be.Eye, interface{}) {}

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
	eye.Show("_", New().Grow("When", w).Grow("Cell", cell))
}
