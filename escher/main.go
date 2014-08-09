// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/understand"

	"github.com/gocircuit/escher/faculty"
	_ "github.com/gocircuit/escher/faculty/basic"
	// _ "github.com/gocircuit/escher/faculty/circuit"
)

var (
	flagLex  = flag.Bool("lex", false, "parse and show faculties without running")
	flagSrc  = flag.String("src", "", "program source directory")
)

func main() {
	flag.Parse()
	if *flagSrc == "" {
		fatalf("source directory must be specified with -src")
	}
	if *flagLex {
		fmt.Println(load(*flagSrc).Print("", "   "))
	} else {
		think.Space(load(*flagSrc)).Materialize("main")
		select{} // wait forever
	}
}

func load(src string) understand.Faculty {
	faculty.Root.UnderstandDirectory(src)
	return faculty.Root
}
