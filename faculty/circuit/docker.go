// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

import (
	"errors"
	"fmt"
	"log"
	"sync"

	dkr "github.com/gocircuit/circuit/client/docker"
	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/think"
)

// Docker
type Docker struct{}

func (x Docker) Materialize() think.Reflex {
	// Create reflex synapses
	cmdEndo, cmdExo := think.NewSynapse()
	spawnEndo, spawnExo := think.NewSynapse()
	exitEndo, exitExo := think.NewSynapse()
	ioEndo, ioExo := think.NewSynapse()
	serverEndo, serverExo := think.NewSynapse()
	//
	go func() {
		p := &docker{
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

// docker is the materialized docker reflex
type docker struct {
	id     string // ID of this docker reflex instance
	reExit *think.ReCognizer
	reIO   *think.ReCognizer
	arg    struct {
		sync.Mutex
		server string // root-level anchor of the server where the docker is to be started
		cmd    *dkr.Run
	}
	ready chan struct{}    // notify loop that arguments are ready
	spawn chan interface{} // notify loop of spawn strobes
}

func (p *docker) CognizeServer(v interface{}) {
	a, ok := v.(string)
	if !ok {
		panic("docker server anchor is non-string")
	}
	p.arg.Lock()
	defer p.arg.Unlock()
	if p.arg.server != "" {
		panic("docker server anchor already set")
	}
	p.arg.server = a
	if p.arg.cmd != nil {
		close(p.ready)
	}
}

//
//	{
//		Image "ubuntu64"
//		Memory 10e9
//		CpuShares 23
//		Lxc {
//			"??"
//			"??"
//		}
//		Volume {
//			"/haha"
//			"/mnt/all"
//		}
//		Entry "??"
//		Env {
//			"PATH=/abc:/bin"
//			"LESS=less"
//		}
//		Dir "/home/petar"
//		Path "/bin/ls"
//		Args { "-l", "/" }
//	}
//
func (p *docker) CognizeCommand(v interface{}) {
	img, ok := v.(Image)
	if !ok {
		log.Printf("Non-image sent to Docker.Command (%v)", v)
		return
	}
	p.arg.Lock()
	defer p.arg.Unlock()
	if p.arg.cmd != nil {
		panic("docker command already set")
	}
	p.arg.cmd = &dkr.Run{}
	p.arg.cmd.Image = img.String("Image") // mandatory
	if img.Has("Memory") {
		p.arg.cmd.Memory = int64(plumb.AsInt(img["Memory"]))
	}
	if img.Has("CpuShares") {
		p.arg.cmd.CpuShares = int64(plumb.AsInt(img["CpuShares"]))
	}
	lxc := img.Walk("Lxc")
	for _, key := range lxc.Sort() {
		p.arg.cmd.Lxc = append(p.arg.cmd.Lxc, lxc.String(key))
	}
	vol := img.Walk("Volume")
	for _, key := range vol.Sort() {
		p.arg.cmd.Volume = append(p.arg.cmd.Volume, vol.String(key))
	}
	if img.Has("Entry") {
		p.arg.cmd.Entry = img.String("Entry")
	}
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
	log.Printf("circuit docker command (%v)", Linearize(img.Print("", "t")))
	if p.arg.server != "" {
		close(p.ready)
	}
}

func (p *docker) CognizeSpawn(v interface{}) {
	p.spawn <- v
	log.Printf("circuit docker spawn (%v)", Linearize(fmt.Sprintf("%v", v)))
}

func (p *docker) loop() {
	<-p.ready // make sure arguments (command and server) have been received
	for {
		spwn := <-p.spawn
		var x Image
		if exit := p.spawnDocker(spwn); exit != nil {
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
		log.Printf("circuit docker exit sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *docker) spawnDocker(spwn interface{}) error {
	p.arg.Lock()
	server := p.arg.server
	cmd := p.arg.cmd
	p.arg.Unlock()
	//
	anchor := program.Client.Walk([]string{server, "escher", program.Name, "circuit.Docker", p.id})
	container, err := anchor.MakeDocker(*cmd)
	if err != nil {
		panic("invalid docker run argument")
	}
	defer anchor.Scrub()
	g := Image{
		"Spawn":  spwn,
		"Stdin":  container.Stdin(),
		"Stdout": container.Stdout(),
		"Stderr": container.Stderr(),
	}
	log.Printf("circuit docker io (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	p.reIO.ReCognize(g)
	log.Printf("circuit docker waiting (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	stat, err := container.Wait()
	if err != nil {
		panic("docker wait aborted by user")
	}
	log.Printf("circuit docker exit (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	var exit error
	if stat.State.ExitCode != 0 {
		return errors.New(fmt.Sprintf("docker exit code: %d", stat.State.ExitCode != 0))
	}
	return exit
}
