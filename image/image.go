// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package image

import (
	"bytes"
	"fmt"
	"sort"
)

// Image is…
type Image map[string]interface{}

// Make creates a singleton node star and an eye into it.
func Make() Image {
	return make(Image)
}

func (x Image) Unwrap() map[string]interface{} {
	return (map[string]interface{})(x)
}

// Copy returns a Image-recursive (non-Image children are not recursed into) copy of the star.
func (x Image) Copy() Image {
	t := Make()
	for key, v := range x {
		switch u := v.(type) {
		case Image:
			t.Grow(key, u.Copy())
		default:
			t.Grow(key, v)
		}
	}
	return t
}

// Len returns the number of stars.
func (x Image) Len() int {
	return len(x)
}

func (x Image) Attach(y Image) Image {
	for key, v := range y {
		if _, present := x[key]; present {
			panic(1)
		}
		x[key] = v
	}
	return x
}

func (x Image) Grow(key string, v interface{}) Image {
	if _, present := x[key]; present {
		panic(4)
	}
	x[key] = v
	return x
}

func (x Image) Abandon(key string) Image {
	delete(x, key)
	return x
}

func (x Image) Walk(key string) Image {
	v := x[key]
	if v != nil {
		return v.(Image)
	}
	return Image{}
}

func (x Image) Has(key string) bool {
	_, present := x[key]
	return present
}

func (x Image) String(key string) string {
	return x[key].(string)
}

func (x Image) Int(key string) int {
	return x[key].(int)
}

func (x Image) Float(key string) float64 {
	return x[key].(float64)
}

func (x Image) Complex(key string) complex128 {
	return x[key].(complex128)
}

// Sort returns the keys in s in sorted order.
func (x Image) Sort() []string {
	lex := make([]string, 0, len(x))
	for key, _ := range x {
		lex = append(lex, key)
	}
	sort.Strings(lex)
	return lex
}

func Same(s, t Image) bool {
	return s.Contains(t) && t.Contains(s)
}

func (x Image) Contains(t Image) bool {
	for key, tv := range t {
		sv, present := x[key]
		if !present {
			return false
		}
		switch tu := tv.(type) {
		case Image:
			su, ok := sv.(Image)
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

func (x Image) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "{\n")
	var keys []string
	for k, _ := range x {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		v := x[key]
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

func linearize(s string) string {
	x := []byte(s)
	for i, b := range x {
		if b == '\n' {
			x[i] = ';'
		}
		if b == '\t' {
			x[i] = ' '
		}
	}
	return string(x)
}
