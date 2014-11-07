// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package reflect

import (
	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Addresses struct {}

func (Addresses) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Addresses) CognizeIdiom(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, reflectAddresses(v.(be.Idiom), nil))
}

func (Addresses) Cognize(eye *be.Eye, v interface{}) {}

func reflectAddresses(u be.Idiom, path []Name) be.Idiom {
	r := be.NewIdiom()
	for n, v := range Circuit(u).Gate {
		switch t := v.(type) {
		case be.Idiom:
			Circuit(u).Include(n, reflectAddresses(t, append(path, n)))
		default:
			Circuit(u).Include(n, be.NewNoun(Address{append(path, n)}))
		}
	}
	return r
}
