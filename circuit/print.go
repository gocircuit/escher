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

func (u *circuit) super() (super Name) {
	for n, m := range u.Symbols() {
		if _, ok := m.(Super); ok {
			if super != nil {
				panic("two supers")
			}
			super = n
		}
	}
	if super == nil {
		panic("no super")
	}
	return
}

func (u *circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	super := u.super()
	if super != nil {
		fmt.Fprintf(&w, "%v ", super)
	}
	valves := u.Valves(super)
	if len(valves) > 0 {
		w.WriteString("(")
		var i int
		for vn, _ := range valves {
			fmt.Fprintf(&w, "%v", vn)
			i++
			if i < len(valves) {
				w.WriteString(", ")
			}
		}
		w.WriteString(") ")
	}
	w.WriteString("{\n")
	for n, p := range u.Symbols() {
		if n == super {
			continue
		}
		w.WriteString(prefix)
		w.WriteString(indent)
		switch t := p.(type) {
		case Circuit:
			fmt.Fprintf(&w, "%v %v\n", n, t.Print(prefix + indent, indent))
		case string:
			fmt.Fprintf(&w, "%v %q\n", n, t)
		default:
			fmt.Fprintf(&w, "%v %v\n", n, t)
		}
	}
	w.WriteString("}")
	return w.String()
}
