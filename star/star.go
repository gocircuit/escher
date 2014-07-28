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
	Choice Choice
	Value interface{}
}

type Choice map[string]*Star

// Make creates a singleton node star and an eye into it.
func Make() *Star {
	return &Star{
		Choice: make(Choice),
	}
}

func (s *Star) scrub() {
	s.Choice = nil
	s.Value = nil
}

// Copy returns a complete copy of the star with the same point-of-view into it.
func (s *Star) Copy(exclude ...string) *Star {
	t := Make()
	t.Show(s.Interface())
	for fwd, choice := range s.Choice {
		if fwd == Parent || contains(exclude, fwd) {
			continue
		}
		_, rev := s.Reverse(fwd)
		t.Merge(fwd, choice.Copy(rev))
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
	t, ok := s.Choice[fwd]
	if !ok {
		return nil, ""
	}
	for rev, r := range t.Choice {
		if r == s {
			return t, rev
		}
	}
	panic(3)
}

const Parent = ""

func (s *Star) Merge(fwd string, t *Star) *Star {
	if _, ok := s.Choice[fwd]; ok {
		panic("forward clash")
	}
	if _, ok := t.Choice[Parent]; ok {
		panic("Parent clash")
	}
	if s.Value != nil && len(s.Choice) > 0 {
		panic(1)
	}
	if t.Value != nil && len(t.Choice) > 0 {
		panic(1)
	}
	s.Choice[fwd], t.Choice[Parent] = t, s
	return s
}

// Point-of-view

func (s *Star) Grow(fwd string, value interface{}) *Star {
	if value == nil {
		panic(1)
	}
	if _, ok := value.(*Star); ok {
		panic(2)
	}
	s.Merge(fwd, Make().Show(value))
	return s
}

func (s *Star) Up() (t *Star) {
	defer s.collect()
	t, _ = s.Reverse(Parent)
	return t
}

func (s *Star) collect() {
	if s.Value != nil {
		return
	}
	if len(s.Choice) != 1 {
		return
	}
	for fwd, _ := range s.Choice {
		if fwd != Parent {
			panic(4)
		}
		t, rev := s.Reverse(fwd)
		Split(t, rev)
		s.scrub()
		return
	}
	panic(1)
}

func (s *Star) Down(fwd string) (t *Star) {
	t, _ = s.Reverse(fwd)
	if t != nil {
		return t
	}
	if fwd == Parent { // if at root
		return nil
	}
	t = Make()
	s.Merge(fwd, t)
	return t
}

func Split(s *Star, fwd string) (parent, child *Star) {
	t, rev := s.Reverse(fwd)
	delete(t.Choice, rev)
	delete(s.Choice, fwd)
	return s, t
}

// See returns the value stored at this node.
func (s *Star) Interface() interface{} {
	return s.Value
}

func (s *Star) String() string {
	return s.Value.(string)
}

func (s *Star) Int() int {
	return s.Value.(int)
}

func (s *Star) Float() float64 {
	return s.Value.(float64)
}

func (s *Star) Complex() complex128 {
	return s.Value.(complex128)
}

func (s *Star) Star() *Star {
	return s.Value.(*Star)
}

// Show sets the value stored at this node.
func (s *Star) Show(v interface{}) *Star {
	if v != nil && len(s.Choice) > 1 {
		panic("value in a non-terminal node")
	}
	s.Value = v
	return s
}

// Comparison

func Same(s, t *Star) bool {
	return s.Contains(t) && t.Contains(s)
}

func SameValue(x, y interface{}) bool {
	// star values should only be basic Go types that are directly comparable with ==
	if _, ok := x.(*Star); ok {
		panic(1)
	}
	if _, ok := y.(*Star); ok {
		panic(1)
	}
	return x == y
}

func (s *Star) Contains(t *Star, exclude ...string) bool {
	if !SameValue(s.Value, t.Value) {
		return false
	}
	for tfwd, tchoice := range t.Choice {
		if tfwd == Parent || contains(exclude, tfwd) {
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
	if s.Value != nil {
		fmt.Fprintf(&w, "%v", s.Value)
	} else {
		fmt.Fprintf(&w, "{\n")
		for fwd, choice := range s.Choice {
			if fwd == Parent || contains(exclude, fwd) {
				continue
			}
			_, rev := s.Reverse(fwd)
			fmt.Fprintf(
				&w, "%s%s%s %s\n", 
				prefix, indent, fwd,
				choice.Print(prefix+indent, indent, rev),
			)
		}
		fmt.Fprintf(&w, "%s}", prefix)
	}
	return w.String()
}
