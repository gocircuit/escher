// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/hoijui/escher/be"
	// . "github.com/hoijui/escher/circuit"
	"github.com/hoijui/escher/faculty"
)

func init() {
	faculty.Register(be.NewSink(), "e", "Ignore")
	faculty.Register(be.NewMaterializer(&Grow{}), "e", "Grow")
	faculty.Register(be.NewMaterializer(&be.Union{}), "e", "Fork")
	faculty.Register(be.NewMaterializer(&Lens{}), "e", "Lens")
	faculty.Register(be.NewMaterializer(&Alternate{}), "e", "Alternate")
	faculty.Register(be.NewMaterializer(&Alternate{}), "e", "Alt")
	faculty.Register(be.NewMaterializer(&OneWayDoor{}), "e", "OneWayDoor")
	faculty.Register(be.NewMaterializer(Repeat{}), "e", "Repeat")
}
