// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"log"

	"github.com/gocircuit/escher/a"
	cir "github.com/gocircuit/escher/circuit"
)

func ParseVerb(src string) (verb cir.Verb) {
	defer func() {
		if r := recover(); r != nil {
			verb = cir.Verb{}
		}
	}()
	t := a.NewSrcString(src)
	verb = cir.Verb(SeeVerb(t).(cir.Circuit))
	if t.Len() != 0 {
		log.Printf("Non-address characters at end of %q", src)
		panic(1)
	}
	return verb
}

func Parse(src string) (cir.Name, cir.Value) {
	return SeePeer(a.NewSrcString(src))
}

func ParseCircuit(src string) cir.Circuit {
	n, v := Parse(src)
	if _, ok := n.(Nameless); !ok {
		panic("not a circuit")
	}
	return v.(cir.Circuit)
}
