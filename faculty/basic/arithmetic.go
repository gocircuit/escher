// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	"strconv"
	"sync"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	faculty.Register("Sum", Sum{})
	faculty.Register("IntString", be.NewNativeMaterializer(IntString{}))
	// faculty.Root.Grow("Prod", Prod{})
}

// IntString
type IntString struct{}

func (IntString) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return IntString{}
}

func (IntString) CognizeInt(eye *be.Eye, v interface{}) {
	eye.Show("String", strconv.Itoa(v.(int)))
}

func (IntString) CognizeString(eye *be.Eye, v interface{}) {
	i, err := strconv.Atoi(v.(string))
	if err != nil {
		panic(err)
	}
	eye.Show("Int", i)
}

// Sum
type Sum struct{}

func (Sum) Materialize() (be.Reflex, Value) {
	x := &sum{lit: New()}
	reflex, _ := be.NewEyeCognizer(x.Cognize, "X", "Y", "Sum")
	return reflex, Sum{}
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

func (x *sum) Cognize(eye *be.Eye, dv Name, dvalue interface{}) {
	dvalve := dv.(string)
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
