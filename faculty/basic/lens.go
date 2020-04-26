// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"sync"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
)

type Lens struct {
	valve []cir.Name
	sync.Mutex
	history cir.Circuit // histories from both valves { ValveOne { … }, ValveTwo { … } }
}

func (g *Lens) Spark(eye *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	mvg := matter.CircuitAt("View").Gate
	if len(mvg) < 1 || len(mvg) > 2 {
		panic("lens can have one or two endpoints")
	}
	g.history = cir.New()
	for vlv := range mvg {
		g.valve = append(g.valve, vlv)
		g.history.Grow(vlv, cir.New())
	}
	return g // return self in residual to expose query interface
}

func (g *Lens) OverCognize(eye *be.Eye, valve cir.Name, value interface{}) {
	g.remember(valve, value)
	for _, v := range g.valve {
		if v != valve {
			eye.Show(v, value)
		}
	}
}

func (g *Lens) remember(valve cir.Name, value cir.Value) {
	g.Lock()
	defer g.Unlock()
	h := g.history.CircuitAt(valve) // valve history circuit
	h.Grow(h.Len(), cir.DeepCopy(value))
}

func (g *Lens) Peek() cir.Circuit {
	g.Lock()
	defer g.Unlock()
	return cir.DeepCopy(g.history).(cir.Circuit)
}
