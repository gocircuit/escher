// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	"fmt"
	"strconv"

	"github.com/gocircuit/escher/tree"
)

// Design is one of the built-in designs listed below.
type Design interface{
	String() string
}

type (
	NameDesign string
	AbsNameDesign string
	StringDesign string
	IntDesign int
	FloatDesign float64
	ComplexDesign complex128
	TreeDesign tree.Tree
)

func (x TreeDesign) String() string {
	return string((tree.Tree)(x).Marshal())
}

func (x ComplexDesign) String() string {
	return fmt.Sprintf("Complex(%g)", x)
}

func (x FloatDesign) String() string {
	return fmt.Sprintf("Float(%g)", x)
}

func (x IntDesign) String() string {
	return fmt.Sprintf("Int(%d)", x)
}

func (x NameDesign) String() string {
	return string(x)
}

func (x AbsNameDesign) String() string {
	return string(x)
}

func (x StringDesign) String() string {
	return fmt.Sprintf("String(%s)", strconv.Quote(chop(string(x))))
}

func chop(x string) string {
	if len(x) < 21 {
		return x
	}
	return x[:20]+"…"
}

// Reflex represents the name of a reflex and its arguments.
type Reflex struct {
	Name  string
	Valve []string
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
	return fmt.Sprintf("Peer{%s, %v}", p.Name, p.Design)
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
