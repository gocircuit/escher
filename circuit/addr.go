// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"bytes"
	"fmt"
)

// DefaultValve
const DefaultValve = ""

// Address ...
type Address struct {
	Path []Name
}

func NewAddress(name ...Name) Address {
	return Address{name}
}

func (a Address) Chop() ([]Name, Name) {
	p := a.Path
	return p[:len(p)-1], p[len(p)-1]
}

func (a Address) Append(b Address) Address {
	c := Address{make([]Name, len(a.Path) + len(b.Path))}
	n := copy(c.Path, a.Path)
	m := copy(c.Path[n:], b.Path)
	if n != len(a.Path) || m != len(b.Path) {
		panic(0)
	}
	return c
}

func (a Address) Same(r Value) bool {
	b, ok := r.(Address)
	if !ok {
		return false
	}
	if len(a.Path) != len(b.Path) {
		return false
	}
	for i, j := range a.Path {
		if !Same(j, b.Path[i]) {
			return false
		}
	}
	return true
}

func (a Address) Copy() Address {
	b := Address{make([]Name, len(a.Path))}
	copy(b.Path, a.Path)
	return b
}

func (a Address) Simplify() interface{} {
	if len(a.Path) == 1 {
		return a.Simple()
	}
	return a
}

func (a Address) Simple() string {
	if len(a.Path) != 1 {
		panic("address not simple")
	}
	return a.Path[0].(string)
}

func (a Address) String() string {
	var w bytes.Buffer
	for i, x := range a.Path {
		fmt.Fprintf(&w, "%v", x)
		if i + 1 < len(a.Path) {
			w.WriteString(".")
		}
	}
	return w.String()
}

func (a Address) Strings() []string {
	var s = make([]string, len(a.Path))
	for i, x := range a.Path {
		s[i] = x.(string)
	}
	return s
}

func (a Address) Name() string {
	return a.Path[len(a.Path)-1].(string)
}

func (a Address) Circuit() Circuit {
	x := New()
	for i, j := range a.Path {
		x.Grow(i, j)
	}
	return x
}
