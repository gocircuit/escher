// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package help

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
	// "github.com/gocircuit/escher/be"
	// "github.com/gocircuit/escher/faculty"
)

func init() {
	// faculty.Register("help.Analyze", Analyze{})
}

func Infer(idiom Circuit) Circuit {
	result := New()
	for xName, xView := range idiom.Flow {
		xAddress, ok := idiom.Gate[xName].(Address)
		if !ok {
			continue
		}
		for xValve, vector := range xView {
			yAddress, ok :=  idiom.Gate[vector.Gate].(Address) // vector.Gate = yName
			if !ok {
				continue
			}
			xxName, xxGate := xAddress.String(), xAddress
			yyName, yyGate := yAddress.String(), yAddress
			result.Include(xxName, xxGate)
			result.Include(yyName, yyGate)
			result.Link(Vector{xxName, xValve}, Vector{yyName, vector.Valve}) // vector.Valve = yValve
		}
	}
	return result
}
