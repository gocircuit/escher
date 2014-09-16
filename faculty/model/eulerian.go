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

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

/*
	Eulerian traverses the hierarchy of circuits induced by a given top-level/valveless circuit.

	Start Circuit

	Memory *Memory

	View = {
		Circuit Circuit // Current circuit in the exploration sequence
		Index int // Index of this circuit within exploration sequence, 0-based
		Depth int
		Dir string
		Path string // Loop
	}
*/
type Eulerian struct{
	m chan *Memory
}

func (e *Eulerian) Spark() {
	e.m = make(chan *Memory, 1)
}

func (e *Eulerian) memory() *Memory {
	m := <-e.m
	e.m <- m
	return m
}

func (e *Eulerian) CognizeMemory(_ *be.Eye, v interface{}) {
	e.m <- v.(*Memory)
}

func (e *Eulerian) CognizeView(*be.Eye, interface{}) {}

func (e *Eulerian) CognizeStart(eye *be.Eye, v interface{}) {
	euler(
		eye,
		e.memory(),
		&eulerView{
			Circuit: v.(Circuit),
			Index: 0,
			Depth: 0,
		},
	)
}

func euler(eye *be.Eye, m *Memory, v *eulerView) int {
	var n int // number of views shown
	eye.Show("View", v.Circuitize(true))
	n++
	//
	for g, h := range v.Circuit.Gates() {
		switch t := h.(type) {
		case Address:
			x := m.Lookup(t.Path()...) // Resolve addresses once
			if x == nil {
				log.Fatalf("No Eulerian circuit addressed %s", t.String())
			}
			u, ok := x.(Circuit)
			if !ok {
				break // cannot enter non-circuits
			}
			n += euler(
				eye, 
				m, 
				&eulerView{
					Circuit: u,
					Index: v.Index + n,
					Depth: v.Depth + 1,
				},
			)
		case Circuit:
			n += euler(
				eye, 
				m, 
				&eulerView{
					Circuit: t,
					Index: v.Index + n,
					Depth: v.Depth + 1,
				},
			)
		default: // skip non-address gates as data
		}
	}
	v.Index += n
	eye.Show("View", v.Circuitize(false))
	n++
	return n
}

type eulerView struct {
	Circuit Circuit
	Index int
	Depth int
}

func (v *eulerView) Circuitize(entering bool) Circuit {
	var dir = "Return"
	if entering {
		dir = "Enter"
	}
	var p = "Within"
	if !entering && v.Depth == 0 {
		p = "Loop"
	}
	return New().
		Grow("Circuit", v.Circuit).
		Grow("Index", v.Index).
		Grow("Depth", v.Depth).
		Grow("Dir", dir).
		Grow("Path", p)
}
