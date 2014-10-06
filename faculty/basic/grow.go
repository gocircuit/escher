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

type Grow struct{
	sync.Mutex
	u Circuit
}

func (g *Grow) Spark(*be.Matter, ...interface{}) Value {
	g.u = New()
	return &Grow{}
}

func (g *Grow) CognizeKey(eye *be.Eye, v interface{}) {
	g.Lock()
	defer g.Unlock()
	g.u.ReGrow("Key", v)
	g.fire(eye)
}

func (g *Grow) CognizeValue(eye *be.Eye, v interface{}) {
	g.Lock()
	defer g.Unlock()
	g.u.ReGrow("Value", v)
	g.fire(eye)
}

func (g *Grow) CognizeImg(eye *be.Eye, v interface{}) {
	g.Lock()
	defer g.Unlock()
	g.u.ReGrow("Img", v.(Circuit))
	g.fire(eye)
}

func (g *Grow) Cognize(eye *be.Eye, v interface{}) {}

func (g *Grow) fire(eye *be.Eye) {
	if g.u.Len() != 3 {
		return
	}
	eye.Show("", g.u.CircuitAt("Img").Clone().ReGrow(g.u.At("Key"), g.u.At("Value")))
}
