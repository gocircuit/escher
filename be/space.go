// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	"log"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/understand"
)

// Space encloses the entire faculty tree of an Escher program, starting from the root.
type Space understand.Faculty

func (x Space) Faculty() understand.Faculty {
	return understand.Faculty(x)
}

// Materialize creates the reflex, described by the absolute path walk.
func (x Space) Materialize(walk ...interface{}) Reflex {
	return x.materialize(
		&Matter{
			Name: []string{ChainKey("escher", []interface{}{"!spaek¡"})},
		}, 
		walk...)
}

func (x Space) materialize(matter *Matter, walk ...interface{}) Reflex {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Problem materializing %s (%v)", walk, r)
		}
	}()
	within, term := x.Faculty().Walk(walk...)
	switch t := term.(type) {
	case Gate:
		return t.Materialize()
	case GateWithMatter:
		return t.Materialize(matter)
	case *understand.Circuit:
		return x.materializeCircuit(
			&Matter{
				Name: append(matter.Name, ChainKey(matter.LastName(), walk)),
				Source: walk,
				Circuit: t,
				Faculty: within.(understand.Faculty),
				Super: matter,
			}, 
			within.(understand.Faculty), 
			t)
	case nil:
		log.Fatalf("Not found %v", walk)
	default:
		log.Fatalf("Not knowing how to materialize %v/%T", t, t)
	}
	panic(1)
}

func (x Space) materializeCircuit(matter *Matter, withinFac understand.Faculty, cir *understand.Circuit) Reflex {
	// println(cir.Print("	", "\t"))
	peers := make(map[interface{}]Reflex)
	for _, name := range cir.PeerNames() {
		peer := cir.PeerByName(name)
		if _, ok := name.(understand.Super); ok { // skip the super peer of this circuit
			continue
		}
		switch t := peer.Design().(type) {
		case see.Name: ??? cannot distinguish string from link!!!
			peers[name] = x.materialize(matter, t.AsWalk()...)
		case string, int, float64, complex128:
			peers[name] = NewNounReflex(t) // materialize builtin gates
		case Image:
			peers[name] = NewNounReflex(t.Copy()) // materialize images
		default:
			panicf("unknown design: %T/%v", t, t)
		}
	}
	// Connect/attach all reflex memories
	exo := make(Reflex)
	for _, name := range cir.PeerNames() {
		if _, ok := name.(understand.Super); ok { // skip the super peer of this circuit
			continue
		}
		p := cir.PeerByName(name)
		for _, vn := range p.ValveNames() {
			v := p.ValveByName(vn)
			// println(fmt.Sprintf("%s·%s", name, v.Name))
			m1 := peers[name][v.Name]
			if m1 == nil {
				continue
			}
			delete(peers[name], v.Name)
			if _, ok := v.Matching.Of.Name().(understand.Super); ok {
				if _, ok = exo[v.Matching.Name]; ok {
					panic(6)
				}
				exo[v.Matching.Name] = m1
			} else {
				qp, qv := v.Matching.Of.Name(), v.Matching.Name
				m2 := peers[qp][qv]
				delete(peers[qp], qv)
				if m1 == nil || m2 == nil {
					log.Fatalf("No matching valve: %v:%v/%v <=> %v:%v/%v", name, v.Name, m1, qp, qv, m2)
				}
				go Merge(m1, m2)
			}
		}
	}
	// Check for unmatched valves
	for pn, p := range peers {
		for vname, _ := range p {
			panicf("%s:%s not matched", pn, vname)
		}
	}
	return exo
}

func (x Space) Interpret(understood *understand.Circuit) (fresh *understand.Circuit) {
	return x.Faculty().Interpret(understood)
}

func (x Space) Forget(name interface{}) (forgotten interface{}) {
	switch t := x.Faculty().Forget(name).(type) {
	case understand.Faculty:
		return Space(t)
	default:
		return t
	}
	panic(0)
}

func (x Space) Walk(walk ...interface{}) (parent, child interface{}) {
	parent, child = x.Faculty().Walk(walk...)
	if fac, ok := parent.(understand.Faculty); ok {
		parent = Space(fac)
	}
	if fac, ok := child.(understand.Faculty); ok {
		child = Space(fac)
	}
	return
}

func (x Space) Roam(walk ...interface{}) (parent, child interface{}) {
	parent, child = x.Faculty().Roam(walk...)
	if fac, ok := parent.(understand.Faculty); ok {
		parent = Space(fac)
	}
	if fac, ok := child.(understand.Faculty); ok {
		child = Space(fac)
	}
	return
}
