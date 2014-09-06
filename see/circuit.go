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

const DefaultValve = "_"

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
	Name  string
	Valve []string
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

// E.g. “a.X”
type PeerJoin struct {
	Peer  string
	Valve string
}

func (p *PeerJoin) String() string {
	return fmt.Sprintf("%s.%s", p.Peer, p.Valve)
}

// E.g. “Y”
type ValveJoin struct {
	Valve string
}

func (v *ValveJoin) String() string {
	return v.Valve
}

// E.g. “12.1e3”
type DesignJoin struct {
	Design interface{}
}

func (d *DesignJoin) String() string {
	return fmt.Sprintf("%v", d.Design)
}

func Circuitize(name string, img Image) (cir *Circuit) {
	if img == nil {
		return nil
	}
	cir = &Circuit{
		Peer:  make([]*Peer, 0, img.Len()),                        // # explicit peers + 1 (redundant) = # of src children = # peers + child for "$"
		Match: make([]*Matching, 0, img.Walk(MatchingName).Len()), // # of matchings
	}
	cir.Name = name
	for name, v := range img {
		if name == MatchingName {
			cir.seeMatching(v.(Image))
			continue
		}
		cir.Peer = append(
			cir.Peer,
			&Peer{
				Name:   name,
				Design: v,
			},
		)
	}
	return
}

func (cir *Circuit) seeMatching(s Image) {
	for _, x := range s {
		// fmt.Printf("=%s=>\n", string(w))
		m := &Matching{}
		for i := 0; i < 2; i++ {
			y := x.(Image).Walk(i)
			// fmt.Printf("    –%d–>\n    %s\n", i, y.Print("    ", "\t"))

			v := string(y["Valve"].(Name))
			switch p := y["Peer"].(type) {
			case Name:
				if string(p) == "" {
					m.Join[i] = &ValveJoin{
						Valve: v,
					}
					cir.Valve = append(cir.Valve, v)
				} else {
					m.Join[i] = &PeerJoin{
						Peer:  string(p),
						Valve: v,
					}
				}
			default:
				m.Join[i] = &DesignJoin{
					Design: p,
				}
			}
		}
		cir.Match = append(cir.Match, m)
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
