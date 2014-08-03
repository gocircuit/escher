// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"fmt"
	"strconv"
	"github.com/gocircuit/escher/star"
)

// func SeeCircuit(src *Src) (string, *star.Star) {
// 	return SeePeer(src)
// }

func SeeCircuit(src *Src) *Circuit {
	return Circuitize(SeePeer(src))
}

func Circuitize(name string, x *star.Star) (cir *Circuit) {
	img := x.Interface().(*Image).Unwrap()
	cir = &Circuit{
		Peer: make([]*Peer, 0, img.Len()), // # explicit peers + default empty-string peer = # of src children = # peers + child for "$"
		Match: make([]*Matching, 0, img.Down(MatchingName).Len()), // # of matchings
	}
	cir.Name = name
	cir.Peer = append(
		cir.Peer, 
		&Peer{
			Name: "",
			Design: nil, // no design for implied peer
		}, // default empty-string peer
	)
	for name, v := range img.Choice {
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
	for w, x := range s.Choice {
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

type Peer struct {
	Name   string
	Design Design
}

func (p *Peer) String() string {
	return fmt.Sprintf("Peer(%s, %v)", p.Name, p.Design)
}

type Matching struct {
	Join [2]Join
}

// Join is one of PeerJoin, ValveJoin or DesignJoin.
type Join interface{}

// E.g. “a.X”
type PeerJoin struct {
	Peer string
	Valve string
}

// E.g. “Y”
type ValveJoin struct {
	Valve string
}

// E.g. “12.1e3”
type DesignJoin struct {
	Design Design
}
