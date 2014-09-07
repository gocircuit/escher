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

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	kitio "github.com/gocircuit/escher/kit/io"
	"github.com/gocircuit/escher/kit/plumb"
)

// Process
type Process struct{}

func (x Process) Materialize() be.Reflex {
	p := &process{
		spawn: make(chan interface{}),
	}
	reflex, _ := plumb.NewEyeCognizer(p.cognize, "Command", "When", "Exit", "IO")
	return reflex
}

type process struct{
	spawn chan interface{}
	sync.Once // start backloop once
}

func (p *process) cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	switch dvalve {
	case "Command":
		p.Once.Do(
			func() {
				back := &processBack{
					eye: eye, 
					cmd: cognizeCommand(dvalue), 
					spawn: p.spawn,
				}
				go back.loop()
			},
		)
	case "When":
		p.spawn <- dvalue
		// log.Printf("OS process spawning (%v)", Linearize(fmt.Sprintf("%v", value)))
	}
}

//	{
//		Env { "PATH=/abc:/bin", "LESS=less" }
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
//
func cognizeCommand(v interface{}) *exec.Cmd {
	img, ok := v.(Image)
	if !ok {
		panic(fmt.Sprintf("Non-image sent to Process.Command (%v)", v))
	}
	cmd := &exec.Cmd{}
	cmd.Path = img.String(see.Name("Path")) // mandatory
	cmd.Args = []string{cmd.Path}
	if img.Has(see.Name("Dir")) {
		cmd.Dir = img.String(see.Name("Dir"))
	}
	env := img.Walk(see.Name("Env"))
	for _, key := range see.Numbers(env) {
		cmd.Env = append(cmd.Env, env.String(key))
	}
	args := img.Walk(see.Name("Args"))
	for _, key := range see.Numbers(args) {
		cmd.Args = append(cmd.Args, args.String(key))
	}
	// log.Printf("os process command (%v)", Linearize(img.Print("", "")))
	return cmd
}

type processBack struct {
	eye *plumb.Eye
	cmd *exec.Cmd
	spawn <-chan interface{}
}

func (p *processBack) loop() {
	for {
		when := <-p.spawn
		var x Image
		if exit := p.spawnProcess(when); exit != nil {
			x = Image{
				see.Name("When"): when,
				see.Name("Exit"):  1,
			}
			p.eye.Show("Exit", x)
		} else {
			x = Image{
				see.Name("When"): when,
				see.Name("Exit"):  0,
			}
			p.eye.Show("Exit", x)
		}
		// log.Printf("os process exit sent (%v)", Linearize(x.Print("", "")))
	}
}

func (p *processBack) spawnProcess(when interface{}) (err error) {
	var stdin io.WriteCloser
	var stdout io.ReadCloser
	var stderr io.ReadCloser
	stdin, err =  p.cmd.StdinPipe()
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
	stdin = kitio.RunOnCloseWriter(stdin, func() { stdClose <- struct{}{} })
	stdout = kitio.RunOnCloseReader(stdout, func() { stdClose <- struct{}{} })
	stderr = kitio.RunOnCloseReader(stderr, func() { stdClose <- struct{}{} })
	g := Image{
		see.Name("When"):  when,
		see.Name("Stdin"):  stdin,
		see.Name("Stdout"): stdout,
		see.Name("Stderr"): stderr,
	}
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
