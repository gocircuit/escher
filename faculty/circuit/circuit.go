// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// circuit provides Escher gates for building dynamic cloud applications using the circuit runtime of http://gocircuit.org
package circuit

import (
	// "fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gocircuit/circuit/client"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	ns := faculty.Root.Refine("circuit")
	ns.AddTerminal("process", Process{})
	// ns.AddTerminal("docker", Docker{})
	// ns.AddTerminal("chan", Chan{})
	// ns.AddTerminal("subscription", Subscription{})

	// program = &Program{
	// 	Name: "???",
	// 	Client: ???,
	// }
}

// Program…
type Program struct {
	Name string // Textual ID of this Escher program process to be used as part of circuit anchor names
	*client.Client
}

var program *Program

// ChooseID returns a unique textual ID.
func ChooseID() string {
	return strconv.FormatUint(uint64(rand.Int63()), 20)
}
