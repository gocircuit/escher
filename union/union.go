// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package union

import (
	"bytes"
	"fmt"

	// . "github.com/gocircuit/escher/image"
)

// Name is one of: int or string
type Name interface{}

// Meaning is one of: string, int, float64, complex128, *Union
type Meaning interface{}

// Union ...
type Union struct {
	peer map[Name]Meaning
	match map[Name]map[Name]Matching // peer -> valve -> opposing peer and valve
}

// Super is a placeholder meaning for the super peer
type Super struct{}

// Matching ...
type Matching struct {
	Peer [2]Name
	Valve [2]Name
}

func (x Matching) Reverse() Matching {
	x.Peer[0], x.Peer[1] = x.Peer[1], x.Peer[0]
	x.Valve[0], x.Valve[1] = x.Valve[1], x.Valve[0]
	return x
}

// New ...
func New() *Union {
	return &Union{
		peer: make(map[Name]Meaning),
		match: make(map[Name]map[Name]Matching),
	}
}

// Add adds a peer to this union.
func (c *Union) Add(name Name, meaning Meaning) {
	c.peer[name] = meaning
}

// Peer ...
func (c *Union) Peer(name Name) Meaning {
	return c.peer[name]
}

// Match ...
func (c *Union) Match(x Matching) {
	if x.Peer[0] == x.Peer[1] && x.Valve[0] == x.Valve[1] {
		panic("mismatch")
	}
	p := []map[Name]Matching{
		c.valves(x.Peer[0]), 
		c.valves(x.Peer[1]),
	}
	v := x.Valve
	if _, ok := p[0][v[0]]; ok {
		panic("dup")
	}
	if _, ok := p[1][v[1]]; ok {
		panic("dup")
	}
	p[0][v[0]], p[1][v[1]] = x, x.Reverse()
}

func (c *Union) valves(p Name) map[Name]Matching {
	if c.match[p] == nil {
		c.match[p] = make(map[Name]Matching)
	}
	return c.match[p]
}

func (u *Union) Valves(peer Name) map[Name]Matching {
	return u.match[peer]
}

// Follow ...
func (c *Union) Follow(p, v Name) (q, u Name) {
	x, ok := c.valves(p)[v]
	if !ok {
		return nil, nil
	}
	return x.Peer[1], x.Valve[1]
}

func (c *Union) Letters() []string {
	var l []string
	for key, _ := range c.peer {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

func (c *Union) Numbers() []int {
	var l []int
	for key, _ := range c.peer {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

func (u *Union) Peers() map[Name]Meaning {
	return u.peer
}

func (u *Union) String() string {
	return u.Print(nil, "", "\t")
}

func (u *Union) Print(super Name, prefix, indent string) string {
	var w bytes.Buffer
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
	for n, p := range u.Peers() {
		if n == super {
			continue
		}
		w.WriteString(prefix)
		w.WriteString(indent)
		switch t := p.(type) {
		case *Union:
			fmt.Fprintf(&w, "%v %v\n", n, t.Print(n, prefix + indent, indent))
		case string:
			fmt.Fprintf(&w, "%v %q\n", n, t)
		default:
			fmt.Fprintf(&w, "%v %v\n", n, t)
		}
	}
	w.WriteString("}")
	return w.String()
}
