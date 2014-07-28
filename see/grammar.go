// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// The source:
//
//	nand {
//		a and
//		n not
//		a.XandY = n.X
//		= n.notX // the empty-string valve of the nand circuit connects to n's notX valve
//		a.X = X
//		a.Y = 1
//	}
//
// Has the following desugared representation after parsing:
//
//	nand {
//		a Name("and")
//		n Name("not")
//		matching_1 {
//			Left {
//				Peer	Name("a")
//				Valve Name("XandY")
//			}
//			Right {
//				Peer	Name("n")
//				Valve Name("X")
//			}
//		}
//		…
//	}
//
package see

import (
	"fmt"
	"strconv"
)

// Design is a type that is meant to be stored in the value of a star.
type Design interface{
	String() string
}

type (
	Name string
	RootName string
	String string
	Int int
	Float float64
	Complex complex128
)

func (x Complex) String() string {
	return fmt.Sprintf("Complex%g", x)
}

func (x Float) String() string {
	return fmt.Sprintf("Float(%g)", x)
}

func (x Int) String() string {
	return fmt.Sprintf("Int(%d)", x)
}

func (x Name) String() string {
	return fmt.Sprintf("Name(%s)", strconv.Quote(string(x)))
}

func (x RootName) String() string {
	return fmt.Sprintf("RootName(%s)", strconv.Quote(string(x)))
}

func (x String) String() string {
	return fmt.Sprintf("String(%s)", strconv.Quote(chop(string(x))))
}

func chop(x string) string {
	if len(x) < 21 {
		return x
	}
	return x[:20]+"…"
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
