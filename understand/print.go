// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package understand

import (
	"bytes"
	"fmt"

	"github.com/gocircuit/escher/see"
)

func (x *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s%s {\n", prefix, x.Name)
	for _, p := range x.Peer {
		if p.Design == nil {
			p.Design = see.NameDesign("☻")
		}
		fmt.Fprintf(&w,"%s%s%s %s\n", prefix, indent, printable(p.Name), p.Design.String())
		for _, v := range p.Valve {
			fmt.Fprintf(&w, "%s%s%s%s·%s = %s·%s\n", 
				prefix, indent, indent, 
				printable(p.Name), printable(v.Name), 
				printable(v.Matching.Of.Name), printable(v.Matching.Name),
			)
		}
	}
	fmt.Fprintf(&w, "%s}\n", prefix)
	return w.String()
}

func printable(s string) string {
	if s != "" {
		return s
	}
	return "¶"
}