// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
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
	reflex, eye := faculty.NewEye("X", "Y", "Sum")
	go func() {
		x := &sum{
			connected: make(chan struct{}),
		}
		x.reply = eye.Focus(x.ShortCognize)
		close(x.connected)
	}()
	return reflex
}

type sum struct {
	connected chan struct{}
	reply     *faculty.EyeNerve
}

func (s *sum) ShortCognize(imp faculty.Impression) {
	// println(fmt.Sprintf("summing (%v)", Linearize(imp.Print("", " "))))
	<-s.connected
	x, xk := plumb.OptionallyInt(imp.Valve("X").Value())
	y, yk := plumb.OptionallyInt(imp.Valve("Y").Value())
	su, sk := plumb.OptionallyInt(imp.Valve("Sum").Value())
	// println(fmt.Sprintf("SUMMING X=%v/%T Y=%v/%T Sum=%v/%T", x, x, y, y, su, su))
	switch imp.Index(0).Valve() { // determine which valve was most recently updated
	case "X":
		if !sk || !yk {
			return
		}
		z := faculty.MakeImpression().Show(0, "X", x).Show(1, "Y", su-x).Show(2, "Sum", x+y)
		go s.reply.ReCognize(z)
	case "Y":
		if !sk || !xk {
			return
		}
		z := faculty.MakeImpression().Show(0, "Y", y).Show(1, "Sum", x+y).Show(2, "X", su-y)
		go s.reply.ReCognize(z)
	case "Sum":
		if !xk || !yk {
			return
		}
		z := faculty.MakeImpression().Show(0, "Sum", su).Show(1, "X", su-y).Show(2, "Y", su-x)
		go s.reply.ReCognize(z)
	default:
		panic(7)
	}
}
