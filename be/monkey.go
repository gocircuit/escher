// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/see"
)

func filterMonkey(a Address) (addr Address, monkey bool) {
	a = a.Copy()
	if len(a.Path) == 0 {
		return a, false
	}
	f, ok := a.Path[0].(string)
	if !ok {
		return a, false
	}
	if len(f) == 0 {
		return a, false
	}
	if f[0] != '@' {
		return a, false
	}
	n := see.ParseName(f[1:]).(string) // parse name after @
	if n == "" {
		a.Path = a.Path[1:]
	} else {
		a.Path[0] = n
	}
	return a, true
}
