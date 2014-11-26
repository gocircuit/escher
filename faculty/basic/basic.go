// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	// . "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
)

func init() {
	faculty.Register(be.NewSink, "Ignore")
	//
	faculty.Register(be.NewMaterializer(&Grow{}), "Grow")
	faculty.Register(be.NewMaterializer(&be.Union{}), "Fork")
	faculty.Register(be.NewMaterializer(&Lens{}), "Lens")
	//
	faculty.Register(be.NewMaterializer(&Alternate{}), "Alternate")
	faculty.Register(be.NewMaterializer(&Alternate{}), "Alt")
	faculty.Register(be.NewMaterializer(&OneWayDoor{}), "OneWayDoor")
	//
	faculty.Register(be.NewMaterializer(Repeat{}), "Repeat")
}
