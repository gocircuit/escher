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
	"github.com/gocircuit/escher/kit/plumb"
)

// Remember
type Remember struct{}

func (Remember) Materialize(*be.Matter) be.Reflex {
	reflex, _ := plumb.NewEyeCognizer((&remember{}).Cognize, "From", "What", "When", "_")
	return reflex
}

type remember struct {
	sync.Mutex
	from Image
	what Image
	when interface{}
}

func (x *remember) Cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "From":
		x.from = dvalue.(Image)
	case "What":
		x.what = dvalue.(Image)
	case "When":
		x.when = dvalue
	case "_":
	default:
		panic("eh")
	}
	eye.Show(
		"_", 
		Image{
			"Memory": emphasize(x.from, x.what),
			"When": x.when,
		},
	)
}

func emphasize(from, what Image) Image {
	if what == nil {
		return from
	}
	r := Make()
	for name, subwhat := range what {
		v, present := from[name]
		if !present {
			continue
		}
		switch t := v.(type) {
		case Image:
			r[name] = emphasize(subwhat.(Image), t)
		default: // Go primitive types and nil
			r[name] = t
		}
	}
	return r
}