// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	"fmt"

	"github.com/gocircuit/escher/see"
	. "github.com/gocircuit/escher/image"
)

// SugarValve is the default valve name, for sugar gates
const SugarValve = "_"

func Understand(s *see.Circuit) *Circuit {
	x := &Circuit{peer: Make()}
	x.genus = []*see.Circuit{s}
	x.name = s.Name

	// Add the super peer
	sup := &Peer{
		name: s.Name,
		index: 0,
		valve: Make(),
		design: nil, // indicates super
	}
	x.peer[s.Name] = sup

	// Add non-super peers from circuit definition. Valves are not added on this pass
	for i, p := range s.Peer {
		x.addPeer(p.Name, i+1, p.Design)
	}
	var nsugar int // Counter for generating names of desugared peer definitions
	for l, m := range s.Match {
		var end [2]*Valve // reciprocals
		for i, join := range m.Join {
			switch t := join.(type) {
			case *see.DesignJoin: // unfold sugar
				nsugar++
				pn := fmt.Sprintf("sugar#%d", nsugar)
				x.addPeer(pn, ??, t.Design)
				end[i] = x.reserveValve(pn, SugarValve, l)
			case *see.PeerJoin:
				end[i] = x.reserveValve(t.Peer, t.Valve, l)
			case *see.ValveJoin:
				end[i] = x.reserveValve(Super{}, t.Valve, l)
			default:
				panic(fmt.Sprintf("unknown or missing matching endpoint: %T·%v", t, t))
			}
		}
		// Link two ends
		if end[0].Matching != nil || end[1].Matching != nil {
			panic("reuse of valve")
		}
		end[0].Matching, end[1].Matching = end[1], end[0]
	}

	// Verify no dangling/unmatched valves remain
	for _, pn := range x.PeerNames() {
		p := x.PeerByName(pn)
		for _, vn := range p.ValveNames() {
			if p.ValveByName(vn).Matching == nil {
				panic("unmatched valve")
			}
		}
	}

	return x
}

func (x *Circuit) Merge(y *Circuit) {
	if len(y.genus) != 1 {
		panic(1)
	}
	x.genus = append(x.genus, y.genus[0])
	var m int // track max index
	for _, pn := range y.PeerNames() {
		if x.PeerByName(pn) != nil {
			panicf("collision adding peer %s in circuit %s", pn, x.Name())
		}
		q := y.PeerByName(pn).Copy()
		q.index += x.index
		m = max(m, q.index)
		x.attachPeer(q)
	}
	x.index += y.index
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (x *Circuit) attachPeer(p *Peer) {
	if _, ok := x.peer[p.name]; ok {
		panic("peer name already present")
	}
	if p.design == nil {
		panic("peer is missing design")
	}
	x.peer[p.name] = p
}

func (x *Circuit) addPeer(name interface{}, index int, design interface{}) {
	if _, ok := x.peer[name]; ok {
		panic("peer name already present")
	}
	if design == nil {
		panic("peer is missing design")
	}
	p := &Peer{
		name: name,
		index: index,
		valve: Make(),
		design: design,
	}
	x.peer[name] = p
}

// reserveValve returns the addressed valve, creating it if necessary.
// Creating is prohibited solely for the empty-string peer, corresponding to this circuit.
func (x *Circuit) reserveValve(peer, valve interface{}, index int) *Valve {
	p := x.PeerByName(peer)
	if p == nil {
		panic(fmt.Sprintf("peer %v is missing", peer))
	}
	v := p.ValveByName(valve)
	if v != nil {
		return v
	}
	v = &Valve{
		Of: p,
		Name: valve,
		Index: index,
		Matching: nil,
	}
	p.valve[valve] = v
	return v
}
