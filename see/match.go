// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"
	"github.com/gocircuit/escher/star"
)

// A matching is the following syntactic structure:
//
//	Matching —> Join “=” Join NewLine
//	Join —> ID “.” ID / ID / Design
//
// The star representation of a matching is:
//
//	{
//		Kind Name("Matching")
//		Left {
//			Peer Name(…)
//			Valve Name(…) or Design()
//		}
//		Right {
//			Peer Name(…)
//			Valve ??
//		}
//	}
//
func SeeMatching(src *Src, name string) (fwd, rev string, x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x = star.Make()
	x.Grow("Kind", "", Name("Matching"))
	t := src.Copy()
	Space(t)
	if left := SeeJoin(t); left != nil {
		x.Merge("Left", "", left)
	}
	Space(t)
	t.Match("=")
	Whitespace(t)
	if right := SeeJoin(t); right != nil {
		x.Merge("Right", "", right)
	}
	if !Space(t) { // require newline at end
		return "", "", nil
	}
	src.Become(t)
	return name+"_fwd", name+"_rev", x
}

// Join = 3.19 | Peer.Valve | Valve
//
//	{
//		Peer Name("??")
//		Valve Name("??")
//	}
//
func SeeJoin(src *Src) (x *star.Star) {
	if x = seeDesignJoin(src); x != nil { // int, string, etc.
		return x
	}
	if x = seePeerValveJoin(src); x != nil { // peer.valve
		return x
	}
	if x = seeValveJoin(src); x != nil { // valve (or empty string)
		return x
	}
	return nil
}

func seeDesignJoin(src *Src) (x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	d := SeeArithmeticOrStar(t)
	if d == nil {
		return nil
	}
	src.Become(t)
	return star.Make().Merge("Peer", "", d).Grow("Valve", "", Name(""))
}

// seePeerValveJoin…
func seePeerValveJoin(src *Src) (x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	peer := Identifier(t)
	if peer == "" {
		return nil
	}
	t.Match(".")
	valve := Identifier(t)
	if valve == "" {
		return nil
	}
	src.Become(t)
	return star.Make().Grow("Peer", "", Name(peer)).Grow("Valve", "", Name(valve))
}

// seeValveJoin parses a single identifier as a valve name
func seeValveJoin(src *Src) (x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	valve := Identifier(t)
	src.Become(t)
	return star.Make().Grow("Peer", "", Name("")).Grow("Valve", "", Name(valve))
}
