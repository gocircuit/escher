// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	// "fmt"
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
			id: ChooseID(),
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
		"Server": serverExo, // in-only
		"Command": cmdExo, // in-only
		"Spawn": spawnExo, // in-only
		"Exit": exitExo, // out-only
		"IO": ioExo, // out-only
	}
}

// process is the materialized process reflex
type process struct {
	id string  // ID of this process reflex instance
	reExit *think.ReCognizer
	reIO *think.ReCognizer
	arg struct {
		sync.Mutex
		server string // root-level anchor of the server where the process is to be started
		cmd *client.Cmd
	}
	ready chan struct{} // notify loop that arguments are ready
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
	env := img.Walk("Env")
	for _, key := range env.Sort() {
		p.arg.cmd.Env = append(p.arg.cmd.Env, env.String(key))
	}
	args := img.Walk("Args")
	for _, key := range args.Sort() {
		p.arg.cmd.Args = append(p.arg.cmd.Args, args.String(key))
	}
	if p.arg.server != "" {
		close(p.ready)
	}
}

func (p *process) CognizeSpawn(v interface{}) {
	p.spawn <- v
}

func (p *process) loop() {
	<-p.ready // make sure arguments (command and server) have been received
	for {
		spwn := <-p.spawn
		if exit := p.spawnProcess(); exit != nil {
			p.reExit.ReCognize(
				Image{
					"Spawn": spwn,
					"Exit": 1,
				},
			)
		} else {
			p.reExit.ReCognize(
				Image{
					"Spawn": spwn,
					"Exit": 0,
				},
			)
		}
	}
}

func (p *process) spawnProcess() error {
	p.arg.Lock()
	server := p.arg.server
	cmd := p.arg.cmd
	p.arg.Unlock()
	//
	anchor := program.Client.Walk([]string{server, "escher", program.Name, p.id})
	proc, err := anchor.MakeProc(*cmd)
	if err != nil {
		panic("invalid command argument")
	}
	defer anchor.Scrub()
	p.reIO.ReCognize(
		Image{
			"Stdin": proc.Stdin(), 
			"Stdout": proc.Stdout(),
			"Stderr": proc.Stderr(),
		},
	)
	stat, err := proc.Wait()
	if err != nil {
		panic("process wait aborted by user")
	}
	return stat.Exit
}
