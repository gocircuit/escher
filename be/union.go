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

type Union struct {
	field []Name
	flow  map[Name]chan struct{}
	sync.Mutex
	union Circuit
}

func (u *Union) Spark(eye *Eye, matter Circuit, aux ...interface{}) Value {
	// check whether default valve is connected and extract names of connected non-default valves
	var defaultConnected bool
	for vlv, _ := range matter.CircuitAt("View").Gate {
		if vlv == DefaultValve {
			defaultConnected = true
		} else {
			u.field = append(u.field, vlv)
		}
	}
	if !defaultConnected || len(u.field) == 0 {
		log.Fatalf("Fork gate's default valve not linked or has no partition valves. In:\n%v\n",
			matter.CircuitAt("Super").CircuitAt("Design"),
		)
	}
	// allocate flow control channels
	u.flow = make(map[Name]chan struct{})
	for _, f := range u.field {
		u.flow[f] = make(chan struct{}, 1)
		u.flow[f] <- struct{}{} // send initial flow tokens
	}
	//
	u.union = New()
	return nil
}

func (u *Union) Cognize(eye *Eye, value interface{}) {
	y := make(chan struct{})     // block and
	for _, f_ := range u.field { // send updated values to all field valves
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
	<-u.flow[valve] // obtain flow token
	u.Lock()
	defer u.Unlock()
	u.union.Grow(valve, value)         // grow will panic, if gate already exists
	if u.union.Len() == len(u.field) { // flush if all the fields have been set
		w := u.union
		u.union = New() // flush
		for f, _ := range u.flow {
			u.flow[f] <- struct{}{} // replenish flow tokens
		}
		eye.Show(DefaultValve, w)
	}
}
