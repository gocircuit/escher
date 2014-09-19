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
type Address struct {
	path []Name
}

func NewAddressStrings(ss []string) (a Address) {
	a = Address{
		path: make([]Name, len(ss)),
	}
	for i, x := range ss {
		a.path[i] = x
	}
	return
}

func (a Address) Same(r Reducible) bool {
	b, ok := r.(Address)
	if !ok {
		return false
	}
	if len(a.path) != len(b.path) {
		return false
	}
	for i, j := range a.path {
		if !Same(j, b.path[i]) {
			return false
		}
	}
	return true
}

func (a Address) Copy() Reducible {
	c := Address{
		path: make([]Name, len(a.path)),
	}
	copy(c.path, a.path)
	return c
}

func (a Address) Simplify() interface{} {
	if len(a.path) == 1 {
		return a.Simple()
	}
	return a
}

func (a Address) Simple() string {
	if len(a.path) != 1 {
		panic("address not simple")
	}
	return a.path[0].(string)
}

func (a Address) String() string {
	var w bytes.Buffer
	for i, x := range a.path {
		fmt.Fprintf(&w, "%v", x)
		if i + 1 < len(a.path) {
			w.WriteString(".")
		}
	}
	return w.String()
}

func (a Address) Strings() []string {
	var s = make([]string, len(a.path))
	for i, x := range a.path {
		s[i] = x.(string)
	}
	return s
}

func (a Address) Path() (walk []Name) {
	return []Name(a.path)
}

func (a Address) Name() string {
	return a.path[len(a.path)-1].(string)
}
