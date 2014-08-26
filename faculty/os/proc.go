// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"fmt"
	"log"
	"os/exec"
	"sync"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

// Process
type Process struct{}

func (x Process) Materialize() be.Reflex {
	p := &process{
		spawn: make(chan interface{}),
	}
	reflex, _ := plumb.NewEyeCognizer(p.cognize, "Command", "Spawn", "Exit", "IO")
	return reflex
}

type process struct{
	spawn chan interface{}
	sync.Mutex
	cmd *exec.Cmd
}

func (p *process) cognize(eye *plumb.Eye, valve string, value interface{}) {
	switch valve {
	case "Command":
		if p.cognizeCommand(value) {
			p.startBack(eye)
		}
	case "Spawn":
		p.spawn <- value
		log.Printf("os process spawning (%v)", Linearize(fmt.Sprintf("%v", value)))
	}
}

//	{
//		Env { "PATH=/abc:/bin", "LESS=less" }
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
func (p *process) cognizeCommand(v interface{}) (ready bool) {
	img, ok := v.(Image)
	if !ok {
		log.Printf("Non-image sent to Process.Command (%v)", v)
		return false
	}
	p.Lock()
	defer p.Unlock()
	if p.cmd != nil {
		panic("process command already set")
	}
	p.cmd = &exec.Cmd{}
	p.cmd.Path = img.String("Path") // mandatory
	if img.Has("Dir") {
		p.cmd.Dir = img.String("Dir")
	}
	env := img.Walk("Env")
	for _, key := range env.Sort() {
		p.cmd.Env = append(p.cmd.Env, env.String(key))
	}
	args := img.Walk("Args")
	for _, key := range args.Sort() {
		p.cmd.Args = append(p.cmd.Args, args.String(key))
	}
	log.Printf("os process command (%v)", Linearize(img.Print("", "t")))
	return true
}

func (p *process) startBack(eye *plumb.Eye) {
	p.Lock()
	defer p.Unlock()
	f := &processFixed{
		eye: eye,
		cmd: p.cmd,
	}
	go f.backLoop(p.spawn)
}

type processFixed struct {
	eye *plumb.Eye
	cmd *exec.Cmd
}

func (p *processFixed) backLoop(spawn <-chan interface{}) {
	for {
		spwn := <-spawn
		var x Image
		if exit := p.spawnProcess(spwn); exit != nil {
			x = Image{
				"Spawn": spwn,
				"Exit":  1,
			}
			p.eye.Show("Exit", x)
		} else {
			x = Image{
				"Spawn": spwn,
				"Exit":  0,
			}
			p.eye.Show("Exit", x)
		}
		log.Printf("os process exit sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *processFixed) spawnProcess(spwn interface{}) (err error) {
	stdin, err :=  p.cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err := p.cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := p.cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	if err = p.cmd.Start(); err != nil {
		return err
	}
	g := Image{
		"Spawn":  spwn,
		"Stdin":  stdin,
		"Stdout": stdout,
		"Stderr": stderr,
	}
	log.Printf("os process io (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	p.eye.Show("IO", g)
	log.Printf("os process waiting (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	err = p.cmd.Wait()
	switch err.(type) {
	case nil, *exec.ExitError:
	default:
		panic(err)
	}
	log.Printf("os process exit (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	return err
}
