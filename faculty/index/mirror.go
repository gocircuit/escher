// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	// "fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

type Mirror struct {}

func (Mirror) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Mirror) CognizeIndex(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, Circuit(MirrorNative(be.AsIndex(v), nil)))
}

func (Mirror) Cognize(eye *be.Eye, v interface{}) {}

func MirrorNative(u be.Index, path []Name) be.Index {
	r := be.NewIndex()
	for n, v := range Circuit(u).Gate {
		// println(fmt.Sprintf("-> %v:%T", n, v))
		if n == "?" {
			continue
		}
		if IsSymbol(n) {
			Circuit(r).Include(n, v)
		} else {
			switch t := v.(type) {
			case Circuit:
				if be.IsIndex(t) {
					Circuit(r).Include(n, 
						Circuit(
							MirrorNative(be.AsIndex(t), append(path, n)),
						),
					)
				} else {
					Circuit(r).Include(n, v)
				}
			default:
				Circuit(r).Include(n, be.NewNoun(Address{append(path, n)}))
			}
		}
	}
	return r
}
