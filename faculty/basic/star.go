// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"fmt"
	"log"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/be"
)

func init() {
	faculty.Register("Star", MaterializeStar)
}

// MaterializeStar
func MaterializeStar(matter *Matter) (Reflex, Value) {
	return MaterializeStarWithFunc(matter, nil)
}

func MaterializeShow(matter *Matter) (Reflex, Value) {
	return MaterializeStarWithFunc(
		matter, 
		func (name Name, v interface{}) {
			fmt.Printf(":%v = %v\n", name, v)
		},
	)
}

type JuncFunc func(Name, interface{})

func MaterializeStarWithFunc(matter *Matter, jf JuncFunc) (Reflex, Value) {
	if matter.View.Len() < 1 {
		log.Printf("Star not connected")
		return nil, nil
	}
	vlv := make([]Name, 0, matter.View.Len())
	for v, _ := range matter.View.Gate {
		vlv = append(vlv, v)
	}
	j := junction{jf, vlv}
	reflex, _ := NewEyeCognizer(j.Cognize, vlv...)
	return reflex, nil
}

type junction struct {
	f JuncFunc
	valve []Name
}

func (j junction) Cognize(eye *Eye, name Name, value interface{}) {
	if j.f != nil {
		j.f(name, value)
	}
	ch := make(sparkChan, len(j.valve)-1)
	for _, u_ := range j.valve {
		u := u_
		if u == name {
			continue
		}
		go spark(ch, eye, u, value)
	}
	for i := 0; i+1 < len(j.valve); i++ {
		<-ch
	}
}

type sparkChan chan struct{}

func spark(ch sparkChan, eye *Eye, dvalve Name, dvalue interface{}) {
	eye.Show(dvalve, dvalue)
	ch <- struct{}{}
}
