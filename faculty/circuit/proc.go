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

func (x Process) Materialize(matter *be.Matter) be.Reflex {
	p := &process{
		name: matter.LastName(),
		spawn: make(chan interface{}),
	}
	reflex, _ := plumb.NewEyeCognizer(p.cognize, "Command", "Spawn", "Exit", "IO")
	return reflex
}

type process struct{
	name string
	sync.Once // start backloop once
	spawn chan interface{} // notify loop of spawn memes
	sync.Mutex
	cmd *client.Cmd
}

func (p *process) cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	switch dvalve {
	case "Command":
		p.Once.Do(
			func() {
				back := &processBack{
					name: p.name,
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
//		Env { "PATH=/abc:/bin", "LESS=less" }
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
func cognizeProcessCommand(v interface{}) *client.Cmd {
	img, ok := v.(Image)
	if !ok {
		panic(fmt.Sprintf("Non-image sent to Process.Command (%v)", v))
	}
	cmd := &client.Cmd{}
	cmd.Path = img.String("Path") // mandatory
	if img.Has("Dir") {
		cmd.Dir = img.String("Dir")
	}
	env := img.Walk("Env")
	for _, key := range env.Numbers() {
		cmd.Env = append(cmd.Env, env.String(key))
	}
	args := img.Walk("Args")
	for _, key := range args.Numbers() {
		cmd.Args = append(cmd.Args, args.String(key))
	}
	log.Printf("circuit process command (%v)", Linearize(img.Print("", "t")))
	return cmd
}

type processBack struct {
	name string
	eye *plumb.Eye
	cmd *client.Cmd
	spawn <-chan interface{}
}

func (p *processBack) loop() {
	for {
		spwn := <-p.spawn
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
		log.Printf("circuit process exit meme sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *processBack) spawnProcess(spwn interface{}) error {
	// anchor determination
	s := spwn.(Image)
	if s.String("Name") == "" {
		panic("circuit process execution name required")
	}
	if s.String("Server") == "" {
		panic("circuit process server required")
	}
	anchor := program.Client.Walk(
		[]string{
			s.String("Server"), // server name
			p.name, // reflex' unique materialization name
			s.String("Name"), // (dynamic) execution name
		})
	//
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
	log.Printf("circuit process (%v) exited", Linearize(fmt.Sprintf("%v", spwn)))
	if stat.Exit != nil {
		log.Printf("circuit process exit error: %v", stat.Exit)
	}
	return stat.Exit
}
