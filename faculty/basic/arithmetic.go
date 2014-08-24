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
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	faculty.Root.AddTerminal("Sum", Sum{})
	// faculty.Root.AddTerminal("Prod", Prod{})
}

// Sum
type Sum struct{}

func (Sum) Materialize() think.Reflex {
	reflex, eye := plumb.NewEye("X", "Y", "Sum")
	go func() {
		lit := Make() // literals
		for {
			dvalve, dvalue := eye.See()
			lit[dvalve] = dvalue.(int)
			//
			var wg sync.WaitGroup
			wg.Add(2)
			switch dvalve {
			case "X":
				go func() {
					eye.Show("Y", lit.Int("Sum") - lit.Int("X"))
					wg.Done()
				}()
				go func() {
					eye.Show("Sum", lit.Int("Y") + lit.Int("X"))
					wg.Done()
				}()
			case "Y":
				go func() {
					eye.Show("X", lit.Int("Sum") - lit.Int("Y"))
					wg.Done()
				}()
				go func() {
					eye.Show("Sum", lit.Int("Y") + lit.Int("X"))
					wg.Done()
				}()
			case "Sum":
				go func() {
					eye.Show("X", lit.Int("Sum") - lit.Int("Y"))
					wg.Done()
				}()
				go func() {
					eye.Show("Y", lit.Int("Sum") - lit.Int("X"))
					wg.Done()
				}()
			}
			wg.Wait()
		}
	}()
	return reflex
}
