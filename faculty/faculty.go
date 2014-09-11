// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	// "fmt"
	"log"
	"strings"

	. "github.com/gocircuit/escher/circuit"
)

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

// Faculty is a node in a hierarchy of nodes that can hold subnodes as well as circuit designs (themselves circuit structures).
type Faculty Circuit

func NewFaculty(name string) Faculty {
	fty := Faculty(New())
	Circuit(fty).Seal(name)
	Circuit(fty).Include(Genus_{}, NewFacultyGenus())
	return fty
}

func (fty Faculty) Genus() *FacultyGenus {
	g, _ := Circuit(fty).At(Genus_{})
	return g.(*FacultyGenus)
}

func (fty Faculty) Exclude(name string) (forgotten Meaning) {
	return Circuit(fty).Exclude(name)
}

// Roam traverses the hierarchy, creating faculty nodes if necessary, returning the final two nodes.
func (fty Faculty) Roam(walk ...string) (parent, child Meaning) {
	if len(walk) == 0 {
		return nil, fty
	}
	if parent, child = fty.Lookup(walk[0]); parent == nil && child == nil { // If no child, make it
		child = fty.Refine(walk[0])
	}
	fac, ok := child.(Faculty)
	if !ok {
		panic("walking thru a non-faculty")
	}
	return fac.Roam(walk[1:]...)
}

func (fty Faculty) LookupAddress(addr string) (parent, child Meaning) {
	return fty.Lookup(strings.Split(addr, ".")...)
}

// Lookup ...
func (fty Faculty) Lookup(walk ...string) (parent, child Meaning) {
	if len(walk) == 0 {
		return nil, fty
	}

	v, ok := Circuit(fty).At(walk[0])
	if !ok {
		return nil, nil
	}
	switch t := v.(type) {
	case Faculty:
		if len(walk) == 1 {
			return fty, t
		}
		return t.Lookup(walk[1:]...)
	default: // non-faculty children are leaves (e.g. Circuit, Circuit, Gate)
		if len(walk) != 1 {
			panic("walk terminated")
		}
		return fty, t
	}
	panic(7)
}

func (fty Faculty) Refine(name string) Faculty {
	if x, ok := Circuit(fty).At(name); ok {
		return x.(Faculty)
	}
	y := NewFaculty(name)
	y.Genus().Walk = append(fty.Genus().Walk, name)
	Circuit(fty).Grow(name, y)
	return y
}

func (fty Faculty) AddTerminal(name string, term Meaning) {
	if _, ok := Circuit(fty).Include(name, term); ok {
		log.Fatalf("overwriting terminal %v->%T", name, term)
	}
}
