// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"

	. "github.com/gocircuit/escher/image"
)

func SeeCircuit(src *Src) *Circuit {
	if src.Len() == 0 {
		return nil
	}
	for p, un := range SeePeer(src) {
		// fmt.Printf("p=%s v=%v\n", p, un)
		return Circuitize(p, un.(Image))
	}
	return nil
}

// Circuit
type Circuit struct {
	Name  interface{}
	Valve []interface{}
	Peer  []*Peer
	Match []*Matching
}

// Peer
type Peer struct {
	Name   interface{}
	Design interface{}
}

func (p *Peer) String() string {
	return fmt.Sprintf("%v %v", p.Name, p.Design)
}

// Matching
type Matching struct {
	Join [2]Join
}

func (m *Matching) String() string {
	return fmt.Sprintf("%v=%v", m.Join[0], m.Join[1])
}

// Join is one of PeerJoin, ValveJoin or DesignJoin.
type Join interface{}

// PeerJoin
type PeerJoin struct {
	Peer  interface{}
	Valve interface{}
}

func (p *PeerJoin) String() string {
	return fmt.Sprintf("%v:%v", p.Peer, p.Valve)
}

// DesignJoin
type DesignJoin struct {
	Design interface{}
}

func (d *DesignJoin) String() string {
	return fmt.Sprintf("%v", d.Design)
}

// ValveJoin
type ValveJoin struct {
	Valve interface{}
}

func (v *ValveJoin) String() string {
	return fmt.Sprintf("%v", v.Valve)
}

func Circuitize(name interface{}, img Image) (cir *Circuit) {
	if img == nil {
		return nil
	}
	cir = &Circuit{
		Peer:  make([]*Peer, 0, img.Len()), // # explicit peers + 1 (redundant) = # of src children = # peers + child for "$"
		Match: make([]*Matching, 0, img.Walk(Matchings{}).Len()), // # of matchings
	}
	cir.Name = name
	for nm, v := range img {
		if _, ok := nm.(Matchings); ok {
			cir.seeMatching(cir.Name, v.(Image))
			continue
		}
		cir.Peer = append(
			cir.Peer,
			&Peer{
				Name: nm,
				Design: v,
			},
		)
	}
	return
}

func (cir *Circuit) seeMatching(name interface{}, s Image) {
	for _, x := range s {
		// fmt.Printf("=%s=>\n", string(w))
		m := &Matching{}
		for i := 0; i < 2; i++ {
			y := x.(Image).Walk(i)
			// fmt.Printf("    –%d–>\n    %s\n", i, y.Print("    ", "\t"))
			if y.Has("Design") {
				m.Join[i] = &DesignJoin{Design: y.Interface("Design")}
			} else {
				p, v := y["Peer"], y["Valve"]
				if p == name {
					m.Join[i] = &ValveJoin{Valve: v}
					cir.Valve = append(cir.Valve, v)
				} else {
					m.Join[i] = &PeerJoin{Peer:  p, Valve: v}
				}
			}
		}
		cir.Match = append(cir.Match, m)
	}
}

func (c *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%v ", c.Name)
	if len(c.Valve) > 0 {
		w.WriteString("(")
	}
	for i, v := range c.Valve {
		var comma = ", "
		if i+1 == len(c.Valve) {
			comma = ""
		}
		fmt.Fprintf(&w, "%v%s", v, comma)
	}
	if len(c.Valve) > 0 {
		w.WriteString(") {\n")
	} else {
		w.WriteString("{\n")
	}
	for _, p := range c.Peer {
		fmt.Fprintf(&w, "%s%s%v\n", prefix, indent, p)
	}
	for _, m := range c.Match {
		fmt.Fprintf(&w, "%s%s%v\n", prefix, indent, m)
	}
	fmt.Fprintf(&w, "%s}\n", prefix)
	return w.String()
}
