// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	"log"
	"sync"

	. "github.com/gocircuit/escher/circuit"
)

func MaterializeUnion(matter *Matter) (Reflex, Value) {
	return MaterializeNative(matter, &Union{})
}

type Union struct {
	field []Name
	sync.Mutex
	union Circuit
}

func (u *Union) Spark(eye *Eye, matter *Matter, aux ...interface{}) Value {
	var defaultConnected bool
	for vlv, _ := range matter.View.Gate {
		if vlv == DefaultValve {
			defaultConnected = true
		} else {
			u.field = append(u.field, vlv)
		}
	}
	if !defaultConnected || len(u.field) == 0 {
		log.Fatalf("Fork gate's default valve not linked or has no partition valves. In:\n%v\n", matter.Super.Design.(Circuit))
	}
	//
	u.union = New()
	return nil
}

func (u *Union) Cognize(eye *Eye, value interface{}) {
	// log.Printf("%p u: %v", u, value)
	y := make(chan struct{}) // block and
	for _, f_ := range u.field { // send updated conjunction to all field valves
		f := f_
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Fatalf("Union over %v panic on %v: %v", u.field, value, r)
				}
			}()
			eye.Show(f, value.(Circuit).At(f))
			y <- struct{}{}
		}()
	}
	for _ = range u.field {
		<-y
	}
}

func (u *Union) OverCognize(eye *Eye, valve Name, value interface{}) {
	// log.Printf("%p u:%v %v", u, valve, value)
	u.Lock()
	defer u.Unlock()
	if valve == DefaultValve {
		panic(1)
	}
	u.union.Include(valve, value)
	if u.union.Len() == len(u.field) {
		w := u.union.Copy()
		eye.Show(DefaultValve, w)
	}
}
