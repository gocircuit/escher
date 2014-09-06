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

// Super is a symbol for the peer name of the point-of-view peer
type Super struct{}

func (s Super) String() string {
	return "*"
}

// Sugar is a valve name for auto-generated peers
type Sugar int

func (s Sugar) String() string {
	return fmt.Sprintf("sugar#%d", s)
}

// Default is the default valve's name
type Default struct{}

func (s Default) String() string {
	return "^"
}

func Understand(s *see.Circuit) *Circuit {
	x := &Circuit{peer: Make()}
	x.genus = []*see.Circuit{s}
	x.name = s.Name

	// Add the super peer
	sup := &Peer{
		name: Super{},
		index: 0,
		valve: Make(),
		design: nil,
	}
	x.peer[Super{}] = sup
	x.index = 1

	// Add peers from circuit definition, valves are not added on this pass
	for _, p := range s.Peer {
		x.addPeer(p.Name, x.index, p.Design)
		x.index++
	}
	var nsugar int // Counter for generating names of desugared peer definitions
	for _, m := range s.Match {
		var end [2]*Valve // reciprocals
		for i, join := range m.Join {
			switch t := join.(type) {
			case *see.DesignJoin: // unfold sugar
				nsugar++
				p := Sugar(nsugar)
				x.addPeer(p, x.index, t.Design)
				x.index++
				end[i] = x.reserveValve(p, Default{}, x.index)
				x.index++
			case *see.PeerJoin:
				end[i] = x.reserveValve(t.Peer, t.Valve, x.index)
				x.index++
			case *see.ValveJoin:
				end[i] = x.reserveValve(Super{}, t.Valve, x.index)
				x.index++
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
