// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"math"
)

func AsInt(v interface{}) (int, bool) {
	if i, ok := v.(int); ok {
		return i, true
	}
	if f, ok := v.(float64); ok && math.Floor(f) == f {
		return int(f), true
	}
	return 0, false
}
