// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	// "fmt"
	"sync"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
)

// Associate
type Associate struct{}

func (Associate) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer((&association{}).Cognize, "Name", "With", "When", DefaultValve)
	return reflex
}

type association struct {
	sync.Mutex
	name string
	with interface{}
	when interface{}
}

func (x *association) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "Name":
		x.name = dvalue.(string)
	case "With":
		x.with = dvalue
	case "When":
		x.when = dvalue
	case DefaultValve:
	default:
		panic("eh")
	}
	eye.Show(
		DefaultValve, 
		Image{
			x.name: x.with,
			"When": x.when,
		},
	)
}
