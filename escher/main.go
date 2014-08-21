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
	_ "github.com/gocircuit/escher/faculty/io"
	_ "github.com/gocircuit/escher/faculty/io/util"
	facultyos "github.com/gocircuit/escher/faculty/os"
	_ "github.com/gocircuit/escher/faculty/time"
	_ "github.com/gocircuit/escher/faculty/text"
)

var (
	flagUn  = flag.Bool("un", false, "understand and show source without materializing it")
	flagX  = flag.String("x", "", "program source directory X")
	flagY  = flag.String("y", "", "program source directory Y")
	flagName = flag.String("n", "", "execution name")
	flagArg = flag.String("a", "", "program arguments")
	flagDiscover = flag.String("d", "", "multicast UDP discovery address for circuit faculty, if needed")
)

func main() {
	flag.Parse()
	basic.Init(*flagName)
	facultyos.Init(*flagArg)
	loadCircuitFaculty(*flagName, *flagDiscover)
	if *flagX == "" && *flagY == "" {
		fatalf("at least one source directory, X or Y, must be specified with -x or -y, respectively")
	}
	if *flagUn {
		fmt.Println(compile(*flagX, *flagY).Print("", "   "))
	} else {
		think.Space(compile(*flagX, *flagY)).Materialize("main")
		select{} // wait forever
	}
}

func compile(x, y string) understand.Faculty {
	if x != "" {
		faculty.Root.UnderstandDirectory(x)
	}
	if y != "" {
		faculty.Root.UnderstandDirectory(y)
	}
	return faculty.Root
}

func loadCircuitFaculty(name, discover string) {
	if discover == "" {
		circuit.Init(name, nil)
		return
	}
	if name == "" {
		panic("circuit-based Escher programs must have a non-empty name")
	}
	circuit.Init(name, client.DialDiscover(discover, nil))
}
