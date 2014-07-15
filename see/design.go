// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	"bytes"
	"fmt"
	"strconv"
)

func SeeDesign(src *Src) (v Design, ok bool) {
	if v, ok = SeeBasic(src); ok {
		return
	}
	if v, ok = SeeTree(src); ok {
		return
	}
	if v, ok = SeeNameOrPackage(src); ok {
		return
	}
	return nil, false
}

func SeeNoNameDesign(src *Src) (v Design, ok bool) {
	if v, ok = SeeBasic(src); ok {
		return
	}
	if v, ok = SeeTree(src); ok {
		return
	}
	return nil, false
}

func SeeBasic(src *Src) (v Design, ok bool) {
	if v, ok = SeeInt(src); ok {
		return
	}
	if v, ok = SeeFloat(src); ok {
		return
	}
	if v, ok = SeeComplex(src); ok {
		return
	}
	if v, ok = SeeBackquoteString(src); ok {
		return
	}
	if v, ok = SeeDoubleQuoteString(src); ok {
		return
	}
	return nil, false
}

// Name …
func SeeNameOrPackage(src *Src) (np Design, ok bool) {
	l := Identifier(src)
	if l == "" {
		return nil, false
	}
	if l[0] != '@' {
		return NameDesign(l), true
	}
	return RootNameDesign(l[1:]), true
}

// Int …
func SeeInt(src *Src) (i IntDesign, ok bool) {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return 0, false
	}
	r := bytes.NewBufferString(l)
	if n, _ := fmt.Fscanf(r, "%d", &i); n != 1 || r.Len() != 0  {
		return 0, false
	}
	src.Become(t)
	return i, true
}

// Float …
func SeeFloat(src *Src) (f FloatDesign, ok bool) {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return 0, false
	}
	r := bytes.NewBufferString(l)
	if n, _ := fmt.Fscanf(r, "%g", &f); n != 1 || r.Len() != 0 {
		return 0, false
	}
	src.Become(t)
	return f, true
}

// Complex …
func SeeComplex(src *Src) (c ComplexDesign, ok bool) {
	t := src.Copy()
	l := Literal(t)
	if l == "" {
		return 0, false
	}
	r := bytes.NewBufferString(l)
	if n, _ := fmt.Fscanf(r, "%g", &c); n != 1 || r.Len() != 0 {
		return 0, false
	}
	src.Become(t)
	return c, true
}

// SeeBackquoteString …
func SeeBackquoteString(src *Src) (str StringDesign, ok bool) {
	t := src.Copy()
	var quoted string
	quoted, ok = DelimitBackquoteString(t)
	if !ok {
		return "", false
	}
	// var err error
	// if str, err = strconv.Unquote(quoted); err != nil {
	// 	return "", false
	// }
	str = StringDesign(quoted[1:len(quoted)-1])
	src.Become(t)
	return str, true
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
func SeeDoubleQuoteString(src *Src) (sd StringDesign, ok bool) {
	t := src.Copy()
	var quoted string
	quoted, ok = DelimitDoubleQuoteString(t)
	if !ok {
		return "", false
	}
	var err error
	var str string
	if str, err = strconv.Unquote(quoted); err != nil {
		return "", false
	}
	src.Become(t)
	return StringDesign(str), true
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
