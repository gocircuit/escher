// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"fmt"
	"testing"
)

func TestSame(t *testing.T) {
	if !Same(New().Grow("x", nil), New().Grow("x", nil)) {
		t.Errorf("same")
	}
	if !Same(New().Grow("x", DefaultValve), New().Grow("x", DefaultValve)) {
		t.Errorf("same")
	}
}

func TestVerb(t *testing.T) {
	a, b := Circuit(NewLookupVerb("abc", "d", 1)), Circuit(NewLookupVerb("abc", "d", 1))
	if !Same(a, b) {
		t.Errorf("verb same")
	}
	fmt.Printf("%v\n", a)
}
