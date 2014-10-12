// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package think

import (
	// "fmt"
	"sync"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Remember
type Remember struct{}

func (Remember) Materialize() be.Reflex {
	reflex, _ := be.NewEyeCognizer((&remember{}).Cognize, "From", "What", "When", DefaultValve)
	return reflex
}

type remember struct {
	sync.Mutex
	from Circuit
	what Circuit
	when interface{}
}

func (x *remember) Cognize(eye *be.Eye, dvalve string, dvalue interface{}) {
	x.Lock()
	defer x.Unlock()
	switch dvalve {
	case "From":
		x.from = dvalue.(Circuit)
	case "What":
		x.what = dvalue.(Circuit)
	case "When":
		x.when = dvalue
	case DefaultValve:
	default:
		panic("eh")
	}
	eye.Show(
		DefaultValve, 
		Circuit{
			"Memory": emphasize(x.from, x.what),
			"When": x.when,
		},
	)
}

func emphasize(from, what Circuit) Circuit {
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
		case Circuit:
			r[name] = emphasize(subwhat.(Circuit), t)
		default: // Go primitive types and nil
			r[name] = t
		}
	}
	return r
}