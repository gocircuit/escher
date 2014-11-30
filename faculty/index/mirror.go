// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package index

import (
	"fmt"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
)

// The Mirror gate recursively transforms an input circuit into one wherein
// (a) every terminal gate value (a constant or a verb circuit) is substituted
// by a noun materializer whose reflex returns the terminal gate value
// (b) whereas every materializer value is substituted by a noun materializer,
// which returns its own address within the index.
type Mirror struct{ be.Sparkless }

func (Mirror) CognizeIndex(eye *be.Eye, v interface{}) {
	eye.Show(DefaultValve, MirrorIndex(v.(Circuit), nil))
}

func (Mirror) Cognize(eye *be.Eye, v interface{}) {}

func MirrorIndex(u Circuit, addr []Name) Circuit {
	r := New()
	for n, v := range Circuit(u).Gate {
		switch t := v.(type) {
		case Circuit:
			r.Include(n, MirrorIndex(t, append(addr, n)))
		case be.Materializer:
			r.Include(n, be.NewSource(NewAddress(append(addr, n)...)))
		case int:
			r.Include(n, be.NewSource(NewAddress("int")))
		case float64:
			r.Include(n, be.NewSource(NewAddress("float")))
		case complex128:
			r.Include(n, be.NewSource(NewAddress("complex")))
		case string:
			r.Include(n, be.NewSource(NewAddress("string")))
		default:
			r.Include(n, be.NewSource(NewAddress("go", fmt.Sprintf("%T", t))))
		}
	}
	return r
}
