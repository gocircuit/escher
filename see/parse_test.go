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
		x := SeeArithmeticOrNameOrUnion(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%v\n", x)
	}
}

var testMatching = []string {
	`a.X = b.Y`,
	` X = y.Z `,
	` X = "hello"`,
	`123 =`,
	`=`,
	`X=`,
	`X.y=`,
}

func TestMatching(t *testing.T) {
	for _, q := range testMatching {
		x := SeeMatching(NewSrcString(q))
		if x == nil {
			t.Fatalf("problem parsing: %s", q)
			continue
		}
		// fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testPeer = []string{
	`a b`,
	`a @b`,
	`_ "abc"`,
	`a 3.13`,
	`a { }`,
}

func TestPeer(t *testing.T) {
	for _, q := range testPeer {
		x := SeePeer(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		// fmt.Printf("%s %v\n", nm, x.Print("", "\t"))
	}
}

var testUnion = []string{
	`{}`,
	`{
		g {}
		a = b
		x {}
	}`,
	`{
		a b
		c @d
		e 1.23
		f "123"
		 = 0-2i
		 _ 123
	}`,
}

func TestUnion(t *testing.T) {
	for _, q := range testUnion {
		x := SeeUnion(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		//fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testCircuit = []string{
	`nand {
		a and
		n not
		X=a.X
		Y=a.Y
		n.X=a.XandY
		b "3e3"
		n.notX=
		{}=
	}
	`,
}

func TestCircuit(t *testing.T) {
	for _, q := range testCircuit {
		x := SeeCircuit(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%s\n", x.Print("", "    "))
	}
}
