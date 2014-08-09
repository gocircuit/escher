// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	// "fmt"
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

// Lookup looks up the design of a faculty-local name; name should be a single-word identifier.
func (x Space) Lookup(within understand.Faculty, name string) (d interface{}) {
	if strings.Index(name, ".") >= 0 {
		panic(7)
	}
	return within[""].(*understand.Circuit).Peer[name].Design
}

// Materialize creates the reflex, described by the absolute path walk.
func (x Space) Materialize(walk ...string) Reflex {
	within_, term := x.Faculty().Walk(walk...)
	if gate, ok := term.(Gate); ok {
		return gate.Materialize()
	}
	within, cir := within_.(understand.Faculty), term.(*understand.Circuit)
	//println(cir.Print("	", "\t"))
	peers := make(map[string]Reflex)
	for _, peer := range cir.Peer {
		if peer.Name == "" { // skip the super peer of this circuit
			continue
		}
		switch t := peer.Design.(type) {
		case see.RootName:
			peers[peer.Name] = x.Materialize(strings.Split(string(t), ".")...)
		case see.Name: // e.g. “hello.who.is.there”
			peers[peer.Name] = x.materializeName(within, string(t))
		case string, int , float64, complex128, Image:
			peers[peer.Name] = NewNounReflex(t) // materialize builtin gates
		default:
			panicf("unknown design: %T/%v", t, t)
		}
	}
	// Connect/attach all reflex memories
	super := make(Reflex)
	for _, p := range cir.Peer {
		if p.Name == "" {
			continue
		}
		for _, v := range p.Valve {
			m1 := peers[p.Name][v.Name]
			if m1 == nil {
				continue
			}
			delete(peers[p.Name], v.Name)
			if v.Matching.Of.Name == "" {
				if _, ok := super[v.Matching.Name]; ok {
					panic(6)
				}
				super[v.Matching.Name] = m1
			} else {
				qp, qv := v.Matching.Of.Name, v.Matching.Name
				m2 := peers[qp][qv]
				delete(peers[qp], qv)
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
	return super
}

func (x Space) materializeName(within understand.Faculty, name string) Reflex {
	parts := strings.Split(name, ".")
	unfold := x.Lookup(within, parts[0])
	switch t := unfold.(type) {
	case string, int, float64, complex128, Image:
		return NewNounReflex(t) // materialize builtin gates
	case see.Name:
		parts = append(strings.Split(string(t), "."), parts[1:]...)
		return Space(within).Materialize(parts...)
	case see.RootName:
		parts = append(strings.Split(string(t), "."), parts[1:]...)
		return x.Materialize(parts...)
	case nil:
		return nil
	}
	panic("unknown design")
}
