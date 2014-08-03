// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"bytes"
	"fmt"
	"strconv"
	"github.com/gocircuit/escher/star"
)

// SeeArithmeticOrImage parses a built-in type (int, float, complex, string, name) into a star.
func SeeArithmeticOrImage(src *Src) (x *star.Star) {
	if x = SeeArithmetic(src); x != nil {
		return
	}
	if x = SeeImage(src); x != nil {
		return
	}
	return nil
}

func SeeArithmeticOrNameOrImage(src *Src) (x *star.Star) {
	if x = SeeArithmetic(src); x != nil {
		return
	}
	if x = SeeName(src); x != nil {
		return
	}
	if x = SeeImage(src); x != nil {
		return
	}
	return nil
}

func SeeArithmetic(src *Src) (x *star.Star) {
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
	return nil
}

// Name …
func SeeName(src *Src) (x *star.Star) {
	l := Identifier(src)
	if l == "" {
		return nil
	}
	if l[0] != '@' {
		return Wrap(Name(l))
	}
	return Wrap(RootName(l))
}

// Int …
func SeeInt(src *Src) *star.Star {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return nil
	}
	r := bytes.NewBufferString(l)
	var i Int
	if n, _ := fmt.Fscanf(r, "%d", &i); n != 1 || r.Len() != 0  {
		return nil
	}
	src.Become(t)
	return Wrap(i)
}

// Float …
func SeeFloat(src *Src) *star.Star {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return nil
	}
	r := bytes.NewBufferString(l)
	var f Float
	if n, _ := fmt.Fscanf(r, "%g", &f); n != 1 || r.Len() != 0 {
		return nil
	}
	src.Become(t)
	return Wrap(f)
}

// Complex …
func SeeComplex(src *Src) *star.Star {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return nil
	}
	r := bytes.NewBufferString(l)
	var c Complex
	if n, _ := fmt.Fscanf(r, "%g", &c); n != 1 || r.Len() != 0 {
		return nil
	}
	src.Become(t)
	return Wrap(c)
}

// SeeBackquoteString …
func SeeBackquoteString(src *Src) *star.Star {
	t := src.Copy()
	quoted, ok := DelimitBackquoteString(t)
	if !ok {
		return nil
	}
	str := String(quoted[1:len(quoted)-1])
	src.Become(t)
	return Wrap(str)
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
			return src.SkipString(src.Len()-buf.Len()), true
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
				return src.SkipString(src.Len()-buf.Len()), true
			}
			q = 2
		case 2:
			if r != '`' { // two consecutive backquotes and then different character
				buf.UnreadRune()
				buf.UnreadRune()
				return src.SkipString(src.Len()-buf.Len()), true
			}
			q = 0
		}
	}
}

// SeeDoubleQuoteString …
func SeeDoubleQuoteString(src *Src) *star.Star {
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
	return Wrap(String(str))
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
				return src.SkipString(src.Len()-buf.Len()), true
			default:
			}
		}
	}
}
