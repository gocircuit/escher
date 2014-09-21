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

// DefaultValve
const DefaultValve = ""

// Address ...
type Address struct {
	*address
}

type address []Name

func NewAddressStrings(ss []string) (a Address) {
	p := make(address, len(ss))
	for i, x := range ss {
		p[i] = x
	}
	return Address{&p}
}

func NewAddress(nn []Name) (a Address) {
	p := make(address, len(nn))
	copy(p, nn)
	return Address{&p}
}

func (a Address) Same(r Reducible) bool {
	b, ok := r.(Address)
	if !ok {
		return false
	}
	return a.address.Same(*b.address)
}

func (a address) Same(b address) bool {
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
	b := a.address.Copy()
	return Address{&b}
}

func (a address) Copy() address {
	x := make([]Name, len(a))
	copy(x, a)
	return x
}

func (a Address) Simplify() interface{} {
	if len(*a.address) == 1 {
		return a.Simple()
	}
	return a
}

func (a address) Simple() string {
	if len(a) != 1 {
		panic("address not simple")
	}
	return a[0].(string)
}

func (a address) String() string {
	var w bytes.Buffer
	for i, x := range a {
		fmt.Fprintf(&w, "%v", x)
		if i + 1 < len(a) {
			w.WriteString(".")
		}
	}
	return w.String()
}

func (a address) Strings() []string {
	var s = make([]string, len(a))
	for i, x := range a {
		s[i] = x.(string)
	}
	return s
}

func (a address) Path() []Name {
	return []Name(a)
}

func (a address) Name() string {
	return a[len(a)-1].(string)
}
