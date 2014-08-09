// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	// println("Loading basic faculty")
	faculty.Root.AddTerminal("sum", Sum{})
	faculty.Root.AddTerminal("prod", Prod{})
}

// Sum
type Sum struct{}

func (Sum) Materialize() think.Reflex {
	reflex, eye := faculty.NewEye("X", "Y", "Sum")
	go func() {
		x := &sum{
			ready: make(chan struct{}),
		}
		x.reply = eye.Focus(x.ShortCognize)
		close(x.ready)
	}()
	return reflex
}

type sum struct {
	ready chan struct{}
	reply *faculty.EyeReCognizer
}

func (s *sum) ShortCognize(imp faculty.Impression) {
	// println(fmt.Sprintf("imp=%v", imp))
	<-s.ready
	x, xk := imp.Valve("X").Value().(int)
	y, yk := imp.Valve("Y").Value().(int)
	su, sk := imp.Valve("Sum").Value().(int)
	switch imp.Index(2).Valve() { // determine which valve we are computing for
	case "X":
		if !sk || !yk {
			return
		}
		s.reply.ReCognize(faculty.MakeImpression().Show(0, "X", su - y))
	case "Y":
		if !sk || !xk {
			return
		}
		s.reply.ReCognize(faculty.MakeImpression().Show(0, "Y", su - x))
	case "Sum":
		if !xk || !yk {
			return
		}
		s.reply.ReCognize(faculty.MakeImpression().Show(0, "Sum", x + y))
	default:
		panic(7)
	}
}

// Prod
type Prod struct{}

func (Prod) Materialize() think.Reflex {
	reflex, eye := faculty.NewEye("X", "Y", "Prod")
	go func() {
		x := &prod{
			ready: make(chan struct{}),
		}
		x.reply = eye.Focus(x.ShortCognize)
		close(x.ready)
	}()
	return reflex
}

type prod struct {
	ready chan struct{}
	reply *faculty.EyeReCognizer
}

func (s *prod) ShortCognize(imp faculty.Impression) {
	// println(fmt.Sprintf("imp=%v", imp))
	<-s.ready
	x, xk := imp.Valve("X").Value().(int)
	y, yk := imp.Valve("Y").Value().(int)
	pr, pk := imp.Valve("Prod").Value().(int)
	switch imp.Index(2).Valve() { // determine which valve we are computing for
	case "X":
		if !pk || !yk || y == 0 {
			return
		}
		s.reply.ReCognize(faculty.MakeImpression().Show(0, "X", pr / y))
	case "Y":
		if !pk || !xk || x == 0 {
			return
		}
		s.reply.ReCognize(faculty.MakeImpression().Show(0, "Y", pr / x))
	case "Prod":
		if !xk || !yk {
			return
		}
		s.reply.ReCognize(faculty.MakeImpression().Show(0, "Prod", x * y))
	default:
		panic(7)
	}
}
