// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package spin

import (
	// "log"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

/*
	View {
		Time int
		Position complex128
		Orientation complex128 // dilation and rotation
	}
*/

func init() {
	faculty.Register("spin.Age", be.NewNativeMaterializer(Age{}))
	faculty.Register("spin.Die", be.NewNativeMaterializer(&Die{}))
	faculty.Register("spin.Split", be.NewNativeMaterializer(&Split{}))
	//
	faculty.Register("spin.Dilate", be.NewNativeMaterializer(&Dilate{}))
	faculty.Register("spin.Move", be.NewNativeMaterializer(Move{}))
}

// Age…
type Age struct{}

func (Age) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Age) CognizeView(eye *be.Eye, val interface{}) {
	v := val.(Circuit)
	v.Gate["Time"] = v.IntAt("Time") + 1
	eye.Show(DefaultValve, v)
}

func (Age) Cognize(*be.Eye, interface{}) {}
