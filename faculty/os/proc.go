// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/gocircuit/escher/be"
	cir "github.com/gocircuit/escher/circuit"
	kio "github.com/gocircuit/escher/kit/io"
)

// Process
type Process struct{ be.Sparkless }

func (Process) CognizeCommand(eye *be.Eye, dvalue interface{}) {
	x := cir.New()
	if exit := spawnProcess(eye, cognizeCommand(dvalue)); exit != nil {
		x.Grow("Exit", 1)
		eye.Show("Exit", x)
	} else {
		x.Grow("Exit", 0)
		eye.Show("Exit", x)
	}
}

func spawnProcess(eye *be.Eye, cmd *exec.Cmd) (err error) {
	var stdin io.WriteCloser
	var stdout io.ReadCloser
	var stderr io.ReadCloser
	stdin, err = cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err = cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatalf("Problem starting %s (%v)", cmd.Path, err)
		return err
	}
	// We cannot call cmd.Wait before all std streams have been closed.
	stdClose := make(chan struct{}, 3)
	stdin = kio.RunOnCloseWriter(stdin, func() { stdClose <- struct{}{} })
	stdout = kio.RunOnCloseReader(stdout, func() { stdClose <- struct{}{} })
	stderr = kio.RunOnCloseReader(stderr, func() { stdClose <- struct{}{} })
	g := cir.New().
		Grow("Stdin", stdin).
		Grow("Stdout", stdout).
		Grow("Stderr", stderr)
	// log.Printf("os process io (%v)", Linearize(fmt.Sprintf("%v", when)))
	eye.Show("IO", g)
	<-stdClose
	<-stdClose
	<-stdClose
	// log.Printf("os process waiting (%v)", Linearize(fmt.Sprintf("%v", when)))
	err = cmd.Wait()
	switch err.(type) {
	case nil, *exec.ExitError:
	default:
		panic(err)
	}
	// log.Printf("os process exit (%v)", Linearize(fmt.Sprintf("%v", when)))
	return err
}

func (Process) CognizeExit(*be.Eye, interface{}) {}

func (Process) CognizeIO(*be.Eye, interface{}) {}

//	{
//		Env { "PATH=/abc:/bin", "LESS=less" }
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
//
func cognizeCommand(v interface{}) *exec.Cmd {
	img, ok := v.(cir.Circuit)
	if !ok {
		panic(fmt.Sprintf("Non-image sent to Process.Command (%v)", v))
	}
	cmd := &exec.Cmd{}
	cmd.Path = img.StringAt("Path") // mandatory
	cmd.Args = []string{cmd.Path}
	if dir, ok := img.StringOptionAt("Dir"); ok {
		cmd.Dir = dir
	}
	if env, ok := img.CircuitOptionAt("Env"); ok {
		for _, key := range env.Numbers() {
			cmd.Env = append(cmd.Env, env.StringAt(key))
		}
	}
	if args, ok := img.CircuitOptionAt("Args"); ok {
		for _, key := range args.Numbers() {
			cmd.Args = append(cmd.Args, args.StringAt(key))
		}
	}
	// log.Printf("os process command (%v)", Linearize(img.Print("", "")))
	return cmd
}
