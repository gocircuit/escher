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
	"strings"
	"time"

	"github.com/gocircuit/circuit/client"
	"github.com/gocircuit/escher/faculty"
)

func Init(name string, client *client.Client) {
	rand.Seed(time.Now().UnixNano())
	ns := faculty.Root.Refine("circuit")
	ns.AddTerminal("process", Process{})
	// ns.AddTerminal("docker", Docker{})
	// ns.AddTerminal("chan", Chan{})
	// ns.AddTerminal("subscription", Subscription{})

	if name = strings.TrimSpace(name); name == "" {
		panic("Escher/Circuit program executions must have a non-empty instance name")
	}
	program = &Program{
		Name: name,
		Client: client,
	}
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
