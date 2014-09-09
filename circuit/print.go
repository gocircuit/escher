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
	for n, m := range u.Images() {
		if _, ok := m.(Super); ok {
			if super != nil {
				println("alrea", fmt.Sprintf("%v vs %v", super, n))
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
	for n, p := range u.image {
		if n == super {
			continue
		}
		w.WriteString(prefix)
		w.WriteString(indent)
		switch t := p.(type) {
		case Circuit:
			fmt.Fprintf(&w, "%v\n", t.Print(prefix + indent, indent))
		case string:
			fmt.Fprintf(&w, "%v %q\n", n, t)
		default:
			fmt.Fprintf(&w, "%v %v\n", n, t)
		}
		for _, m := range u.real[n] { // links
			fmt.Fprintf(&w, "%s%s%s:%s = %s:%s\n", 
				prefix, indent,  
				m.Image[0], m.Valve[0],
				m.Image[1], m.Valve[1],
			)
		}
	}
	w.WriteString(prefix + "}")
	return w.String()
}
