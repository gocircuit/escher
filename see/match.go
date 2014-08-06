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
			x = NoImage
		}
	}()
	x = star.Make()
	t := src.Copy()
	Space(t)
	x.Merge("0", SeeJoin(t).Star)
	Whitespace(t)
	t.Match("=")
	Whitespace(t)
	x.Merge("1", SeeJoin(t).Star)
	if !Space(t) { // require newline at end
		return NoImage
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
	if x = seeDesignJoin(src); x.Lit() { // int, string, etc.
		return x
	}
	if x = seePeerValveJoin(src); x.Lit() { // peer.valve
		return x
	}
	if x = seeValveJoin(src); x.Lit() { // valve (or empty string)
		return x
	}
	return NoImage
}

func seeDesignJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = NoImage
		}
	}()
	t := src.Copy()
	dimg := SeeArithmeticOrImage(t)
	if dimg == nil {
		return NoImage
	}
	src.Become(t)
	return Imagine(star.Make().Merge("Peer", dimg).Grow("Valve", Name("")))
}

// seePeerValveJoin…
func seePeerValveJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = NoImage
		}
	}()
	t := src.Copy()
	peer := Identifier(t)
	if peer == "" {
		return NoImage
	}
	t.Match(".")
	valve := Identifier(t)
	if valve == "" {
		return NoImage
	}
	src.Become(t)
	return Imagine(star.Make().Grow("Peer", Name(peer)).Grow("Valve", Name(valve)))
}

// seeValveJoin parses a single identifier as a valve name. 
// It captures the empty string.
func seeValveJoin(src *Src) (x Image) {
	defer func() {
		if r := recover(); r != nil {
			x = NoImage
		}
	}()
	t := src.Copy()
	valve := Identifier(t)
	src.Become(t)
	return star.Make().Grow("Peer", Name("")).Grow("Valve", Name(valve))
}
