// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"log"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Fork struct{}

func (Fork) Materialize(matter *be.Matter) (be.Reflex, Value) {
	var arm []string
	var defaultUsed bool
	for vlv, _ := range matter.Valve {
		if vlv == "" { // 
			defaultUsed = true
		} else {
			arm = append(arm, vlv.(string))
		}
	}
	if !defaultUsed || len(arm) == 0 {
		log.Fatalf("Fork gate's default valve not linked or has no partition valves. In:\n%v\n", matter.Super.Design.(Circuit))
	}
	return be.MaterializeUnion(arm...)
}
