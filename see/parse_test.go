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
		x := SeeArithmeticOrNameOrStar(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%v\n", x.Print("", "\t"))
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
		_, _, x := SeeMatching(NewSrcString(q), "$")
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		// fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testPeer = []string{
	`a\v b`,
	`\a @b`,
	`\ "abc"`,
	`a\ 3.13`,
	`a\ {}`,
	`\ { }`,
}

func TestPeer(t *testing.T) {
	for _, q := range testPeer {
		_, _, x := SeePeer(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		// fmt.Printf("%s %v\n", nm, x.Print("", "\t"))
	}
}

var testStar = []string{
	`{}`,
	`{
		g\ {}
		a = b
		\ {}
	}`,
	`{
		a\ b
		c\ @d
		e\ 1.23
		f\ "123"
		 = 0-2i
		 \ 123
	}`,
}

func TestStar(t *testing.T) {
	for _, q := range testStar {
		x := SeeStar(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		//fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testCircuit = []string{
	`nand\ {
		a\ and
		n\ not
		X=a.X
		Y=a.Y
		n.X=a.XandY
		b\ "3e3"
		n.notX=
		{}=
	}
	`,
}

func TestCircuit(t *testing.T) {
	for _, q := range testCircuit {
		f, r, x := SeeCircuit(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%s\\%s %v\n", f, r, x.Print("", "\t"))
	}
}
