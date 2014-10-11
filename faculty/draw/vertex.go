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
	mass map[*Vertex]float64
}

func (x *Vertex) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	x.n = float64(len(matter.Super.Design.(Circuit).Gate)) // number of vertices in enclosing circuit
	x.view = matter.View
	x.eye = eye
	x.mass = make(map[*Vertex]float64)
	x.mass[x] = 1
	go x.push(x)
	return nil
}

const teleport = 0.11

func (x *Vertex) push(vertex *Vertex) {
	x.Lock()
	amt := x.mass[vertex]
	if amt <= 1 / (x.n * x.n) { // if error is small enough
		amt = 0
	} else {
		x.mass[vertex] = amt / 2 // lazy random walk
		amt = amt / 2
	}
	x.Unlock()
	if amt == 0 {
		return
	}
	for nbr, _ := range x.view.Gate {
		p := (1 - teleport) * amt / float64(len(x.view.Gate))
		x.eye.Show(
			nbr, 
			New().
				Grow("Vertex", vertex).
				Grow("Mass", p),
		)
	}
	vertex.OverCognize(
		nil, nil, 
		New().
			Grow("Vertex", vertex).
			Grow("Mass", teleport * amt),
	)
}

func (x *Vertex) OverCognize(_ *be.Eye, _ Name, val interface{}) {
	vertex := val.(Circuit).At("Vertex").(*Vertex)
	x.Lock()
	u, ok := x.mass[vertex]
	if !ok {
		u = 0
	}
	x.mass[vertex] = u + val.(Circuit).FloatAt("Mass")
	x.Unlock()
	go x.push(vertex)
}
