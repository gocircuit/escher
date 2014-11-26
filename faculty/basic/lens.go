// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Lens struct {
	valve []Name
	sync.Mutex
	history Circuit // histories from both valves { ValveOne { … }, ValveTwo { … } }
}

func (g *Lens) Spark(eye *be.Eye, matter Circuit, aux ...interface{}) Value {
	mvg := matter.CircuitAt("View").Gate
	if len(mvg) < 1 || len(mvg) > 2 {
		panic("lens can have one or two endpoints")
	}
	g.history = New()
	for vlv, _ := range mvg {
		g.valve = append(g.valve, vlv)
		g.history.Grow(vlv, New())
	}
	return g // return self in residual to expose query interface
}

func (g *Lens) OverCognize(eye *be.Eye, valve Name, value interface{}) {
	g.remember(valve, value)
	for _, v := range g.valve {
		if v != valve {
			eye.Show(v, value)
		}
	}
}

func (g *Lens) remember(valve Name, value Value) {
	g.Lock()
	defer g.Unlock()
	h := g.history.CircuitAt(valve) // valve history circuit
	h.Grow(h.Len(), DeepCopy(value))
}

func (g *Lens) Peek() Circuit {
	g.Lock()
	defer g.Unlock()
	return DeepCopy(g.history).(Circuit)
}
