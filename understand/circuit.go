// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package understand

import (
	"fmt"

	"github.com/gocircuit/escher/see"
)

// TODO: Eventually Circuit can be represented as a tree.Tree
type Circuit struct {
	Genus []*see.Circuit // origin
	Name string
	Peer map[string]*Peer // peers and self; self corresponds to the empty string
}

type Peer struct {
	Name string
	Design see.Design
	Valve map[string]*Valve
}

type Valve struct {
	Of *Peer
	Name string
	Matching *Valve
}

func (x *Circuit) Merge(y *Circuit) {
	x.Genus = append(x.Genus, y.Genus[0])
	for name, p := range y.Peer {
		println(x.Name, "		merging peer", name)
		if _, ok := x.Peer[name]; ok {
			panicf("collision adding peer %s in circuit %s", name, x.Name)
		}
		x.Peer[name] = p
	}
}

func Understand(s *see.Circuit) *Circuit {
	x := &Circuit{
		Peer: make(map[string]*Peer),
	}
	x.Genus = []*see.Circuit{s}
	x.Name = s.Name

	// Add “this” circuit as the empty-string peer
	x.Peer[""] = &Peer{
		Name: "",
		Valve: make(map[string]*Valve),
		Design: nil,
	}

	// Add peers from circuit definition, valves are not added on this pass
	for _, p := range s.Peer {
		x.addPeer(p.Name, p.Design)
	}
	var nsugar int // Counter for generating names of desugared peer definitions
	for _, m := range s.Match {
		var end [2]*Valve // reciprocal
		for i, j := range m.Join {
			switch t := j.(type) {
			case *see.DesignJoin: // unfold sugar
				nsugar++
				p := fmt.Sprintf("_sugar_%d", nsugar)
				x.addPeer(p, t.Design)
				end[i] = x.reserveValve(p, "") // Anonymous designs have one empty-string valve
			case *see.PeerJoin:
				end[i] = x.reserveValve(t.Peer, t.Valve)
			case *see.ValveJoin:
				end[i] = x.reserveValve("", t.Valve)
			case nil: // match other argument to empty-string valve of this circuit
				end[i] = x.reserveValve("", "")
			default:
				panic(fmt.Sprintf("unknown or missing matching endpoint: %T·%v", j, j))
			}
		}
		// Link two ends
		if end[0].Matching != nil || end[1].Matching != nil {
			panic("reuse of valve")
		}
		end[0].Matching, end[1].Matching = end[1], end[0]
	}

	// Verify no unmatched valves remain
	// for _, peer := range x.Peer {
	// 	for _, v := range peer.Valve {
	// 		if v.Matching == nil {
	// 			panic("unmatched valve")
	// 		}
	// 	}
	// }
	return x
}

func (x *Circuit) addPeer(name string, design see.Design) {
	if _, ok := x.Peer[name]; ok {
		panic("peer already present")
	}
	if design == nil {
		panic("peer is mising design")
	}
	x.Peer[name] = &Peer{
		Name: name,
		Valve: make(map[string]*Valve),
		Design: design,
	}
}

// reserveValve returns the addressed valve, making it if necessary.
// Making is prohibited solely for the empty-string peer, corresponding to this circuit.
func (x *Circuit) reserveValve(peer, valve string) *Valve {
	p, ok := x.Peer[peer]
	if !ok {
		panic("peer is missing")
	}
	v, ok := p.Valve[valve]
	if ok {
		return v
	}
	v = &Valve{
		Of: p,
		Name: valve,
		Matching: nil,
	}
	p.Valve[valve] = v
	return v
}
