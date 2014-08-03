// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"
	"strconv"
	"github.com/gocircuit/escher/star"
)

func SeeCircuit(src *Src) *Circuit {
	if src.Len() == 0 {
		return nil
	}
	return Circuitize(SeePeer(src))
}

func Circuitize(name string, x *star.Star) (cir *Circuit) {
	if x == nil {
		return nil
	}
	img := x.Interface().(*Image).Unwrap()
	cir = &Circuit{
		Peer: make([]*Peer, 0, img.Len()), // # explicit peers + 1 (redundant) = # of src children = # peers + child for "$"
		Match: make([]*Matching, 0, img.Down(MatchingName).Len()), // # of matchings
	}
	cir.Name = name
	for name, v := range img.Arm {
		if name == star.Parent {
			continue
		}
		if name == MatchingName {
			cir.seeMatching(v)
			continue
		}
		cir.Peer = append(
			cir.Peer,
			&Peer{
				Name: name,
				Design: v.Interface().(Design),
			},
		)
	}
	return
}

func (cir *Circuit) seeMatching(s *star.Star) {
	for w, x := range s.Arm {
		if string(w) == star.Parent {
			continue
		}
		// fmt.Printf("=%s=>\n", string(w))
		m := &Matching{}
		for i := 0; i < 2; i++ {
			y := x.Down(strconv.Itoa(i))
			// fmt.Printf("    –%d–>\n    %s\n", i, y.Print("    ", "\t"))

			v := string(y.Down("Valve").Interface().(Name))
			switch p := y.Down("Peer").Interface().(type) {
			case Name:
				if string(p) == "" {
					m.Join[i] = &ValveJoin{
						Valve: v,
					}
				} else {
					m.Join[i] = &PeerJoin{
						Peer: string(p),
						Valve: v,
					}
				}
			case Design:
				m.Join[i] = &DesignJoin{
					Design: p,
				}
			default:
				panic(1) // parsing bug
			}
		}
		cir.Match = append(cir.Match, m)
	}
}

type Circuit struct {
	Name string
	Valve []string
	Peer    []*Peer
	Match []*Matching
}

func (c *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s (", c.Name)
	for i, v := range c.Valve {
		var comma = ", "
		if i + 1 == len(c.Valve) {
			comma = ""
		}
		fmt.Fprintf(&w, "%s%s", v, comma)
	}
	w.WriteString(") {\n")
	for _, p := range c.Peer {
		fmt.Fprintf(&w, "%s%s%v\n", prefix, indent, p)
	}
	for _, m := range c.Match {
		fmt.Fprintf(&w, "%s%s%v\n", prefix, indent, m)
	}
	fmt.Fprintf(&w, "%s}\n", prefix)
	return w.String()
}

type Peer struct {
	Name   string
	Design Design
}

func (p *Peer) String() string {
	return fmt.Sprintf("%s %v", p.Name, p.Design)
}

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
	Peer string
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
	Design Design
}

func (d *DesignJoin) String() string {
	return d.Design.String()
}
