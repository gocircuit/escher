// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	// "fmt"
// 	"log"

// 	"github.com/gocircuit/escher/faculty/basic"
	// . "github.com/gocircuit/escher/circuit"
// 	"github.com/gocircuit/escher/be"
// 	"github.com/gocircuit/escher/kit/plumb"
)

// p is the predicate circuit.
// 
// h is the template for the output circuit.
// The output y will have the same structure as h, while the meanings of the
// images in h are substituted according to the following rule.
//
// (a) If the meaning in h is irreducible (int, float, complex, string), it is kept the same.
// (b) If it is an address, it substituted with a corresponding value from the predicate match
// or from the auxiliary input. Addresses starting “p.name” refer to the meaning in x matching p.
// (c) If it is a circuit, ... 
//
// a is the auxiliary input circuit.
// x is the input circuit.
// func transform(p, h, a, x Circuit) (y Circuit) {
// 	y = h.Copy()
// 	for n, v := range y.Images() {
// 		switch t := 
// 	}
// }
