// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocircuit/circuit/client"

	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/understand"

	// Load faculties
	"github.com/gocircuit/escher/faculty/acid"
	"github.com/gocircuit/escher/faculty/basic"
	"github.com/gocircuit/escher/faculty/circuit"
	_ "github.com/gocircuit/escher/faculty/escher"
	_ "github.com/gocircuit/escher/faculty/i"
	_ "github.com/gocircuit/escher/faculty/io"
	_ "github.com/gocircuit/escher/faculty/io/util"
	facultyos "github.com/gocircuit/escher/faculty/os"
	_ "github.com/gocircuit/escher/faculty/path"
	_ "github.com/gocircuit/escher/faculty/text"
	_ "github.com/gocircuit/escher/faculty/shelah"
	_ "github.com/gocircuit/escher/faculty/think"
	_ "github.com/gocircuit/escher/faculty/time"
	_ "github.com/gocircuit/escher/faculty/web/twitter"
)

var (
	flagShow     = flag.String("show", "", "compile and display object at given path; don't run")
	flagX        = flag.String("x", "", "program source directory X")
	flagY        = flag.String("y", "", "program source directory Y")
	flagName     = flag.String("n", "", "execution name")
	flagArg      = flag.String("a", "", "program arguments")
	flagDiscover = flag.String("d", "", "multicast UDP discovery address for circuit faculty, if needed")
)

func main() {
	flag.Parse()
	if *flagX == "" && *flagY == "" {
		fatalf("at least one source directory, X or Y, must be specified with -x or -y, respectively")
	}
	// Initialize faculties
	basic.Init(*flagName)
	facultyos.Init(*flagArg)
	loadCircuitFaculty(*flagName, *flagDiscover, *flagX, *flagY)
	//
	switch {
	case *flagShow != "":
		walk := strings.Split(*flagShow, ".")
		if len(walk) == 2 && walk[0] == "" && walk[1] == "" { // -show .
			walk = nil
		}
		_, cd := compile(*flagX, *flagY).Walk(walk...)
		switch t := cd.(type) {
		case *understand.Circuit:
			fmt.Println(t.Print("", "\t"))
		case understand.Faculty:
			fmt.Println(t.Print("", "\t"))
		default:
			fmt.Printf("%T/%v\n", t, t)
		}
	default:
		be.Space(compile(*flagX, *flagY)).Materialize("main")
		select {} // wait forever
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

func loadCircuitFaculty(name, discover, x, y string) {
	acid.Init(x, y)
	if discover == "" {
		circuit.Init(name, nil)
		return
	}
	if name == "" {
		panic("circuit-based Escher programs must have a non-empty name")
	}
	circuit.Init(name, client.DialDiscover(discover, nil))
}
