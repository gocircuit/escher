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
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&Star{}), "e", "Star")
	faculty.Register(be.NewMaterializer(&Star{}, StarFunc(show)), "e", "Show")
	faculty.Register(be.NewMaterializer(&Star{}, StarFunc(show1)), "e", "Show1")
	faculty.Register(be.NewMaterializer(&Star{}, StarFunc(show2)), "e", "Show2")
}

type StarFunc func(Name, interface{})

type Star struct {
	f    StarFunc
	view Circuit
}

func (s *Star) Spark(_ *be.Eye, matter Circuit, aux ...interface{}) Value {
	s.view = matter.CircuitAt("View")
	if len(aux) == 1 {
		s.f = aux[0].(StarFunc)
	}
	return nil
}

func (s *Star) OverCognize(eye *be.Eye, name Name, value interface{}) {
	if s.f != nil {
		s.f(name, value)
	}
	for gn_, _ := range s.view.Gate {
		gn := gn_
		if gn == name {
			continue
		}
		go eye.Show(gn, value)
	}
}

func show(name Name, v interface{}) {
	fmt.Printf("Showing:%v = %v\n", name, String(v))
}

func show1(name Name, v interface{}) {
	var w bytes.Buffer
	Print(&w, Format{"", "\t", 1}, v)
	fmt.Printf("Showing:%v = %v\n", name, w.String())
}

func show2(name Name, v interface{}) {
	var w bytes.Buffer
	Print(&w, Format{"", "\t", 2}, v)
	fmt.Printf("Showing:%v = %v\n", name, w.String())
}
