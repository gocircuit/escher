// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Root.AddTerminal("Sum", Sum{})
	faculty.Root.AddTerminal("Prod", Prod{})
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
	reply *faculty.EyeReCognizer
}

func (s *sum) ShortCognize(imp faculty.Impression) {
	println(fmt.Sprintf("summing %v", imp.Print("", "   ")))
	<-s.connected
	x, xk := AsInt(imp.Valve("X").Value())
	y, yk := AsInt(imp.Valve("Y").Value())
	su, sk := AsInt(imp.Valve("Sum").Value())
	switch imp.Index(2).Valve() { // determine which valve we are computing for
	case "X":
		println("<-------X")
		if !sk || !yk {
			return
		}
		go s.reply.ReCognize(faculty.MakeImpression().Show(0, "X", su - y))
	case "Y":
		println("<-------Y")
		if !sk || !xk {
			return
		}
		go s.reply.ReCognize(faculty.MakeImpression().Show(0, "Y", su - x))
	case "Sum":
		println("<-------SUM")
		if !xk || !yk {
			println("SUM SKIP")
			return
		}
		go s.reply.ReCognize(faculty.MakeImpression().Show(0, "Sum", x + y))
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
			connected: make(chan struct{}),
		}
		x.reply = eye.Focus(x.ShortCognize)
		close(x.connected)
	}()
	return reflex
}

type prod struct {
	connected chan struct{}
	reply *faculty.EyeReCognizer
}

func (s *prod) ShortCognize(imp faculty.Impression) {
	// println(fmt.Sprintf("imp=%v", imp))
	<-s.connected
	x, xk := AsInt(imp.Valve("X").Value())
	y, yk := AsInt(imp.Valve("Y").Value())
	pr, pk := AsInt(imp.Valve("Prod").Value())
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
