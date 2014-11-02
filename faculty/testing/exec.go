// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package testing

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Exec receives values from FilterAll and executes the included test circuits
// in separate OS processes.
type Exec struct {}

func (Exec) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	return nil
}

func (Exec) CognizeIn(eye *be.Eye, v interface{}) {
	x := v.(Circuit)
	//
	addr := x.AddressAt("Address")
	cmd := exec.Command(os.Args[0], "-src", srcDir, addr.String())

	var success bool
	if err := cmd.Run(); err != nil {
		fmt.Printf("- Test %v (%v)\n", addr, err)
		success = false
	} else {
		fmt.Printf("+ Test %v (ok)\n", addr)
		success = true
	}
	r := New().
		Grow("Address", addr).
		Grow("Result", success)
	eye.Show("Out", r)
}

func (Exec) CognizeOut(eye *be.Eye, v interface{}) {}
