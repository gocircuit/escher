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
	"github.com/gocircuit/escher/think"
)

// Process
type Process struct{}

func (x Process) Materialize() think.Reflex {
	// Create reflex synapses
	cmdEndo, cmdExo := think.NewSynapse()
	spawnEndo, spawnExo := think.NewSynapse()
	exitEndo, exitExo := think.NewSynapse()
	ioEndo, ioExo := think.NewSynapse()
	serverEndo, serverExo := think.NewSynapse()
	//
	go func() {
		p := &process{
			id:    ChooseID(),
			ready: make(chan struct{}),
			spawn: make(chan interface{}),
		}
		p.reExit = exitEndo.Focus(think.DontCognize)
		p.reIO = ioEndo.Focus(think.DontCognize)
		serverEndo.Focus(p.CognizeServer)
		cmdEndo.Focus(p.CognizeCommand)
		spawnEndo.Focus(p.CognizeSpawn)
		p.loop()
	}()
	//
	return think.Reflex{
		"Server":  serverExo, // in-only
		"Command": cmdExo,    // in-only
		"Spawn":   spawnExo,  // in-only
		"Exit":    exitExo,   // out-only
		"IO":      ioExo,     // out-only
	}
}

// process is the materialized process reflex
type process struct {
	id     string // ID of this process reflex instance
	reExit *think.ReCognizer
	reIO   *think.ReCognizer
	arg    struct {
		sync.Mutex
		server string // root-level anchor of the server where the process is to be started
		cmd    *client.Cmd
	}
	ready chan struct{}    // notify loop that arguments are ready
	spawn chan interface{} // notify loop of spawn strobes
}

func (p *process) CognizeServer(v interface{}) {
	a, ok := v.(string)
	if !ok {
		panic("process server anchor is non-string")
	}
	p.arg.Lock()
	defer p.arg.Unlock()
	if p.arg.server != "" {
		panic("process server anchor already set")
	}
	p.arg.server = a
	if p.arg.cmd != nil {
		close(p.ready)
	}
}

//
//	{
//		Env {
//			"PATH=/abc:/bin"
//			"LESS=less"
//		}
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
//
func (p *process) CognizeCommand(v interface{}) {
	img, ok := v.(Image)
	if !ok {
		log.Printf("Non-image sent to Process.Command (%v)", v)
		return
	}
	p.arg.Lock()
	defer p.arg.Unlock()
	if p.arg.cmd != nil {
		panic("process command already set")
	}
	p.arg.cmd = &client.Cmd{}
	p.arg.cmd.Path = img.String("Path") // mandatory
	if img.Has("Dir") {
		p.arg.cmd.Dir = img.String("Dir")
	}
	env := img.Walk("Env")
	for _, key := range env.Sort() {
		p.arg.cmd.Env = append(p.arg.cmd.Env, env.String(key))
	}
	args := img.Walk("Args")
	for _, key := range args.Sort() {
		p.arg.cmd.Args = append(p.arg.cmd.Args, args.String(key))
	}
	log.Printf("circuit process command (%v)", Linearize(img.Print("", "t")))
	if p.arg.server != "" {
		close(p.ready)
	}
}

func (p *process) CognizeSpawn(v interface{}) {
	p.spawn <- v
	log.Printf("circuit process spawn (%v)", Linearize(fmt.Sprintf("%v", v)))
}

func (p *process) loop() {
	<-p.ready // make sure arguments (command and server) have been received
	for {
		spwn := <-p.spawn
		var x Image
		if exit := p.spawnProcess(spwn); exit != nil {
			x = Image{
				"Spawn": spwn,
				"Exit":  1,
			}
			p.reExit.ReCognize(x)
		} else {
			x = Image{
				"Spawn": spwn,
				"Exit":  0,
			}
			p.reExit.ReCognize(x)
		}
		log.Printf("circuit process exit sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *process) spawnProcess(spwn interface{}) error {
	p.arg.Lock()
	server := p.arg.server
	cmd := p.arg.cmd
	p.arg.Unlock()
	//
	anchor := program.Client.Walk([]string{server, "escher", program.Name, "circuit.Process", p.id})
	proc, err := anchor.MakeProc(*cmd)
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
	p.reIO.ReCognize(g)
	log.Printf("circuit process waiting (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	stat, err := proc.Wait()
	if err != nil {
		panic("process wait aborted by user")
	}
	log.Printf("circuit process exit (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	return stat.Exit
}
