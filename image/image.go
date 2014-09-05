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

// Image is ...
type Image map[interface{}]interface{}

// Make creates a singleton node star and an eye into it.
func Make() Image {
	return make(Image)
}

func (x Image) Unwrap() map[interface{}]interface{} {
	return (map[interface{}]interface{})(x)
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
			panic(2)
		}
		x[key] = v
	}
	return x
}

func (x Image) Grow(key interface{}, v interface{}) Image {
	if _, present := x[key]; present {
		panic(4)
	}
	x[key] = v
	return x
}

func (x Image) Abandon(key interface{}) Image {
	delete(x, key)
	return x
}

func (x Image) Cut(key interface{}) interface{} {
	v := x[key]
	delete(x, key)
	return v
}

func (x Image) Walk(key interface{}) Image {
	v := x[key]
	if v != nil {
		return v.(Image)
	}
	return Image{}
}

func (x Image) Has(key interface{}) bool {
	_, present := x[key]
	return present
}

func (x Image) Interface(key interface{}) interface{} {
	v, ok := x[key]
	if !ok {
		panic(3)
	}
	return v
}

func (x Image) OptionalInterface(key interface{}) interface{} {
	return x[key]
}

func (x Image) String(key interface{}) string {
	return x[key].(string)
}

func (x Image) OptionalString(key interface{}) string {
	v, ok := x[key]
	if !ok {
		return ""
	}
	return v.(string)
}

func (x Image) Int(key interface{}) int {
	return x[key].(int)
}

func (x Image) OptionalInt(key interface{}) int {
	v, ok := x[key]
	if !ok {
		return 0
	}
	return v.(int)
}

func (x Image) Float(key interface{}) float64 {
	return x[key].(float64)
}

func (x Image) Complex(key interface{}) complex128 {
	return x[key].(complex128)
}

func (x Image) Names() []interface{} {
	names := make([]interface{}, 0, len(x))
	for k, _ := range x {
		names = append(names, k)
	}
	return names
}

// Letters returns the string keys in s in sorted order.
func (x Image) Letters() []string {
	lex := make([]string, 0, len(x))
	for key, _ := range x {
		k, ok := key.(string)
		if !ok {
			continue
		}
		lex = append(lex, k)
	}
	sort.Strings(lex)
	return lex
}

func (x Image) Numbers() []int {
	series := make([]int, 0, len(x))
	for key, _ := range x {
		k, ok := key.(int)
		if !ok {
			continue
		}
		series = append(series, k)
	}
	sort.Ints(series)
	return series
}

func (x Image) LetterValues() (v []interface{}) {
	for _, n := range x.Letters() {
		v = append(v, x[n])
	}
	return
}

func (x Image) NumberValues() (v []interface{}) {
	for _, n := range x.Numbers() {
		v = append(v, x[n])
	}
	return
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

type Pretty Image

func (x Pretty) String() string {
	return Image(x).Print("", "\t")
}

func (x Image) PrintLine() string {
	return Linearize(x.Print("", ""))
}

func (x Image) Print(prefix, indent string) string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "{\n")
	// fmt.Fprintf(&w, "%s%s// letters\n", prefix, indent)
	for _, key := range x.Letters() {
		v := x[key]
		var t string
		switch u := v.(type) {
		case Printer:
			t = u.Print(prefix+indent, indent)
		case string:
			t = fmt.Sprintf("%q", u)
		default:
			t = fmt.Sprintf("%v", v)
		}
		fmt.Fprintf(&w, "%s%s%s %s\n", prefix, indent, key, t)
	}
	// fmt.Fprintf(&w, "%s%s// numbers\n", prefix, indent)
	for _, key := range x.Numbers() {
		v := x[key]
		var t string
		switch u := v.(type) {
		case Printer:
			t = u.Print(prefix+indent, indent)
		case string:
			t = fmt.Sprintf("%q", u)
		default:
			t = fmt.Sprintf("%v", v)
		}
		fmt.Fprintf(&w, "%s%s#%v %s\n", prefix, indent, key, t)
	}
	fmt.Fprintf(&w, "%s}", prefix)
	return w.String()
}

func Linearize(s string) string {
	x := []byte(s)
	for i, b := range x {
		if b == '\n' {
			x[i] = ','
		}
		if b == '\t' {
			x[i] = ' '
		}
	}
	return string(x)
}
