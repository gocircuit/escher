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
	faculty.Register("draw.Circle", be.NewNativeMaterializer(&Circuit{}))
}

// Circuit…
type Circuit struct{}

func (x *Circuit) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (x *Circuit) CognizeCircuit(_ *be.Eye, _ Name, val interface{}) {
	v := val.(Circuit)
	r := New()
	// colors
	var color = []string{"white", "black"}
	// add circles for gates
	var i int
	for _, name := range v.SortedLetters() {
		r.Grow(i, New().
			Grow("cx", ?).
			Grow("cy", ?).
			Grow("r", ?).
			Grow("fill", ?).
		)
		i++
	}
	for _, name := range v.SortedLetters() {
		??
	}

	eye.Show(DefaultValve, r)
}

func (x *Circuit) Cognize(_ *be.Eye, _ Name, val interface{}) {}
