// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	"sync"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
)

func init() {
	faculty.Register("Sum", Sum{})
	// faculty.Root.Grow("Prod", Prod{})
}

// Sum
type Sum struct{}

func (Sum) Materialize() be.Reflex {
	x := &sum{lit: New()}
	reflex, _ := be.NewEyeCognizer(x.Cognize, "X", "Y", "Sum")
	return reflex
}

type sum struct {
	sync.Mutex
	lit Circuit // literals
}

func (x *sum) save(valve string, value int) {
	x.Lock()
	defer x.Unlock()
	x.lit.Include(valve, value)
}

func (x *sum) u(valve string) int {
	x.Lock()
	defer x.Unlock()
	return x.lit.IntOrZeroAt(valve)
}

func (x *sum) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	x.save(dvalve, plumb.AsInt(dvalue))
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(2)
	switch dvalve {
	case "X":
		go func() { // Cognize
			defer func() {
				recover()
			}()
			defer wg.Done()
			eye.Show("Y", x.u("Sum") - x.u("X"))
		}()
		go func() {
			defer func() {
				recover()
			}()
			defer wg.Done()
			eye.Show("Sum", x.u("Y") + x.u("X"))
		}()
	case "Y":
		go func() {
			defer func() {
				recover()
			}()
			defer wg.Done()
			eye.Show("X", x.u("Sum") - x.u("Y"))
		}()
		go func() {
			defer func() {
				recover()
			}()
			defer wg.Done()
			eye.Show("Sum", x.u("Y") + x.u("X"))
		}()
	case "Sum":
		go func() {
			defer func() {
				recover()
			}()
			defer wg.Done()
			eye.Show("X", x.u("Sum") - x.u("Y"))
		}()
		go func() {
			defer func() {
				recover()
			}()
			defer wg.Done()
			eye.Show("Y", x.u("Sum") - x.u("X"))
		}()
	}
}
