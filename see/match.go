// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/gocircuit/escher/star"
)

// A MATCHING is one of the following syntactic structures
//
//	MATCHING: JOIN “=” JOIN
//	JOIN: ID “.” ID | ID | DESIGN
//
// The star encoding of a MATCHING is:
//
//	{
//		"Matching" // empty-string choice
//		Left {
//			Peer "" // string indicates a peer name; star is a circuit or a built-in design
//			Valve "X"
//		}
//		Right {
//			Peer "f"
//			Valve "A"
//		}
//	}
//
func SeeMatching(src *Src) (m *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			m = nil
		}
	}()
	m = star.Make()
	t := src.Copy()
	Space(t)
	if m.Join[0] = SeeJoin(t); m.Join[0] == nil {
		return nil
	}
	if Space(t) {
		src.Become(t)
		return
	}
	t.Match("=")
	Space(t)
	if m.Join[1] = SeeJoin(t); m.Join[1] == nil {
		return nil
	}
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return
}

// “one.Two” or “Wolf” or “+3.12e2”
func SeeJoin(src *Src) Join {
	if j := parseDesignJoin(src); j != nil {
		//fmt.Printf("Design=%T/%v [%s]\n", j, j, src.String())
		return j
	}
	if j := parsePeerJoin(src); j != nil {
		return j
	}
	if j := parseValveJoin(src); j != nil {
		return j
	}
	return nil
}

// “one.Two”
func parsePeerJoin(src *Src) (peer *PeerJoin) {
	defer func() {
		if r := recover(); r != nil {
			peer = nil
		}
	}()
	t := src.Copy()
	peer = &PeerJoin{}
	if peer.Peer = Identifier(t); peer.Peer == "" {
		return nil
	}
	t.Match(".")
	if peer.Valve = Identifier(t); peer.Valve == "" {
		return nil
	}
	src.Become(t)
	return
}

// “Wolf”
func parseValveJoin(src *Src) (valve *ValveJoin) {
	defer func() {
		if r := recover(); r != nil {
			valve = nil
		}
	}()
	t := src.Copy()
	valve = &ValveJoin{}
	if valve.Valve = Identifier(t); valve.Valve == "" {
		return nil
	}
	src.Become(t)
	return
}

// “+3.12e5”
func parseDesignJoin(src *Src) (design *DesignJoin) {
	defer func() {
		if r := recover(); r != nil {
			design = nil
		}
	}()
	t := src.Copy()
	d, ok := SeeNoName(t)
	if !ok {
		return nil
	}
	if _, ok := d.(Name); ok {
		return nil
	}
	if _, ok := d.(RootName); ok {
		return nil
	}
	src.Become(t)
	return &DesignJoin{d}
}
