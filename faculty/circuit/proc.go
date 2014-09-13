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
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/plumb"
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
	sync.Once // start backloop once
	spawn chan interface{} // notify loop of spawn memes
}

func (p *process) cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	switch dvalve {
	case "Command":
		p.Once.Do(
			func() {
				back := &processBack{
					eye: eye, 
					cmd: cognizeProcessCommand(dvalue), 
					spawn: p.spawn,
				}
				go back.loop()
			},
		)
	case "Spawn":
		p.spawn <- dvalue
		log.Printf("circuit process spawning (%v)", Linearize(fmt.Sprintf("%v", dvalue)))
	}
}

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
func cognizeProcessCommand(v interface{}) *client.Cmd {
	img, ok := v.(Circuit)
	if !ok {
		panic(fmt.Sprintf("Non-image sent to Process.Command (%v)", v))
	}
	cmd := &client.Cmd{}
	cmd.Path = img.StringAt("Path") // mandatory
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
	log.Printf("circuit process command (%v)", Linearize(img.Print("", "t")))
	return cmd
}

type processBack struct {
	eye *plumb.Eye
	cmd *client.Cmd
	spawn <-chan interface{}
}

func (p *processBack) loop() {
	for {
		spwn := <-p.spawn
		x := New().Grow("Spawn", spwn)
		if exit := p.spawnProcess(spwn); exit != nil {
			x.Grow("Exit", 1)
			p.eye.Show("Exit", x)
		} else {
			x.Grow("Exit", 0)
			p.eye.Show("Exit", x)
		}
		log.Printf("circuit process exit meme sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *processBack) spawnProcess(spwn interface{}) error {
	// anchor determination
	s := spwn.(Circuit)
	anchor := program.Client.Walk(
		[]string{
			s.StringAt("Server"), // server name
			s.StringAt("Name"), // (dynamic) execution name
		})
	//
	proc, err := anchor.MakeProc(*p.cmd)
	if err != nil {
		panic("invalid command argument")
	}
	defer anchor.Scrub()
	g := New().
		Grow("Spawn",  spwn,).
		Grow("Stdin",  proc.Stdin()).
		Grow("Stdout", proc.Stdout()).
		Grow("Stderr", proc.Stderr())
	log.Printf("circuit process io (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	p.eye.Show("IO", g)
	log.Printf("circuit process waiting (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	stat, err := proc.Wait()
	if err != nil {
		panic("process wait aborted by user")
	}
	log.Printf("circuit process (%v) exited", Linearize(fmt.Sprintf("%v", spwn)))
	if stat.Exit != nil {
		log.Printf("circuit process exit error: %v", stat.Exit)
	}
	return stat.Exit
}
