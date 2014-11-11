// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Generalize struct {}

func (Generalize) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Generalize) CognizeIndex(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, GeneralizeIndex(v.(Circuit)))
}

func (Generalize) Cognize(eye *be.Eye, v interface{}) {}

func GeneralizeIndex(u Circuit) Circuit {
	r := generalizeIndex(be.AsIndex(u))
	r.Memorize("int", "$", "int")
	r.Memorize("float64", "$", "float")
	r.Memorize("complex128", "$", "complex")
	r.Memorize("string", "$", "string")
	r.Memorize("bool", "$", "bool")
	r.Memorize("circuit", "$", "circuit")
	return Circuit(r)
}

func generalizeIndex(x be.Index) be.Index {
	r := be.NewIndex()
	for n, v := range Circuit(x).Gate {
		switch t := v.(type) {
		case Circuit:
			if be.IsIndex(t) {
				Circuit(r).Include(n, Circuit(generalizeIndex(be.AsIndex(t))))
			} else {
				Circuit(r).Include(n, GeneralizeCircuit(t))
			}
		default: // retain all other values
			Circuit(r).Include(n, v)
		}
	}
	return r
}

func GeneralizeCircuit(u Circuit) Circuit {
	r := u.Copy()
	for n, v := range r.Gate {
		switch v.(type) {
		case int:
			r.Include(n, NewAddress("$", "int"))
		case float64:
			r.Include(n, NewAddress("$", "float"))
		case complex128:
			r.Include(n, NewAddress("$", "complex"))
		case string:
			r.Include(n, NewAddress("$", "string"))
		case bool:
			r.Include(n, NewAddress("$", "bool"))
		case Circuit:
			r.Include(n, NewAddress("$", "circuit"))
		case Address:
			r.Include(n, v)
		default:
			panic("unknown type")
		}
	}
	return r
}
