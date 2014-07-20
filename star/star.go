// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package star

import (
	"bytes"
	"fmt"
)

type Star struct {
	star map[string]*Star
	value interface{}
	pebble bool
}

func pebble(s *Star) *bool {
	if s.pebble == true {
		panic(3)
	}
	s.pebble = true
	return &s.pebble
}

func unpebble(p *bool) {
	if !*p {
		panic(3)
	}
	*p = false
}

// Creation

// Make creates a singleton node star.
func Make() *Star {
	return &Star{
		star: make(map[string]*Star),
	}
}

// Copy returns a complete copy of the star with the same point-of-view into it.
func (s *Star) Copy() *Star {
	defer unpebble(pebble(s))
	t := Make()
	for name, ss := range s.star {
		if ss.pebble {
			continue
		}
		t.star[name] = ss.Copy()
	}
	return t
}

// Chane of point-of-view

// Traverse gives a different point-of-view on the same star, by moving the current root
// along the branch labeled name.
func (s *Star) Traverse(name string) *Star {
	return s.star[name]
}

// See returns the value stored at this node.
func (s *Star) See() interface{} {
	return s.value
}

// Show sets the value stored at this node.
func (s *Star) Show(v interface{}) {
	s.value = v
}

// Mutations

func (s *Star) Grow(name string) *Star {
	if _, ok := s.star[name]; ok {
		panic("name exists")
	}
	child := Make()
	s.star[name] = child
	child.star[""] = s
	return s
}

func (s *Star) Shrink(name string) *Star {
	t, ok := s.star[name]
	if !ok {
		panic(8)
	}
	delete(s.star, name)
	delete(t.star, "")
	return s
}

// Comparison

func Same(s, t *Star) bool {
	return s.Contains(t) && t.Contains(s)
}

func (s *Star) Contains(t *Star) bool {
	defer unpebble(pebble(s))
	for name, tt := range t.star {
		if tt.pebble { // maybe pebbling should be on t?
			continue
		}
		ss, ok := s.star[name]
		if !ok {
			return false
		}
		return ss.Contains(tt)
	}
	return true
}

// Printing

func (s *Star) Print(prefix, indent string, without ...string) string {
	defer unpebble(pebble(s))
	var w bytes.Buffer
	fmt.Fprintf(&w, "%s{\n", prefix)
	for name, ss := range s.star {
		if ss.pebble {
			continue
		}
		fmt.Fprintf(&w, "%s%s%s %s\n", prefix, indent, name, ss.Print(prefix+indent, indent, ""))
	}
	fmt.Fprintf(&w, "%s}", prefix)
	return w.String()
}
