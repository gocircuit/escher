// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

import (
	"fmt"
	"testing"
)

// func TestReflex(t *testing.T) {
// 	r := SeeReflex(NewSrcString("reflex Name(A, B, C) // comment"))
// 	fmt.Printf("r=%v\n", r)
// }

// func TestInt(t *testing.T) {
// 	v, ok := Int(NewSrcString("-31***"))
// 	fmt.Printf("v=%v ok=%v\n", v, ok)
// }

// func TestFloat(t *testing.T) {
// 	v, ok := Float(NewSrcString("-07.1e23$$$"))
// 	fmt.Printf("v=%v ok=%v\n", v, ok)
// }

// func TestComplex(t *testing.T) {
// 	v, ok := Complex(NewSrcString("(-07.1e23+2i)$$$"))
// 	fmt.Printf("v=%v ok=%v\n", v, ok)
// }

// func TestBackquoteString(t *testing.T) {
// 	v, ok := BackquoteString(NewSrcString("`hello\\\\tthere`"))
// 	fmt.Printf("v=%v ok=%v\n", v, ok)
// }

// func TestDoubleQuoteString(t *testing.T) {
// 	v, ok := DoubleQuoteString(NewSrcString(`"hello\tthere"`))
// 	fmt.Printf("v=%v ok=%v\n", v, ok)
// }

// func TestScope(t *testing.T) {
// 	s, ok := Scope(NewSrcString("[1, 2, `a`]…"))
// 	fmt.Printf("s=%v, ok=%v\n", s, ok)
// }

// func TestField(t *testing.T) {
// 	n, s, ok := Field(NewSrcString("baha: [1, 2, `a`] ,…"))
// 	fmt.Printf("n=%v, s=%v, ok=%v\n", n, s, ok)
// }

// func TestRecord(t *testing.T) {
// 	r, ok := Record(NewSrcString("{ baha: [1, 2, `a`], empty: [], mama: [`3a`]}"))
// 	fmt.Printf("--\n%v\n--\nok=%v\n", (*record.Record)(r).String(), ok)
// }

// func TestMatch(t *testing.T) {
// 	m := Match(NewSrcString("a.b = c.d // ha\n"))
// 	fmt.Printf("match=%v\n", m)
// }

var testSource = []string{
	`
NaMo { // comment
	and And
	not Not

	str "stringißh"
	num +12.3e5
	msg {
		msg: ["http://gocircuit.org/hello.html"],
		num: [12.3e5],  // number
	} // string

	A = and.A // matching
	B = and.B
	not.B = C
	and.C = not.A
	X = src
	msg.Src = Y
	not.N = +3.14e00 // assign constants directly to wires, only on the right side

	// peer declarations are not sensitive to order within the block
	src ` + "`" + `
<html>
<head><title>E.g.</title></head>
<body>Hello world!</body>
</html>
` + "`" + `
	3.14 // ok
}
`,
}

func TestSyntax(t *testing.T) {
	for i, s := range testSource {
		src := NewSrcString(s)
		c := SeeCircuit(src)
		if c == nil {
			t.Fatalf("#%d misparses", i)
		}
		fmt.Printf("circuit=%v\n---\n", c)
	}
}
