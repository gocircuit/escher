// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	// "fmt"
	// "strings"
)

func SeeMatching(src *Src) (m *Matching) {
	defer func() {
		if r := recover(); r != nil {
			m = nil
		}
	}()
	m = &Matching{}
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
	d, ok := SeeNoNameDesign(t)
	if !ok {
		return nil
	}
	if _, ok := d.(NameDesign); ok {
		return nil
	}
	if _, ok := d.(AbsNameDesign); ok {
		return nil
	}
	src.Become(t)
	return &DesignJoin{d}
}
