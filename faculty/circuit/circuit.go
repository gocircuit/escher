// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// circuit provides Escher gates for building dynamic cloud applications using the circuit runtime of http://gocircuit.org
package circuit

import (
	"fmt"
	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/tree"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	println("Loading http://gocircuit.org faculty")
	ns := faculty.Root.Refine("circuit")
	ns.AddTerminal("proc", Process{})
	// ns.AddTerminal("docker", Docker{})
	// ns.AddTerminal("chan", Chan{})
	// ns.AddTerminal("subscription", Subscription{})
}
