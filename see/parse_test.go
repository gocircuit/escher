// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"fmt"
	"testing"
)

var testDesign = []string{
	`3.19e-2`,
	`22`,
	`"ha" `,
	"`la` ",
	`1-2i`,
	`name`,
	`@root`,
}

func TestDesign(t *testing.T) {
	for _, q := range testDesign {
		x := SeeDesignOrName(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
		}
		fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testSource = []string{
	`
NaMo { // comment
	and And
	not Not

	str "stringißh"
	num +12.3e5
	msg {
		msg "http://gocircuit.org/hello.html",
		num 12.3e5  // number
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

// func TestSyntax(t *testing.T) {
// 	for i, s := range testSource {
// 		src := NewSrcString(s)
// 		c := SeeCircuit(src)
// 		if c == nil {
// 			t.Fatalf("#%d misparses", i)
// 		}
// 		fmt.Printf("circuit=%v\n---\n", c)
// 	}
// }
