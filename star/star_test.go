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

	// Construct
	s := Make()
	s.Show(1)
	s = Traverse(s, "f1", "")
	s.Show(2)

	// Inverse construct
	r := Make()
	r.Show(2)
	r = Traverse(r, "", "f1")
	r.Show(1)
	r = Traverse(r, "f1", "")
	if !Same(s, r) {
		t.Errorf("mismatch")
	}

	// Criss-cross split-merge
	x, y := Split(s, "", "f1")
	a, b := Split(r, "", "f1")
	x.Merge("xx", "", b)
	a.Merge("xx", "", y)
	if !Same(x, a) {
		t.Errorf("mismatch")
	}
}
