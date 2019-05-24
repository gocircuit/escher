// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"bytes"
	"fmt"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&Star{}), "e", "Star")
	faculty.Register(be.NewMaterializer(&Star{}, StarFunc(show)), "e", "Show")
	faculty.Register(be.NewMaterializer(&Star{}, StarFunc(show1)), "e", "Show1")
	faculty.Register(be.NewMaterializer(&Star{}, StarFunc(show2)), "e", "Show2")
}

type StarFunc func(cir.Name, interface{})

type Star struct {
	f    StarFunc
	view cir.Circuit
}

func (s *Star) Spark(_ *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	s.view = matter.CircuitAt("View")
	if len(aux) == 1 {
		s.f = aux[0].(StarFunc)
	}
	return nil
}

func (s *Star) OverCognize(eye *be.Eye, name cir.Name, value interface{}) {
	if s.f != nil {
		s.f(name, value)
	}
	for gn_ := range s.view.Gate {
		gn := gn_
		if gn == name {
			continue
		}
		go eye.Show(gn, value)
	}
}

func show(name cir.Name, v interface{}) {
	fmt.Printf("Showing:%v = %v\n", name, cir.String(v))
}

func show1(name cir.Name, v interface{}) {
	var w bytes.Buffer
	cir.Print(&w, cir.Format{"", "\t", 1}, v)
	fmt.Printf("Showing:%v = %v\n", name, w.String())
}

func show2(name cir.Name, v interface{}) {
	var w bytes.Buffer
	cir.Print(&w, cir.Format{"", "\t", 2}, v)
	fmt.Printf("Showing:%v = %v\n", name, w.String())
}
