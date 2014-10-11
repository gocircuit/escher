// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package draw

import (
	"io"
	"io/ioutil"
	// "log"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("draw.Vertex", be.NewGateMaterializer(&Vertex{}))
}

// Vertex…
type Vertex struct{
	mass map[Name]float64
}

func (x *Vertex) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	x.mass = make(map[Name]float64)
	go func() {
		??
	}()
	return Vertex{}
}

func (x *Vertex) OverCognize(eye *be.Eye, v interface{}) {
	??
}
