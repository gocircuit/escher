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

var testValue = []string{
	`3.19e-2`,
	`22`,
	`"ha" `,
	"`la` ",
	`1-2i`,
	`name`,
	`name.family`,
	`.`,
}

func TestValue(t *testing.T) {
	for _, q := range testValue {
		x := SeeValue(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%v\n", x)
	}
}

var testReal = []string{
	`a:X = b:Y`,
	` X:a = y:Z `,
	` X:_ = "hello"`,
	`123 = "a"`,
	`X:y = a:_`,
	`X:y = :_`,
}

func TestReal(t *testing.T) {
	for _, q := range testReal {
		x, _ := SeeLink(NewSrcString(q), 0)
		if x == nil {
			t.Fatalf("problem parsing: %s", q)
			continue
		}
		// fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testPeer = []string{
	// `a b`,
	// `_ "abc"`,
	// `a 3.13,`,
	// `a { },`,
	// `a;`,
	// `"abc"`,
	// `"ha",`,
	// `a { "cd" }; `,
	`main {
		s Show
		s:_ = "¡Hello, world!"
		hello {
			a A
			hello:_ = a:_
		}
	}
	`,
	// `
	// header {
	// 	merge text.Merge
	// 	merge:First = ""
	// 	merge:Second = Title:_
	// 	merge:Third = ""
	// 	: = merge:
	// }
	// `,
}

func TestPeer(t *testing.T) {
	for _, q := range testPeer {
		x, m := SeePeer(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%v %v\n", x, m)
	}
}

var testCircuit = []string{
	`{}`,
	`{ a {}}`,
	`{
		a:y = b:z
	}`,
	`{
		g {}
		a:y = b:z
	}`,
	`{
		a b
		c d
		e 1.23
		f "123"
		a:_ = 0-2i
		_ 123
	}`,
	`{
		g {},
		a:q = b:p,
		x {};
		y {a, b, c, "def"; }
	}`,
	`{
		a and
		n not
		nand:X=a:X
		nand:Y=a:Y
		n:X=a:XandY
		b "3e3"
		n:notX=nand:_
		X:_=nand:a1,
		w "abcd",
	}
	`,
	`{
		t time.Ticker
		s sum
		out show
		t:Duration = 1e9
		t:Tick = s:Sum
		s:X = 5e9
		s:Y = out:Object
		s:Z = { "a", "b", "c" }
	}
	`,
	`{
		a       "a"
		b       12
		c       1.2
		d       1-2i
		e       {}
		f       { 1; "2"; 3.1; a, "b" }
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
		// fmt.Printf("CIR=\n%v\n", x)
	}
}

var testChamber = []string{
	`BuildIndexPage {
	        wf io.util.WriteFile
	        wf:Content = IndexPage
	        wf:Name = "index.html"
	        : = wf:Ready
	}
	`,
}

func TestChamber(t *testing.T) {
	for _, q := range testChamber {
		x := SeeChamber(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		//fmt.Printf("%v\n", x.Print("", "\t"))
	}
}
