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
	"sync"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	kio "github.com/gocircuit/escher/kit/io"
)

// Process
type Process struct {
	spawn     chan interface{}
	sync.Once // start backloop once
}

func (p *Process) Spark(eye *be.Eye, _ Circuit, _ ...interface{}) Value {
	p.spawn = make(chan interface{})
	return nil
}

func (p *Process) CognizeCommand(eye *be.Eye, dvalue interface{}) {
	p.Once.Do(
		func() {
			back := &processBack{
				eye:   eye,
				cmd:   cognizeCommand(dvalue),
				spawn: p.spawn,
			}
			go back.loop()
		},
	)
}

func (p *Process) CognizeWhen(eye *be.Eye, dvalue interface{}) {
	p.spawn <- dvalue
}

func (p *Process) CognizeExit(*be.Eye, interface{}) {}

func (p *Process) CognizeIO(*be.Eye, interface{}) {}

//	{
//		Env { "PATH=/abc:/bin", "LESS=less" }
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
//
func cognizeCommand(v interface{}) *exec.Cmd {
	img, ok := v.(Circuit)
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

type processBack struct {
	eye   *be.Eye
	cmd   *exec.Cmd
	spawn <-chan interface{}
}

func (p *processBack) loop() {
	for {
		when := <-p.spawn
		x := New().Grow("When", when)
		if exit := p.spawnProcess(when); exit != nil {
			x.Grow("Exit", 1)
			p.eye.Show("Exit", x)
		} else {
			x.Grow("Exit", 0)
			p.eye.Show("Exit", x)
		}
		// log.Printf("os process exit sent (%v)", Linearize(x.Print("", "")))
	}
}

func (p *processBack) spawnProcess(when interface{}) (err error) {
	var stdin io.WriteCloser
	var stdout io.ReadCloser
	var stderr io.ReadCloser
	stdin, err = p.cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err = p.cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err = p.cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err = p.cmd.Start(); err != nil {
		log.Fatalf("Problem starting %s (%v)", p.cmd.Path, err)
		return err
	}
	// We cannot call cmd.Wait before all std streams have been closed.
	stdClose := make(chan struct{}, 3)
	stdin = kio.RunOnCloseWriter(stdin, func() { stdClose <- struct{}{} })
	stdout = kio.RunOnCloseReader(stdout, func() { stdClose <- struct{}{} })
	stderr = kio.RunOnCloseReader(stderr, func() { stdClose <- struct{}{} })
	g := New().
		Grow("When", when).
		Grow("Stdin", stdin).
		Grow("Stdout", stdout).
		Grow("Stderr", stderr)
	// log.Printf("os process io (%v)", Linearize(fmt.Sprintf("%v", when)))
	p.eye.Show("IO", g)
	<-stdClose
	<-stdClose
	<-stdClose
	// log.Printf("os process waiting (%v)", Linearize(fmt.Sprintf("%v", when)))
	err = p.cmd.Wait()
	switch err.(type) {
	case nil, *exec.ExitError:
	default:
		panic(err)
	}
	// log.Printf("os process exit (%v)", Linearize(fmt.Sprintf("%v", when)))
	return err
}
