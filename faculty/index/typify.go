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

type Typify struct {}

func (Typify) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Typify) CognizeIndex(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, TypifyIndex(v.(be.Index)))
}

func (Typify) Cognize(eye *be.Eye, v interface{}) {}

func TypifyIndex(x be.Index) be.Index {
	r := typifyIndex(x)
	r.Memorize("int", "", "int")
	r.Memorize("float64", "", "float")
	r.Memorize("complex128", "", "complex")
	r.Memorize("string", "", "string")
	r.Memorize("bool", "", "bool")
	r.Memorize("circuit", "", "circuit")
	return r
}

func typifyIndex(x be.Index) be.Index {
	r := be.NewIndex()
	for n, v := range Circuit(x).Gate {
		switch t := v.(type) {
		case be.Index:
			Circuit(r).Include(n, typifyIndex(t))
		case Circuit:
			Circuit(r).Include(n, TypifyCircuit(t))
		default: // retain all other values
			Circuit(r).Include(n, v)
		}
	}
	return r
}

func TypifyCircuit(u Circuit) Circuit {
	r := u.Copy()
	for n, v := range r.Gate {
		switch v.(type) {
		case int:
			r.Include(n, NewAddress("", "int"))
		case float64:
			r.Include(n, NewAddress("", "float"))
		case complex128:
			r.Include(n, NewAddress("", "complex"))
		case string:
			r.Include(n, NewAddress("", "string"))
		case bool:
			r.Include(n, NewAddress("", "bool"))
		case Circuit:
			r.Include(n, NewAddress("", "circuit"))
		case Address:
			r.Include(n, v)
		default:
			panic("unknown type")
		}
	}
	return r
}
