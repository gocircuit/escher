// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	. "github.com/gocircuit/escher/circuit"
)

func SeeValue(src *Src) (x Value) {
	if x = SeeCircuit(src); x != nil {
		return
	}
	if x = SeeValueNoCircuit(src); x != nil {
		return
	}
	return nil
}

func SeeValueOrNil(src *Src) (x Value) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	return SeeValue(src)
}

func SeeValueNoCircuit(src *Src) (x Value) {
	if x = SeeInt(src); x != nil {
		return
	}
	if x = SeeFloat(src); x != nil {
		return
	}
	if x = SeeComplex(src); x != nil {
		return
	}
	if x = SeeBackquoteString(src); x != nil {
		return
	}
	if x = SeeDoubleQuoteString(src); x != nil {
		return
	}
	if x = SeeAddress(src); x != nil {
		return
	}
	return nil
}

func ParseName(src string) Name {
	t := NewSrcString(src)
	n := SeeName(t)
	if t.Len() != 0 {
		log.Fatalf("Non-name characters at end of %q", src)
	}
	return n
}

func SeeName(src *Src) Name {
	x := src.Consume(IsIdentifier)
	if x == "" { // empty string is allowed as name
		return x
	}
	i, err := strconv.Atoi(x) // recognize ints and return them as such
	if err == nil {
		return i
	}
	return x
}

func ParseAddress(src string) Address {
	t := NewSrcString(src)
	a := SeeAddress(t).(Address)
	if t.Len() != 0 {
		log.Printf("Non-address characters at end of %q", src)
		panic(1)
	}
	return a
}

// SeeAddress ...
func SeeAddress(src *Src) interface{} {
	t := src.Copy()
	delimit := t.Consume(IsIdentifierOrDot)
	x := strings.Split(delimit, ".")
	if len(x) == 0 {
		return nil
	}
	if len(x) == 1 && x[0] == "" {
		return nil
	}
	var addr Address
	for _, a := range x {
		addr.Path = append(addr.Path, ParseName(a))
	}
	src.Become(t)
	return addr
}

// Int …
func SeeInt(src *Src) interface{} {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return nil
	}
	r := bytes.NewBufferString(l)
	var i int
	if n, _ := fmt.Fscanf(r, "%d", &i); n != 1 || r.Len() != 0 {
		return nil
	}
	src.Become(t)
	return i
}

// Float …
func SeeFloat(src *Src) interface{} {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return nil
	}
	r := bytes.NewBufferString(l)
	var f float64
	if n, _ := fmt.Fscanf(r, "%g", &f); n != 1 || r.Len() != 0 {
		return nil
	}
	src.Become(t)
	return f
}

// Complex …
func SeeComplex(src *Src) interface{} {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return nil
	}
	r := bytes.NewBufferString(l)
	var c complex128
	if n, _ := fmt.Fscanf(r, "%g", &c); n != 1 || r.Len() != 0 {
		return nil
	}
	src.Become(t)
	return c
}

// SeeBackquoteString …
func SeeBackquoteString(src *Src) interface{} {
	t := src.Copy()
	quoted, ok := DelimitBackquoteString(t)
	if !ok {
		return nil
	}
	str := quoted[1 : len(quoted)-1]
	src.Become(t)
	return str
}

func DelimitBackquoteString(src *Src) (string, bool) {
	var m int // number of bytes accepted into the quoted portion
	buf := src.Buffer()
	// first backquote
	r, n, err := buf.ReadRune()
	if err != nil || r != '`' {
		return "", false
	}
	m += n
	//
	var q int // number of backquotes right behind cursor, not counting opening backquote
	for {
		r, n, err = buf.ReadRune()
		if err != nil {
			if q != 1 { // reached end without finding closing backquote
				return "", false
			}
			return src.SkipString(src.Len() - buf.Len()), true
		}
		switch q {
		case 0:
			if r != '`' {
				break
			}
			q = 1
		case 1:
			if r != '`' { // previous backquote was closing
				buf.UnreadRune()
				return src.SkipString(src.Len() - buf.Len()), true
			}
			q = 2
		case 2:
			if r != '`' { // two consecutive backquotes and then different character
				buf.UnreadRune()
				buf.UnreadRune()
				return src.SkipString(src.Len() - buf.Len()), true
			}
			q = 0
		}
	}
}

// SeeDoubleQuoteString …
func SeeDoubleQuoteString(src *Src) interface{} {
	t := src.Copy()
	quoted, ok := DelimitDoubleQuoteString(t)
	if !ok {
		return nil
	}
	var err error
	var str string
	if str, err = strconv.Unquote(quoted); err != nil {
		return nil
	}
	src.Become(t)
	return str
}

func DelimitDoubleQuoteString(src *Src) (string, bool) {
	var m int // number of bytes accepted into the quoted portion
	buf := src.Buffer()
	// first quote
	r, n, err := buf.ReadRune()
	if err != nil || r != '"' {
		return "", false
	}
	m += n
	//
	var backslash bool // true if previous character is backslash
	for {
		r, n, err = buf.ReadRune()
		if err != nil {
			return "", false // reached end of string without closing quote
		}
		if backslash {
			backslash = false
		} else {
			switch r {
			case '\\':
				backslash = true
			case '"':
				return src.SkipString(src.Len() - buf.Len()), true
			default:
			}
		}
	}
}
