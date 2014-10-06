// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"container/list"
	// "fmt"
	"log"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/kit/memory"
)

/*
	Hamiltonian traverses the hierarchy of circuits induced by a given top-level/valveless circuit.

	Start = {
		Circuit Circuit
		Vector Vector
	}

	Memory Memory

	View = {
		Circuit Circuit // Current circuit in the exploration sequence
		Vector Vector
		Index int // Index of this circuit within exploration sequence, 0-based
		Depth int
		Dir string
		Path string // Loop
	}
*/
type Hamiltonian struct{
	mem plumb.Given
}

func (h *Hamiltonian) Spark(*be.Matter, ...interface{}) Value {
	h.mem.Init()
	return &Hamiltonian{}
}

func (h *Hamiltonian) CognizeMemory(_ *be.Eye, v interface{}) {
	h.mem.Fix(v)
}

func (h *Hamiltonian) CognizeView(*be.Eye, interface{}) {}

func (h *Hamiltonian) CognizeStart(eye *be.Eye, dv interface{}) {
	var in = dv.(Circuit)
	var start = hamiltonianView{
		Circuit: in.CircuitAt("Circuit"),
		Vector: Vector(in.CircuitAt("Vector")),
		Index: 0,
		Depth: 0,
	}
	var v = start
	var memory list.List
	for {
		eye.Show("View", v.Circuitize()) // yield current hamiltonianView

		vg, vv := v.Vector.Reduce()
		if vg == Super {
			e := memory.Front() // backtrack
			if e == nil {
				log.Fatalf("short memory")
			}
			u := e.Value.(hamiltonianView)
			memory.Remove(e)
			//
			v.Circuit = u.Circuit
			ug, _ := u.Vector.Reduce()
			v.Vector = v.Circuit.Follow(NewVector(ug, vv))
			v.Depth--
		} else {
			t, ok := v.Circuit.At(vg).(Address)
			if !ok {
				log.Fatalf("unknown gate meaning %T", t)
			}
			if memory.Len() > 100 {
				log.Fatalf("memory overload")
				// memory.Remove(memory.Front())
			}
			memory.PushFront(v) // remember
			//
			lookup := h.mem.Use().(Memory).Lookup(t)
			if lookup == nil {
				log.Fatalf("No Hamiltonian circuit addressed %s", t.String())
			}
			v.Circuit = lookup.(Circuit) // transition to next circuit
			v.Vector = v.Circuit.Follow(NewVector(Super, vv))
			v.Depth++
		}
		v.Index++
		//
		// log.Printf("%s vs %s = %v", v.Vector, start.Vector, Same(v.Vector, start.Vector))
		if Same(v.Circuit, start.Circuit) && Same(v.Vector, start.Vector) {
			eye.Show("View", v.Circuitize().Grow("Path", "Loop")) // yield current hamiltonianView
			return
		}
	}
}

type hamiltonianView struct {
	Circuit Circuit
	Vector Vector
	Index int
	Depth int
}

func (v *hamiltonianView) Dir() string {
	if vg, _ := v.Vector.Reduce(); vg == Super {
		return "Up"
	}
	return "Down"
}

func (v *hamiltonianView) Circuitize() Circuit {
	return New().
		Grow("Circuit", v.Circuit).
		Grow("Vector", Circuit(v.Vector)).
		Grow("Index", v.Index).
		Grow("Depth", v.Depth).
		Grow("Dir", v.Dir())
}
