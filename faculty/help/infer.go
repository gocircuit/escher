// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package help

import (
	// "fmt"

	. "github.com/gocircuit/escher/circuit"
)

func Project(idiom Circuit) Circuit {
	r := New()
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
			yValve := vector.Valve

			// include address gates
			xgName, xgGate := xAddress.String(), xAddress
			ygName, ygGate := yAddress.String(), yAddress
			r.Include(xgName, xgGate)
			r.Include(ygName, ygGate)

			// include valve gates
			xvName := xgName + "#" + NameString(xValve)
			yvName := ygName + "#" + NameString(yValve)
			if !r.Has(xvName) {
				r.Include(xvName, "Valve")
			}
			if !r.Has(yvName) {
				r.Include(yvName, "Valve")
			}

			// link address to valve
			r.Link(Vector{xgName, xValve}, Vector{xvName, DefaultValve})
			r.Link(Vector{ygName, yValve}, Vector{yvName, DefaultValve})

			// link valve gates
			r.Link(Vector{xvName, fan(r, xvName)}, Vector{yvName, fan(r, yvName)})
		}
	}
	return r
}

func fan(r Circuit, vName Name) int {
	return len(r.Flow[vName]) - 1 // discount the default valve
}
