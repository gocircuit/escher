// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"

	"github.com/gocircuit/circuit/client"

	"github.com/gocircuit/escher/think"
	"github.com/gocircuit/escher/understand"
	"github.com/gocircuit/escher/faculty"

	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/faculty/circuit"
	facultyos "github.com/gocircuit/escher/faculty/os"
	_ "github.com/gocircuit/escher/faculty/time"
)

var (
	flagLex  = flag.Bool("lex", false, "parse and show faculties without running")
	flagSrc  = flag.String("src", "", "program source directory")
	flagName = flag.String("name", "", "execution name")
	flagArg = flag.String("arg", "", "program arguments")
	flagDiscover = flag.String("discover", "", "multicast UDP discovery address for circuit faculty, if needed")
)

func main() {
	flag.Parse()
	basic.Init(*flagName)
	facultyos.Init(*flagArg)
	if *flagSrc == "" {
		fatalf("source directory must be specified with -src")
	}
	if *flagLex {
		fmt.Println(compile(*flagSrc).Print("", "   "))
	} else {
		loadCircuitFaculty(*flagName, *flagDiscover)
		think.Space(compile(*flagSrc)).Materialize("main")
		select{} // wait forever
	}
}

func compile(src string) understand.Faculty {
	faculty.Root.UnderstandDirectory(src)
	return faculty.Root
}

func loadCircuitFaculty(name, discover string) {
	if discover == "" {
		return
	}
	if name == "" {
		panic("circuit-based Escher programs must have a non-empty name")
	}
	circuit.Init(name, client.DialDiscover(discover, nil))
}
