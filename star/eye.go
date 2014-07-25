// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package star

import (
	"bytes"
	"fmt"
)

// Eye…
type Eye struct {
	star *Star
}

// Creation

// Make creates a singleton node star.
func Make() *Eye {
	return &Eye{
		&Star{
			choice: make(map[string]*Star),
		},
	}
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

// Copy returns a complete copy of the star with the same point-of-view into it.
func (s *Star) Copy() *Star {
	defer unpebble(pebble(s))
	t := Make()
	t.Show(s.Interface())
	for name, choice := range s.choice {
		if choice.pebble {
			continue
		}
		t.Merge(name, choice.Copy())
	}
	return t
}

func (s *Star) Merge(name string, t *Star) *Star {
	if _, ok := s.choice[name]; ok {
		panic("clash")
	}
	s.choice[name] = t
	s.below += t.Weight()
	return s
}

// Value

// Number of non-nil value nodes in this star.
func (s *Star) Weight() int {
	if s.value != nil {
		return s.below+1
	}
	return s.below
}

// Point-of-view

// Traverse gives a different point-of-view on the same star, by moving the current rootcalong the branch labeled name.
func (s *Star) Traverse(forward, backward string) *Star {
	t, ok := s.choice[forward]
	if ok {
		if t.choice[backward] != s {
			panic("unintended traversal")
		}
		return t
	}
	t = Make()
	s.choice[forward] = t
	t.choice[backward] = s
	return t
}

func (s *Star) goto(t *Star) {
	??
}

func (s *Star) comefrom(t *Star) {
	??
}

func (s *Star) Split(name string) *Star {
	??
	t, ok := s.choice[name]
	if !ok {
		panic(8)
	}
	delete(s.choice, name)
	delete(t.choice, "")
	return s
}

// See returns the value stored at this node.
func (s *Star) Interface() interface{} {
	return s.value
}

func (s *Star) String() string {
	return s.value.(string)
}

func (s *Star) Int() int {
	return s.value.(int)
}

func (s *Star) Float() float64 {
	return s.value.(float64)
}

func (s *Star) Complex() complex128 {
	return s.value.(complex128)
}

func (s *Star) Star() *Star {
	return s.value.(*Star)
}

// Show sets the value stored at this node.
func (s *Star) Show(v interface{}) {
	s.value = v
}

// Comparison

func SameStar(s, t *Star) bool {
	return s.Contains(t) && t.Contains(s)
}

func SameValue(x, y interface{}) bool {
	return x == y
}

func (s *Star) Contains(t *Star) bool {
	defer unpebble(pebble(s))
	if !SameValue(s.value, t.value) {
		return false
	}
	for name, tchoice := range t.choice {
		if tchoice.pebble {
			continue
		}
		schoice, ok := s.choice[name]
		if !ok {
			return false
		}
		return schoice.Contains(choice)
	}
	return true
}

// Printing

func (s *Star) Print(prefix, indent string) string {
	defer unpebble(pebble(s))
	var w bytes.Buffer
	var value string
	if s.value != nil {
		value = " *"
	}
	fmt.Fprintf(&w, "%s%s{\n", prefix, value)
	for name, choice := range s.choice {
		if choice.pebble {
			continue
		}
		fmt.Fprintf(&w, "%s%s%s(%d) %s\n", prefix, indent, name, choice.Width(), choice.Print(prefix+indent, indent))
	}
	fmt.Fprintf(&w, "%s}", prefix)
	return w.String()
}
