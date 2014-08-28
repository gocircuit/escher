// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	"bytes"
	"fmt"
	"sort"

	. "github.com/gocircuit/escher/image"
)

type printer interface {
	Print(string, string) string
}

func (fty Faculty) Print(prefix, indent string) string {
	var w bytes.Buffer
	var keys []string
	for k, _ := range fty {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := fty[k]
		w.WriteString(prefix)
		w.WriteString(indent)
		switch v.(type) {
		case Faculty:
		default:
			w.WriteString("*")
		}
		w.WriteString(k)
		switch t := v.(type) {
		case Faculty:
			w.WriteString(":\n")
			w.WriteString(t.Print(prefix+indent+indent, indent))
		case *Circuit:
			w.WriteString(" […]\n")
			// w.WriteString("\n")
			// w.WriteString(t.Print(prefix+indent+indent, indent))
		default: // reflex or circuit
			w.WriteString(fmt.Sprintf(" [%T]\n", v))
		}
	}
	return w.String()
}

func (x *Circuit) printValves() string {
	valve := x.Peer[""].Valve
	if len(valve) == 0 {
		return ""
	}
	var w bytes.Buffer
	w.WriteString("(")
	var i int
	for v, _ := range valve {
		w.WriteString(v)
		i++
		if i < len(valve) {
			w.WriteString(", ")
		}
	}
	w.WriteString(")")
	return w.String()
}

func (x *Circuit) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s%s {\n", x.Name, x.printValves())
	// string-named peers
	for _, p := range x.Peer {
		name, ok := p.Name.(string)
		if !ok {
			continue
		}
		if name == "" {
			continue
		}
		if pp, ok := p.Design.(Printer); ok {
			fmt.Fprintf(&w, "%s%s%s %v\n", prefix, indent, printable(name), pp.Print(prefix + indent, indent))
		} else {
			fmt.Fprintf(&w, "%s%s%s %v\n", prefix, indent, printable(name), p.Design)
		}
		for _, v := range p.Valve {
			fmt.Fprintf(&w, "%s%s%s%s.%s = %s.%s\n",
				prefix, indent, indent,
				printable(name), printable(v.Name),
				printable(v.Matching.Of.Name.(string)), printable(v.Matching.Name),
			)
		}
	}
	// int-named peers
	for _, p := range x.Peer {
		name, ok := p.Name.(int)
		if !ok {
			continue
		}
		fmt.Fprintf(&w, "%s%s#%d %v\n", prefix, indent, name, p.Design)
		for range p.Valve {
			panic(1)
		}
	}
	//
	fmt.Fprintf(&w, "%s}\n", prefix)
	return w.String()
}

func printable(s string) string {
	if s != "" {
		return s
	}
	return "@"
}
