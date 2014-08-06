// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package star

import (
	"testing"
)

func TestStar(t *testing.T) {

	s := Make().Grow("1", 1i).Grow("2", 2i)
	r := Star{"1": 1i, "2": 2i}
	if !Same(s, r) {
		t.Errorf("mismatch")
	}
	s.Abandon("1")
	r.Abandon("1")
	if !Same(s, r) {
		t.Errorf("mismatch")
	}
}
