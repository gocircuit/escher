// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package test

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
)

// Exec receives values from FilterAll and executes the included test circuits
// in separate OS processes.
type Exec struct{ be.Sparkless }

func (Exec) CognizeIn(eye *be.Eye, v interface{}) {
	x := v.(cir.Circuit)
	//
	addr := cir.Verb(x.CircuitAt("Address").Copy())
	addr.Gate[""] = "*"
	cmd := exec.Command(os.Args[0], "-src", srcDir, addr.String())

	var success bool
	if err := cmd.Run(); err != nil {
		fmt.Printf("- Test %v (%v)\n", addr, err)
		success = false
	} else {
		fmt.Printf("+ Test %v (ok)\n", addr)
		success = true
	}
	r := cir.New().
		Grow("Verb", cir.Circuit(addr)).
		Grow("Result", success)
	eye.Show("Out", r)
}

func (Exec) CognizeOut(eye *be.Eye, v interface{}) {}
