// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package help

import (
	// "fmt"
	"testing"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/see"
)

func TestProject(t *testing.T) {
	idiom := see.ParseCircuit(`
		{
			x1 abc.X
			x2 abc.X
			y def.Y
			z def.Z
			u rst.U
			x1:YZ = y:X
			x2:YZ = z:X
			y:Y = y:YY
		}
	`)
	println(String(Project(idiom)))
}
