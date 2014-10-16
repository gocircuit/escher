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

type Lens struct{
	valve []Name
	sync.Mutex
	history []? Circuit // histories in both directions
}

// need deep circuit copy 

func (g *Lens) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	if len(matter.View.Gate) > 2 {
		panic("lens can have one or two endpoints")
	}
	for vlv, _ := range matter.View.Gate {
		g.valve = append(g.valve, vlv)
	}
	g.history = make([]Circuit, len(g.valve))
	for i, _ := range g.history {
		g.history[i] = New()
	}
	return ?? // 
}

func (g *Lens) OverCognize(eye *be.Eye, valve Name, value interface{}) {
	g.Lock()
	i := len(g.history.Gate)
	g.history.Gate[i] = Copy(value)
	g.Unlock()
	??
}
