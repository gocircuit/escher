// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package think

import (
	"strings"

	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/understand"
)

// Space encloses the entire faculty tree of an Escher program, starting from the root.
type Space understand.Faculty

func (x Space) Faculty() understand.Faculty {
	return understand.Faculty(x)
}

// Lookup…
// name should be a single-word identifier.
func (x Space) Lookup(within understand.Faculty, name string) (d see.Design) {
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
	println(cir.Print("	", "\t"))
	peers := make(map[string]Reflex)
	for _, peer := range cir.Peer {
		if peer.Name == "" { // skip the super peer of this circuit
			continue
		}
		switch t := peer.Design.(type) {
		case see.StringDesign, see.IntDesign , see.FloatDesign , see.ComplexDesign , see.TreeDesign:
			peers[peer.Name] = materializeBuiltin(t)
		case see.RootNameDesign:
			peers[peer.Name] = x.Materialize(strings.Split(string(t), ".")...)
		case see.NameDesign: // e.g. “hello.who.is.there”
			peers[peer.Name] = x.materializeName(within, string(t))
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
				println("m1", p.Name, v.Name)
			}
			delete(peers[p.Name], v.Name)
			if v.Matching.Of.Name == "" {
				super[v.Matching.Name] = m1
			} else {
				m2 := peers[v.Matching.Of.Name][v.Matching.Name]
				if m2 == nil {
					println("m2", v.Matching.Of.Name, v.Matching.Name)
				}
				delete(peers[v.Matching.Of.Name], v.Matching.Name)
				Merge(m1, m2)
			}
		}
	}
	// Check for unmatched valves
	for _, p := range cir.Peer {
		for _, v := range p.Valve {
			panicf("%s.%s not matched", p.Name, v.Name)
		}
	}
	return super
}

func (x Space) materializeName(within understand.Faculty, name string) Reflex {
	parts := strings.Split(name, ".")
	unfold := x.Lookup(within, parts[0])
	switch t := unfold.(type) {
	case see.StringDesign, see.IntDesign, see.FloatDesign, see.ComplexDesign, see.TreeDesign:
		if len(parts) != 1 {
			panic("constant designs do not have sub-faculties")
		}
		return materializeBuiltin(unfold)
	case see.NameDesign:
		parts = append(strings.Split(string(t), "."), parts[1:]...)
		return Space(within).Materialize(parts...)
	case see.RootNameDesign:
		parts = append(strings.Split(string(t), "."), parts[1:]...)
		return x.Materialize(parts...)
	case nil:
		return nil
	}
	panic("unknown design")
}

func materializeBuiltin(d see.Design) Reflex {
	switch t := d.(type) {
	case see.StringDesign:
		return NewNounReflex(string(t))
	case see.IntDesign:
		return NewNounReflex(int(t))
	case see.FloatDesign:
		return NewNounReflex(float64(t))
	case see.ComplexDesign:
		return NewNounReflex(complex128(t))
	case see.TreeDesign:
		return NewNounReflex(tree.Tree(t))
	}
	panic("unknown builting design")
}
