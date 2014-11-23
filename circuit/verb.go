// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	. "github.com/gocircuit/escher/a"
)

// DefaultValve
const DefaultValve = ""

// Verb is an interpretation of a circuit, which also has a shorthand syntax, e.g.
//	*http.Server
//	@http.Server
type Verb Circuit

func NewVerb(addr ...Name) Verb {
	x := New()
	for i, n := range addr {
		x.Gate[i] = n
	}
	return Verb(x)
}

func NewLookupVerb(addr ...Name) Verb {
	x := NewVerb(addr...)
	x.Gate[""] = "@"
	return x
}

func NewMaterializeVerb(addr ...Name) Verb {
	x := NewVerb(addr...)
	x.Gate[""] = "*"
	return x
}

func IsVerb(v Value) bool {
	u, ok := v.(Circuit)
	if !ok {
		return false
	}
	s, _ := u.StringOptionAt("")
	return s == "*" || s == "@"
}

func (a Verb) Address() (addr []Name) {
	for _, i := range Circuit(a).SortedNumbers() {
		addr = append(addr, a.Gate[i])
	}
	return
}

func (a Verb) Verb() Value {
	return a.Gate[""]
}

func (a Verb) Append(b Verb) Verb {
	c := Verb(Circuit(a).Copy())
	index := Circuit(b).SortedNumbers()
	for _, i := range index {
		c.Gate[len(index)+i] = b.Gate[i]
	}
	return c
}

func (a Verb) compactible() bool {
	for _, n := range Circuit(a).SortedNumbers() {
		s, ok := a.Gate[n].(string)
		if !ok {
			return false
		}
		if strings.IndexAny(s, "@*.\n") >= 0 {
			return false
		}
	}
	return true
}

func (a Verb) Print(w io.Writer, f Format) {
	if !a.compactible() {
		Circuit(a).Print(w, f)
		return
	}
	io.WriteString(w, a.summarize())
}

func (a Verb) String() string {
	if !a.compactible() {
		return Circuit(a).String()
	}
	return a.summarize()
}

func (a Verb) summarize() string {
	index := Circuit(a).SortedNumbers()
	var w bytes.Buffer
	if v, ok := a.Gate[""]; ok {
		w.WriteString(fmt.Sprintf("%v", v))
	}
	for _, i := range index {
		x := a.Gate[i]
		fmt.Fprintf(&w, "%v", x)
		if i+1 < len(index) {
			w.WriteString(RefineSymbolString)
		}
	}
	return w.String()
}
