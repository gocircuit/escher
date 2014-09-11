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
)

type Printer interface {
	Print(prefix, indent string) string
}

func (u *circuit) super() (super Name, ok bool) {
	for n, m := range u.Images() {
		if _, ok := m.(Super); ok {
			if super != nil {
				//fmt.Printf("X=%v\n", u.Images())
				panic("two supers")
			}
			super = n
		}
	}
	if super == nil {
		return nil, false
	}
	return super, true
}

func (u *circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	super, ok := u.super()
	if ok { // if super is present
		fmt.Fprintf(&w, "%v", super)
		valves := u.Valves(super)
		w.WriteString("(")
		if len(valves) > 0 {
			var i int
			for vn, _ := range valves {
				fmt.Fprintf(&w, "%v", vn)
				i++
				if i < len(valves) {
					w.WriteString(", ")
				}
			}
		}
		w.WriteString(") ")
	}
	w.WriteString("{\n")
	// letters
	for _, n := range u.Letters() {
		p := u.image[n]
		if n == super {
			continue
		}
		w.WriteString(prefix + indent)
		PrintMeaning(&w, prefix+indent, indent, n, p)
	}
	// numbers
	for _, n := range u.Numbers() {
		p := u.image[n]
		if n == super {
			continue
		}
		w.WriteString(prefix + indent)
		PrintMeaning(&w, prefix+indent, indent, n, p)
	}
	//
	o := make(Orient)
	for _, valves := range u.real {
		for _, re := range valves {
			if o.Has(re.Image[1], re.Valve[1]) {
				continue
			}
			o.Include(re.Image[0], re.Valve[0])
			//
			fmt.Fprintf(&w, "%s%s%s:%s = %s:%s\n", 
				prefix, indent,  
				re.Image[0], re.Valve[0],
				re.Image[1], re.Valve[1],
			)
		}
	}
	w.WriteString(prefix + "}")
	return w.String()
}

func PrintMeaning(w io.Writer, prefix, indent string, n Name, p Meaning) {
	switch t := p.(type) {
	case Printer:
		fmt.Fprintf(w, "%v %v\n", n, t.Print(prefix, indent))
	case string:
		fmt.Fprintf(w, "%v %q\n", n, t)
	default:
		fmt.Fprintf(w, "%v %v\n", n, t)
	}
}