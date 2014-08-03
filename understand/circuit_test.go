// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	//"fmt"
	"testing"

	"github.com/gocircuit/escher/see"
)

var src = []string{
	`
// Nand is a circuit.
Nand { // comments are everywhere
	and And
	not Not

	str "stringißh"
	num +12.3e5
	msg {
		msg "http://gocircuit.org/hello.html"
		num 12.3e5 // number
	} // string

	A = and.A // matching
	Y = and.Y
	not.X = X
	and.W = not.U
	and.D = not.V
	msg.Src = msg.A
	not.N = +3.14e00 // assign constants directly to wires, only on the right side

	// peer declarations are not sensitive to order within the block
	src ` + "`" +`
<html>
<head><title>E.g.</title></head>
<body>Hello world!</body>
</html>
` + "`" +`
	=5.14 // return, 
}
// end comment
`,
	`
namarupa{
	nama Name
	rupa 123
}`,
	`
circuit {
	="123"
}
`,
}

func TestUnderstand(t *testing.T) {
	//fmt.Printf(src)
	for _, r := range src {
		s := see.SeeCircuit(see.NewSrcString(r))
		t := Understand(s)
		printf2(t.Print("", "\t"))
	}
}

func TestFaculty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			printf("recovered: <%v>", r)
		}
	}()
	ns := NewFaculty()
	ns.UnderstandDirectory("/Users/petar/0/src/github.com/gocircuit/escher/understand/testdata")
}
