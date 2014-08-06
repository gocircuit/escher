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
//		$ {
//			0 {
//				Peer	Name("a")
//				Valve Name("XandY")
//			}
//			1 {
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
)

// Design is a type that is meant to be stored in the value of a star.
type Design interface{
	String() string
}

// Name
type Name string

func (x Name) String() string {
	return fmt.Sprintf("Name(%s)", string(x))
}

// RootName
type RootName string

func (x RootName) String() string {
	return fmt.Sprintf("Story(%s)", string(x))
}
