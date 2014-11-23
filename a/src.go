// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package a

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Src represents a sub-string within an ambient backing string.
type Src struct {
	back []byte
}

func NewSrcString(s string) *Src {
	return &Src{[]byte(s)}
}

func (src *Src) Become(g *Src) {
	src.back = g.back
}

func (src *Src) Bytes() []byte {
	return src.back
}

func (src *Src) String() string {
	return string(src.Bytes())
}

func (src *Src) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(src.Bytes())
}

func (src *Src) Len() int {
	return len(src.back)
}

func (src *Src) IsEmpty() bool {
	return src.Len() == 0
}

func (src *Src) RuneAt(at int) (r rune) {
	r, _ = utf8.DecodeRune(src.back[at:])
	return
}

func (src *Src) Skip(n int) {
	src.back = src.back[n:]
}

func (src *Src) SkipString(n int) string {
	r := string(src.back[:n])
	src.back = src.back[n:]
	return r
}

func (src *Src) Copy() *Src {
	return &Src{
		back: src.back,
	}
}

func (src *Src) SliceFrom(at int) *Src {
	return &Src{
		back: src.back[at:],
	}
}

func (src *Src) TryMatch(prefix string) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	src.Match(prefix)
	return true
}

func (src *Src) Match(prefix string) {
	if strings.HasPrefix(src.String(), prefix) {
		src.Skip(len(prefix))
		return
	}
	panic("no match")
}

type ConsumeFunc func(rune) bool

func (src *Src) Consume(match ConsumeFunc) string {
	buf := src.Bytes()
	var m int
	for r, n := utf8.DecodeRune(buf); r != utf8.RuneError; r, n = utf8.DecodeRune(buf) {
		if !match(r) {
			break
		}
		buf = buf[n:]
		m += n
	}
	return src.SkipString(m)
}

func panicf(format string, arg ...interface{}) {
	panic(fmt.Sprintf(format, arg...))
}

func printf(format string, arg ...interface{}) {
	println(fmt.Sprintf(format, arg...))
}
