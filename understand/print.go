// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	"bytes"
	"fmt"

	. "github.com/gocircuit/escher/image"
)

type printer interface {
	Print(string, string) string
}

func (fty Faculty) Print(prefix, indent string) string {
	var w bytes.Buffer
	sd := fty.Genus().SourceDir
	w.WriteString("{ ")
	for _, acid := range sd.Letters() {
		w.WriteString(acid)
		w.WriteString("=")
		w.WriteString(sd.String(acid))
		w.WriteString(" ")
	}
	w.WriteString("}")
	keys := Image(fty).Letters()
	for _, k := range keys {
		v := fty[k]
		w.WriteString("\n" + prefix + indent)
		//
		switch v.(type) {
		case *Circuit:
			w.WriteString("*")
		}
		w.WriteString(k)
		switch v.(type) {
		case Faculty:
			// w.WriteString(":")
		}
		//
		switch t := v.(type) {
		case Faculty:
			w.WriteString(" ")
			w.WriteString(t.Print(prefix + indent, indent))
		case *Circuit:
			w.WriteString(" (…)")
			// w.WriteString("\n")
			// w.WriteString(t.Print(prefix+indent+indent, indent))
		default: // reflex or circuit
			w.WriteString(fmt.Sprintf(" [%T]", v))
		}
	}
	return w.String()
}

func (x *Circuit) printValves() string {
	sup := x.PeerByName("")
	vnames := sup.ValveNames()
	if len(vnames) == 0 {
		return ""
	}
	var w bytes.Buffer
	w.WriteString("(")
	var i int
	for _, vn := range vnames {
		w.WriteString(vn)
		i++
		if i < len(vnames) {
			w.WriteString(", ")
		}
	}
	w.WriteString(")")
	return w.String()
}

func (x *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s%s {\n", x.Name(), x.printValves())
	// string-named peers
	for _, name_ := range x.PeerNames() {
		p := x.PeerByName(name_)
		if name_ == "" {
			continue
		}
		name := nonemptify(print(name_))
		if pp, ok := p.Design().(Printer); ok {
			fmt.Fprintf(&w, "%s%s%s %v\n", prefix, indent, name, pp.Print(prefix + indent, indent))
		} else {
			fmt.Fprintf(&w, "%s%s%s %v\n", prefix, indent, name, p.Design())
		}
		for _, vn := range p.ValveNames() {
			v := p.ValveByName(vn)
			fmt.Fprintf(&w, "%s%s%s%s.%s = %s.%s\n",
				prefix, indent, indent,
				name, nonemptify(vn),
				nonemptify(print(v.Matching.Of.Name())), nonemptify(v.Matching.Name),
			)
		}
	}
	// int-named peers
	// for _, p := range x.Peer {
	// 	name, ok := p.Name.(int)
	// 	if !ok {
	// 		continue
	// 	}
	// 	fmt.Fprintf(&w, "%s%s#%d %v\n", prefix, indent, name, p.Design)
	// 	for _ = range p.Valve {
	// 		panic(1)
	// 	}
	// }
	//
	fmt.Fprintf(&w, "%s}", prefix)
	return w.String()
}

func print(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func nonemptify(s string) string {
	if s != "" {
		return s
	}
	return "@"
}
