// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	"log"
	"strings"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/understand"
)

// Space encloses the entire faculty tree of an Escher program, starting from the root.
type Space understand.Faculty

func (x Space) Faculty() understand.Faculty {
	return understand.Faculty(x)
}

// Lookup looks up the design for a root name; name should be a textual path.
func (x Space) Lookup(within understand.Faculty, name string) (d interface{}) {
	if strings.Index(name, ".") >= 0 {
		panic(7)
	}
	return within[""].(*understand.Circuit).Peer[name].Design
}

// Super describes the circuit that commissioned the present materialization.
type Super struct {
	Circuit *understand.Circuit
	Faculty understand.Faculty
	Super *Super
}

// Materialize creates the reflex, described by the absolute path walk.
func (x Space) Materialize(walk ...string) Reflex {
	return x.materialize(nil, walk...)
}

func (x Space) materialize(super *Super, walk ...string) Reflex {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Problem materializing %s (%v)", strings.Join(walk,"."), r)
		}
	}()
	within, term := x.Faculty().Walk(walk...)
	switch t := term.(type) {
	case Builtin:
		return t.Materialize(super)
	case Gate:
		return t.Materialize()
	case *understand.Circuit:
		return x.materializeCircuit(super, within.(understand.Faculty), t)
	}
	panic(1)
}

func (x Space) materializeCircuit(super *Super, withinFac understand.Faculty, cir *understand.Circuit) Reflex {
	super = &Super{
		Circuit: cir,
		Faculty: withinFac,
		Super: super,
	}
	// println(cir.Print("	", "\t"))
	peers := make(map[string]Reflex)
	for _, peer := range cir.Peer {
		name, ok := peer.Name.(string)
		if !ok {
			log.Fatalf("circuit peers cannot have non-textual names, in circuit %s at peer %v", cir.Name, peer.Name)
		}
		if name == "" { // skip the super peer of this circuit
			continue
		}
		switch t := peer.Design.(type) {
		case see.RootPath:
			peers[name] = x.materialize(super, []string(t)...)
		case see.Path: // e.g. “hello.who.is.there”
			peers[name] = x.materializePath(super, withinFac, []string(t))
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
	for _, p := range cir.Peer {
		name := p.Name.(string)
		if name == "" {
			continue
		}
		for _, v := range p.Valve {
			// println(fmt.Sprintf("%s·%s", name, v.Name))
			m1 := peers[name][v.Name]
			if m1 == nil {
				continue
			}
			delete(peers[name], v.Name)
			if v.Matching.Of.Name == "" {
				if _, ok := exo[v.Matching.Name]; ok {
					panic(6)
				}
				exo[v.Matching.Name] = m1
			} else {
				// ap, av := name, v.Name
				qp, qv := v.Matching.Of.Name.(string), v.Matching.Name
				m2 := peers[qp][qv]
				delete(peers[qp], qv)
				// go func() {
				// 	println(fmt.Sprintf("%s·%s/%v and %s·%s/%v", ap, av, m1, qp, qv, m2))
				// 	Merge(m1, m2)
				// }()
				go Merge(m1, m2)
			}
		}
	}
	// Check for unmatched valves
	for pname, p := range peers {
		for vname, _ := range p {
			panicf("%s.%s not matched", pname, vname)
		}
	}
	return exo
}

func (x Space) materializePath(super *Super, within understand.Faculty, parts []string) Reflex {
	unfold := x.Lookup(within, parts[0])
	switch t := unfold.(type) {
	case string, int, float64, complex128:
		return NewNounReflex(t) // materialize builtin gates
	case Image:
		return NewNounReflex(t.Copy())
	case see.Path:
		parts = append([]string(t), parts[1:]...)
		return Space(within).materialize(super, parts...)
	case see.RootPath:
		parts = append([]string(t), parts[1:]...)
		return x.materialize(super, parts...)
	case nil:
		return nil
	}
	panic("unknown design")
}

func (x Space) Interpret(understood *understand.Circuit) (fresh *understand.Circuit) {
	return x.Faculty().Interpret(understood)
}

func (x Space) Forget(name string) (forgotten interface{}) {
	switch t := x.Faculty().Forget(name).(type) {
	case understand.Faculty:
		return Space(t)
	default:
		return t
	}
	panic(0)
}

func (x Space) Walk(walk ...string) (parent, child interface{}) {
	parent, child = x.Faculty().Walk(walk...)
	if fac, ok := parent.(understand.Faculty); ok {
		parent = Space(fac)
	}
	if fac, ok := child.(understand.Faculty); ok {
		child = Space(fac)
	}
	return
}

func (x Space) Roam(walk ...string) (parent, child interface{}) {
	parent, child = x.Faculty().Roam(walk...)
	if fac, ok := parent.(understand.Faculty); ok {
		parent = Space(fac)
	}
	if fac, ok := child.(understand.Faculty); ok {
		child = Space(fac)
	}
	return
}
