// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package star

import (
	"bytes"
	"fmt"
	"sort"
)

// Star is a node from a symmetric tree, i.e. a tree without a distinct root.
type Star map[string]interface{}

var NoStar = Star{}

// Make creates a singleton node star and an eye into it.
func Make() Star {
	return make(Star)
}

func (s Star) Unwrap() map[string]interface{} {
	return (map[string]interface{})(s)
}

// Copy returns a Star-recursive (non-Star children are not recursed into) copy of the star.
func (s Star) Copy() Star {
	t := Make()
	for key, value := range s {
		switch u := value.(type) {
		case Star:
			t.Grow(key, u.Copy())
		default:
			t.Grow(key, value)
		}
	}
	return t
}

// Len returns the number of stars.
func (s Star) Len() int {
	return len(s)
}

func (s Star) Grow(key string, value interface{}) Star {
	if _, present := s[key]; present {
		panic(4)
	}
	s[key] = value
	return s
}

func (s Star) Abandon(key string) Star {
	delete(s, key)
	return s
}

func (s Star) Walk(key string) Star {
	return s[key].(Star)
}

func (s Star) Take(key string) interface{} {
	return s[key]
}

func (s Star) String(key string) string {
	return s[key].(string)
}

func (s Star) Int(key string) int {
	return s[key].(int)
}

func (s Star) Float(key string) float64 {
	return s[key].(float64)
}

func (s Star) Complex(key string) complex128 {
	return s[key].(complex128)
}

// Sort returns the keys in s in sorted order.
func (s Star) Sort() []string {
	lex := make([]string, 0, len(s))
	for key, _ := range s {
		lex = append(lex, key)
	}
	sort.Strings(lex)
	return lex
}

func Same(s, t Star) bool {
	return s.Contains(t) && t.Contains(s)
}

func (s Star) Contains(t Star) bool {
	for key, tv := range t {
		sv, present := s[key]
		if !present {
			return false
		}
		switch tu := tv.(type) {
		case Star:
			su, ok := sv.(Star)
			if !ok || !su.Contains(tu) {
				return false
			}
		default:
			if sv != tv {
				return false
			}
		}
	}
	return true
}

type Printer interface {
	Print(prefix, indent string) string
}

func (s Star) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "{\n")
	for key, v := range s {
		var t string
		switch u := v.(type) {
		case Printer:
			t = u.Print(prefix+indent, indent)
		default:
			t = fmt.Sprintf("%v", v)
		}
		fmt.Fprintf(&w, "%s%s%s %s\n", prefix, indent, key, t)
	}
	fmt.Fprintf(&w, "%s}", prefix)
	return w.String()
}
