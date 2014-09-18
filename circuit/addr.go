// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"bytes"
	"fmt"
	// "strings"
)

// Address ...
type Address []Name

func NewAddressStrings(ss []string) (a Address) {
	a = make(Address, len(ss))
	for i, x := range ss {
		a[i] = x
	}
	return a
}

func (a Address) Same(r Reducible) bool {
	b, ok := r.(Address)
	if !ok {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, j := range a {
		if !Same(j, b[i]) {
			return false
		}
	}
	return true
}

func (a Address) Copy() Reducible {
	c := make(Address, len(a))
	copy(c, a)
	return c
}

func (a Address) Simplify() interface{} {
	if len(a) == 1 {
		return a.Simple()
	}
	return a
}

func (a Address) Simple() string {
	if len(a) != 1 {
		panic("address not simple")
	}
	return a[0].(string)
}

func (a Address) String() string {
	var w bytes.Buffer
	for i, x := range a {
		fmt.Fprintf(&w, "%v", x)
		if i + 1 < len(a) {
			w.WriteString(".")
		}
	}
	return w.String()
}

func (a Address) Strings() []string {
	var s = make([]string, len(a))
	for i, x := range a {
		s[i] = x.(string)
	}
	return s
}

func (a Address) Path() (walk []Name) {
	return []Name(a)
}

func (a Address) Name() string {
	return a[len(a)-1].(string)
}
