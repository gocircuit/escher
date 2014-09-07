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
		return Circuitize(p.(string), un.(Image))
	}
	return nil
}

// Circuit
type Circuit struct {
	Name string
	Valve []string
	Peer  []*Peer
	Match []*Matching
}

// Peer
type Peer struct {
	Name   string
	Design interface{}
}

func (p *Peer) String() string {
	return fmt.Sprintf("%s %v", p.Name, p.Design)
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
	Peer  string
	Valve string
}

func (p *PeerJoin) String() string {
	return fmt.Sprintf("%s:%s", p.Peer, p.Valve)
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
	Valve string
}

func (v *ValveJoin) String() string {
	return fmt.Sprintf("%s", v.Valve)
}

func Circuitize(name string, img Image) (cir *Circuit) {
	// println(img.Print("", "  "))
	if img == nil {
		return nil
	}
	if len(img.Letters()) != len(img.Numbers()) {
		panic(1)
	}
	cir = &Circuit{
		Peer:  make([]*Peer, len(img.Letters())), // # of peers
		Match: make([]*Matching, img.Walk(Matchings{}).Len()), // # of matchings
	}
	cir.Name = name
	for _, index := range img.Numbers() { // list peers in order of definition, by sorted int name
		pn := img.String(index) // lookup string name (of peer)
		pv := img.Interface(pn) // lookup value of peer
		if pn == cir.Name {
			panic("peer duplicates name with super peer")
		}
		cir.Peer[index] = &Peer{
			Name: pn,
			Design: pv,
		}
	}
	cir.seeMatching(cir.Name, img.OptionalImage(Matchings{}))
	return
}

func (cir *Circuit) seeMatching(name string, s Image) {
	if s == nil { // no matchings section
		return
	}
	for _, xn := range s.Numbers() {
		x := s.Image(xn)
		m := &Matching{}
		for i := 0; i < 2; i++ {
			y := x.Image(i)
			if y.Has("Design") {
				m.Join[i] = &DesignJoin{Design: y.Interface("Design")}
			} else {
				p, v := y.String("Peer"), y.String("Valve")
				if p == name {
					m.Join[i] = &ValveJoin{Valve: v}
					cir.Valve = append(cir.Valve, v)
				} else {
					m.Join[i] = &PeerJoin{Peer:  p, Valve: v}
				}
			}
		}
		cir.Match[xn] = m
	}
}

func (c *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s ", c.Name)
	if len(c.Valve) > 0 {
		w.WriteString("(")
	}
	for i, v := range c.Valve {
		var comma = ", "
		if i+1 == len(c.Valve) {
			comma = ""
		}
		fmt.Fprintf(&w, "%s%s", v, comma)
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
