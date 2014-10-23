// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("Star", be.NewNativeMaterializer(&Star{}))
	faculty.Register("Show", be.NewNativeMaterializer(&Star{}, StarFunc(show)))
	faculty.Register("Show1", be.NewNativeMaterializer(&Star{}, StarFunc(show1)))
	faculty.Register("Show2", be.NewNativeMaterializer(&Star{}, StarFunc(show2)))
}

type StarFunc func(Name, interface{})

type Star struct {
	f StarFunc
	view Circuit
}

func (s *Star) Spark(_ *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	s.view = matter.View
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
	fmt.Printf("Show:%v = %v\n", name, v)
}

func show1(name Name, v interface{}) {
	switch t := v.(type) {
	case Circuit:
		fmt.Printf("Show:%v = %s\n", name, t.Print("", "\t", 1))
	default:
		fmt.Printf("Show:%v = %v\n", name, v)
	}
}

func show2(name Name, v interface{}) {
	switch t := v.(type) {
	case Circuit:
		fmt.Printf("Show:%v = %s\n", name, t.Print("", "\t", 2))
	default:
		fmt.Printf("Show:%v = %v\n", name, v)
	}
}
