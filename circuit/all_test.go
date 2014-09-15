// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"testing"
)

func TestSame(t *testing.T) {
	if !Same(New().Grow("x", nil), New().Grow("x", nil)) {
		t.Errorf("same")
	}
}
