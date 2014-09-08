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

var testMeaning = []string{
	`3.19e-2`,
	`22`,
	`"ha" `,
	"`la` ",
	`1-2i`,
	`name`,
	`name.family`,
}

func TestMeaning(t *testing.T) {
	for _, q := range testMeaning {
		x := SeeMeaning(NewSrcString(q))
		if x == nil {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		fmt.Printf("%v\n", x)
	}
}

var testMatching = []string{
	`a:X = b:Y`,
	` X:a = y:Z `,
	` X:_ = "hello"`,
	`123 = "a"`,
	`X:y = a:_`,
}

func TestMatching(t *testing.T) {
	for _, q := range testMatching {
		x, _ := SeeMatching(NewSrcString(q), 0)
		if x == nil {
			t.Fatalf("problem parsing: %s", q)
			continue
		}
		// fmt.Printf("%v\n", x.Print("", "\t"))
	}
}

var testPeer = []string{
	`a b`,
	`_ "abc"`,
	`a 3.13,`,
	`a { },`,
	`a;`,
	`"abc"`,
	`"ha",`,
	`a { "cd" }; `,
	// `// haha
	// main { // jaja
	// 	proc circuit.Process
	// 	srv os.Arg
	// 	w Way3
	// 	d time.Delay
	// 	forkIO circuit.ForkIO
	// 	forkExit circuit.ForkExit

	// 	srv:Name = "Server"
	// 	proc:Server = srv:Value
	// 	proc:Command = {
	// 		Path "/usr/bin/say"
	// 		Args { "escher" }
	// 	}

	// 	proc:IO = forkIO:Forked

	// 	clunkIn io.Clunk
	// 	clunkOut io.Clunk
	// 	clunkErr io.Clunk
	// 	forkIO:Stdin = clunkIn:IO
	// 	forkIO:Stdout = clunkOut:IO
	// 	forkIO:Stderr = clunkErr:IO

	// 	spawnIgn Ignore
	// 	forkIO:Spawn = spawnIgn:Subject

	// 	proc:Spawn = w:A1
	// 	w:A0 = 1
	// 	w:A2 = d:X
	// 	d:Duration = 1e9 // 1 second
	// 	d:Y = forkExit:Spawn

	// 	exitIgn Ignore
	// 	proc:Exit = forkExit:Forked
	// 	forkExit:Exit = exitIgn:Subject
	// }
	// `, 
	`
	header {
		merge text.Merge
		merge:First = ""
		merge:Second = Title:_
		merge:Third = ""
		header:_ = merge:_
	}
	`,
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
	`{
		a:y = b:z
	}`,
	`{
		g {}
		a:y = b:z
		x {}
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
}

func TestCircuit(t *testing.T) {
	for _, q := range testCircuit {
		x := SeeCircuit(NewSrcString(q))
		if x.IsNil() {
			t.Errorf("problem parsing: %s", q)
			continue
		}
		//fmt.Printf("%v\n", x.Print("", "\t"))
	}
}
