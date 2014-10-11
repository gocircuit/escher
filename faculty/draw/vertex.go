// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package draw

import (
	// "log"
	"sync"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("draw.Vertex", be.NewGateMaterializer(&Vertex{}))
}

// Vertex…
type Vertex struct{
	n float64
	view Circuit
	eye *be.Eye
	sync.Mutex
	mass map[*Vertex]*mass
}

type mass struct {
	Residual float64
	Stationary float64
}

func (x *Vertex) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	x.n = float64(len(matter.Super.Design.(Circuit).Gate)) // number of vertices in enclosing circuit
	x.view = matter.View
	x.eye = eye
	x.mass = make(map[*Vertex]*mass)
	x.mass[x] = &mass{1, 0}
	go x.push(x)
	return nil
}

const teleport = 0.11

func (x *Vertex) pushHere(vertex *Vertex) float64 {
	x.Lock()
	defer x.Unlock()
	m := x.mass[vertex]
	if m.Residual <= 1 / (x.n * x.n) { // if error is small enough, no update necessary
		return 0
	}
	m.Stationary += teleport * m.Residual
	m.Residual = (1 - teleport) * m.Residual / 2 // lazy random walk
	return m.Residual
}

func (x *Vertex) push(vertex *Vertex) {
	amt := x.pushHere(vertex)
	if amt == 0 {
		return
	}
	d := float64(len(x.view.Gate))
	for nbr, _ := range x.view.Gate {
		x.eye.Show(
			nbr, 
			New().
				Grow("Vertex", vertex).
				Grow("Mass", amt / d),
		)
	}
}

func (x *Vertex) OverCognize(_ *be.Eye, _ Name, val interface{}) {
	vertex := val.(Circuit).At("Vertex").(*Vertex)
	x.Lock()
	defer x.Unlock()
	m, ok := x.mass[vertex]
	if !ok {
		m = &mass{0, 0}
		x.mass[vertex] = m
	}
	m.Residual += val.(Circuit).FloatAt("Mass")
	go x.push(vertex)
}
