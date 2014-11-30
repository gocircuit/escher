// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
)

func Panicf(f string, a ...interface{}) {
	panic(fmt.Sprintf(f, a...))
}
