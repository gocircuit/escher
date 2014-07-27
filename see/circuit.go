// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/gocircuit/escher/star"
)

func SeeCircuit(src *Src) (name string, x *star.Star) {
	return SeePeer(src)
}
