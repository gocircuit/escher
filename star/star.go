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

// Star is a node from a symmetric tree, i.e. a tree without a distinct root.
type Star struct {
	choice map[string]*Star
	value interface{}
}

// Make creates a singleton node star and an eye into it.
func Make() *Star {
	return &Star{
		choice: make(map[string]*Star),
	}
}

func (s *Star) scrub() {
	s.choice = nil
	s.value = nil
}

// Copy returns a complete copy of the star with the same point-of-view into it.
func (s *Star) Copy(exclude ...string) *Star {
	t := Make()
	t.Show(s.Interface())
	for fwd, choice := range s.choice {
		if contains(exclude, fwd) {
			continue
		}
		_, rev := s.Reverse(fwd)
		t.Merge(fwd, rev, choice.Copy(rev))
	}
	return t
}

func contains(set []string, s string) bool {
	for _, x := range set {
		if x == s {
			return true
		}
	}
	return false
}

// Reverse returns the name of the choice on fwd that points back to s.
func (s *Star) Reverse(fwd string) (*Star, string) {
	t, ok := s.choice[fwd]
	if !ok {
		return nil, ""
	}
	for rev, r := range t.choice {
		if r == s {
			return t, rev
		}
	}
	panic(3)
}

func (s *Star) Merge(fwd, rev string, t *Star) *Star {
	if _, ok := s.choice[fwd]; ok {
		panic("forward clash")
	}
	if _, ok := t.choice[rev]; ok {
		panic("reverse clash")
	}
	s.choice[fwd], t.choice[rev] = t, s
	return s
}

// Point-of-view

// Traverse gives a different point-of-view on the same star, by moving the current rootcalong the branch labeled name.
func (s *Star) Traverse(fwd, rev string) (t *Star) {
	defer func() {
		if s.value == nil && len(s.choice) == 0 { // garbage-collect behind us
			t.Split(rev, fwd)
			s.scrub()
		}
	}()
	var trev string
	t, trev = s.Reverse(fwd)
	if t != nil {
		if rev != trev {
			panic("unintended reverse")
		}
		return t
	}
	return s.Merge(fwd, rev, Make())
}

func (s *Star) Split(fwd, rev string) (parent, child *Star) {
	t, trev := s.Reverse(fwd)
	if trev != rev {
		panic("unintended reverse")
	}
	delete(t.choice, rev)
	delete(s.choice, fwd)
	return s, t
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
	// star values should only be basic Go types that are directly comparable
	return x == y
}

func (s *Star) Contains(t *Star, exclude ...string) bool {
	if !SameValue(s.value, t.value) {
		return false
	}
	for tfwd, tchoice := range t.choice {
		if contains(exclude, tfwd) {
			continue
		}
		_, trev := t.Reverse(tfwd)
		schoice, srev := s.Reverse(tfwd)
		if schoice == nil || srev != trev {
			return false
		}
		if !schoice.Contains(tchoice, trev) {
			return false
		}
	}
	return true
}

func (s *Star) Print(prefix, indent string, exclude ...string) string {
	var w bytes.Buffer
	var value string
	if s.value != nil {
		value = " *"
	}
	fmt.Fprintf(&w, "%s%s{\n", prefix, value)
	for fwd, choice := range s.choice {
		if contains(exclude, fwd) {
			continue
		}
		_, rev := s.Reverse(fwd)
		fmt.Fprintf(
			&w, "%s%s%s\\%s %s\n", 
			prefix, indent, fwd, rev,
			choice.Print(prefix+indent, indent, rev),
		)
	}
	fmt.Fprintf(&w, "%s}", prefix)
	return w.String()
}
