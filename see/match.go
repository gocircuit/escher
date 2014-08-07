// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	// "fmt"
	. "github.com/gocircuit/escher/image"
)

// A matching is the following syntactic structure:
//
//	Matching => Join “=” Join NewLine
//	Join => ID “.” ID / ID / Design
//
// The star representation of a matching is:
//
//	{
//		Kind Name("Matching")
//		0 {
//			Peer Name(…)
//			Valve Name(…) or Design()
//		}
//		1 {
//			Peer Name(…)
//			Valve ??
//		}
//	}
//
func SeeMatching(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x = Make()
	t := src.Copy()
	Space(t)
	j0 := SeeJoin(t)
	x.Grow("0", j0)
	Whitespace(t)
	t.Match("=")
	Whitespace(t)
	j1 := SeeJoin(t)
	x.Grow("1", j1)
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return
}

// Join = 3.19 | Peer.Valve | Valve
//
//	{
//		Peer Name("??")
//		Valve Name("??")
//	}
//
func SeeJoin(src *Src) (x Image) {
	if x = seeDesignJoin(src); x != nil { // int, string, etc.
		return x
	}
	if x = seePeerValveJoin(src); x != nil { // peer.valve
		return x
	}
	if x = seeValveJoin(src); x != nil { // valve (or empty string)
		return x
	}
	panic(1)
}

func seeDesignJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	dimg := SeeArithmeticOrUnion(t)
	if dimg == nil {
		return nil
	}
	src.Become(t)
	return Image{
		"Peer": dimg,
		"Valve": Name(""),
	}
}

// seePeerValveJoin…
func seePeerValveJoin(src *Src) (x Image) {
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
	return Image{
		"Peer": Name(peer),
		"Valve": Name(valve),
	}
}

// seeValveJoin parses a single identifier as a valve name. 
// It captures the empty string.
func seeValveJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	t := src.Copy()
	valve := Identifier(t)
	src.Become(t)
	return Image{
		"Peer": Name(""),
		"Valve": Name(valve),
	}
}
