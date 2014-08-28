// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocircuit/circuit/client"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
)

// Process
type Process struct{}

func (x Process) Materialize() be.Reflex {
	p := &process{
		id: ChooseID(),
		spawn: make(chan interface{}),
	}
	reflex, _ := plumb.NewEyeCognizer(p.cognize, "Server", "Command", "Spawn", "Exit", "IO")
	return reflex
}

type process struct{
	id string  // ID of this process reflex instance
	spawn chan interface{}
	sync.Mutex
	server *string // root-level anchor of the server where the process is to be started
	cmd *client.Cmd
}

func (p *process) cognize(eye *plumb.Eye, valve string, value interface{}) {
	switch valve {
	case "Server":
		if p.cognizeServer(value) {
			p.startBack(eye)
		}
	case "Command":
		if p.cognizeCommand(value) {
			p.startBack(eye)
		}
	case "Spawn":
		p.spawn <- value
		log.Printf("circuit process spawning (%v)", Linearize(fmt.Sprintf("%v", value)))
	}
}

func (p *process) cognizeServer(v interface{}) (ready bool) {
	a, ok := v.(string)
	if !ok {
		panic("process server anchor is non-string")
	}
	p.Lock()
	defer p.Unlock()
	if p.server != nil {
		panic("process server anchor already set")
	}
	p.server = &a
	return p.cmd != nil
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
		return
	}
	p.Lock()
	defer p.Unlock()
	if p.cmd != nil {
		panic("process command already set")
	}
	p.cmd = &client.Cmd{}
	p.cmd.Path = img.String("Path") // mandatory
	if img.Has("Dir") {
		p.cmd.Dir = img.String("Dir")
	}
	env := img.Walk("Env")
	for _, key := range env.Numbers() {
		p.cmd.Env = append(p.cmd.Env, env.String(key))
	}
	args := img.Walk("Args")
	for _, key := range args.Numbers() {
		p.cmd.Args = append(p.cmd.Args, args.String(key))
	}
	log.Printf("circuit process command (%v)", Linearize(img.Print("", "t")))
	return p.server != nil
}

func (p *process) startBack(eye *plumb.Eye) {
	p.Lock()
	defer p.Unlock()
	f := &processFixed{
		id: p.id,
		eye: eye,
		server: p.server,
		cmd: p.cmd,
	}
	go f.backLoop(p.spawn)
}

type processFixed struct {
	id string
	eye *plumb.Eye
	server *string
	cmd *client.Cmd
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
		log.Printf("circuit process exit sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *processFixed) spawnProcess(spwn interface{}) error {
	anchor := program.Client.Walk([]string{*p.server, "escher", program.Name, "circuit.Process", p.id})
	proc, err := anchor.MakeProc(*p.cmd)
	if err != nil {
		panic("invalid command argument")
	}
	defer anchor.Scrub()
	g := Image{
		"Spawn":  spwn,
		"Stdin":  proc.Stdin(),
		"Stdout": proc.Stdout(),
		"Stderr": proc.Stderr(),
	}
	log.Printf("circuit process io (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	p.eye.Show("IO", g)
	log.Printf("circuit process waiting (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	stat, err := proc.Wait()
	if err != nil {
		panic("process wait aborted by user")
	}
	log.Printf("circuit process exit (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	return stat.Exit
}
