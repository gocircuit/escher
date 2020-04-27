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
	"strings"

	"github.com/hoijui/escher/a"
	cir "github.com/hoijui/escher/circuit"
)

func SeeValueOrNil(src *a.Src) (x cir.Value) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	return SeeValue(src)
}

func SeeValue(src *a.Src) (x cir.Value) {
	if x = SeeCircuit(src); x != nil {
		return
	}
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
	if x = SeeVerb(src); x != nil {
		return
	}
	if x = SeeName(src); x != nil { // must be last since it will consume the empty string
		return
	}
	panic(0)
}

func SeeName(src *a.Src) cir.Name {
	return src.Consume(a.IsIdentifier)
}

// SeeVerb ...
func SeeVerb(src *a.Src) interface{} {
	t := src.Copy()
	verb := ""
	switch {
	case t.TryMatch("*"):
		verb = "*"
	case t.TryMatch("@"):
		verb = "@"
	default:
		return nil
	}
	delimit := t.Consume(a.IsIdentifierOrRefineSymbol)
	xx := strings.Split(delimit, a.RefineSymbolString)
	if len(xx) == 1 && xx[0] == "" {
		xx = nil
	}
	src.Become(t)
	var nn []cir.Name
	for _, x := range xx {
		nn = append(nn, x)
	}
	return cir.Circuit(cir.NewVerbAddress(verb, nn...))
}

// Int …
func SeeInt(src *a.Src) interface{} {
	t := src.Copy()
	l := a.Literal(t)
	if l == a.NullLiteral {
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
func SeeFloat(src *a.Src) interface{} {
	t := src.Copy()
	l := a.Literal(t)
	if l == a.NullLiteral {
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
func SeeComplex(src *a.Src) interface{} {
	t := src.Copy()
	l := a.Literal(t)
	if l == a.NullLiteral {
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
func SeeBackquoteString(src *a.Src) interface{} {
	t := src.Copy()
	quoted, ok := DelimitBackquoteString(t)
	if !ok {
		return nil
	}
	str := quoted[1 : len(quoted)-1]
	src.Become(t)
	return str
}

func DelimitBackquoteString(src *a.Src) (string, bool) {
	var m int // number of bytes accepted into the quoted portion
	buf := src.Buffer()
	// first backquote
	r, n, err := buf.ReadRune()
	if err != nil || r != '`' {
		return a.NullLiteral, false
	}
	m += n
	//
	var q int // number of backquotes right behind cursor, not counting opening backquote
	for {
		r, n, err = buf.ReadRune()
		if err != nil {
			if q != 1 { // reached end without finding closing backquote
				return a.NullLiteral, false
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
func SeeDoubleQuoteString(src *a.Src) interface{} {
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

func DelimitDoubleQuoteString(src *a.Src) (string, bool) {
	var m int // number of bytes accepted into the quoted portion
	buf := src.Buffer()
	// first quote
	r, n, err := buf.ReadRune()
	if err != nil || r != '"' {
		return a.NullLiteral, false
	}
	m += n
	//
	var backslash bool // true if previous character is backslash
	for {
		r, n, err = buf.ReadRune()
		if err != nil {
			return a.NullLiteral, false // reached end of string without closing quote
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
