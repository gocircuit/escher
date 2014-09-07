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
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/see"
)

// Docker
type Docker struct{}

func (x Docker) Materialize(matter *be.Matter) be.Reflex {
	p := &docker{
		name: matter.LastName(),
		spawn: make(chan interface{}),
	}
	reflex, _ := plumb.NewEyeCognizer(p.cognize, "Command", "Spawn", "Exit", "IO")
	return reflex
}

// docker is the materialized docker reflex
type docker struct {
	name string // unique--with respect to materializations--name of this reflex
	sync.Once // start backloop once
	spawn chan interface{} // notify loop of spawn memes
}

func (p *docker) cognize(eye *plumb.Eye, dvalve string, dvalue interface{}) {
	switch dvalve {
	case "Command":
		p.Once.Do(
			func() {
				back := &dockerBack{
					name: p.name,
					eye: eye, 
					cmd: cognizeDockerCommand(dvalue), 
					spawn: p.spawn,
				}
				go back.loop()
			},
		)
	case "Spawn":
		p.spawn <- dvalue
		log.Printf("circuit container spawning (%v)", Linearize(fmt.Sprintf("%v", dvalue)))
	}
}

//	Command example:
//
//		{
//			Image "ubuntu64"
//			Memory 10e9
//			CpuShares 23
//			Lxc {}
//			Volume {
//				"/haha"
//				"/mnt/all"
//			}
//			Entry "entrypoint"
//			Env {
//				"PATH=/abc:/bin"
//				"LESS=less"
//			}
//			Dir "/home/petar"
//			Path "/bin/ls"
//			Args { "-l", "/" }
//		}
//
func cognizeDockerCommand(v interface{}) *dkr.Run {
	img, ok := v.(Image)
	if !ok {
		panic(fmt.Sprintf("non-image sent as circuit container command (%v)", v))
	}
	cmd := &dkr.Run{}
	cmd.Image = img.String(see.Name("Image")) // mandatory
	if img.Has(see.Name("Memory")) {
		cmd.Memory = int64(plumb.AsInt(img[see.Name("Memory")]))
	}
	if img.Has(see.Name("CpuShares")) {
		cmd.CpuShares = int64(plumb.AsInt(img[see.Name("CpuShares")]))
	}
	lxc := img.Walk(see.Name("Lxc"))
	for _, key := range see.Numbers(lxc) {
		cmd.Lxc = append(cmd.Lxc, lxc.String(key))
	}
	vol := img.Walk(see.Name("Volume"))
	for _, key := range see.Numbers(vol) {
		cmd.Volume = append(cmd.Volume, vol.String(key))
	}
	if img.Has(see.Name("Entry")) {
		cmd.Entry = img.String(see.Name("Entry"))
	}
	cmd.Path = img.String(see.Name("Path")) // mandatory
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
	log.Printf("circuit docker command (%v)", Linearize(img.Print("", "t")))
	return cmd
}

type dockerBack struct {
	name string
	eye *plumb.Eye
	cmd *dkr.Run
	spawn <-chan interface{}
}

func (p *dockerBack) loop() {
	for {
		spwn := <-p.spawn
		var x Image
		if exit := p.spawnDocker(spwn); exit != nil {
			x = Image{
				see.Name("Spawn"): spwn,
				see.Name("Exit"):  1,
			}
			p.eye.Show("Exit", x)
		} else {
			x = Image{
				see.Name("Spawn"): spwn,
				see.Name("Exit"):  0,
			}
			p.eye.Show("Exit", x)
		}
		log.Printf("circuit container exit meme sent (%v)", Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *dockerBack) spawnDocker(spwn interface{}) error {
	// anchor determination
	s := spwn.(Image)
	if s.String(see.Name("Name")) == "" {
		panic("container execution name cannot be empty")
	}
	anchor := program.Client.Walk(
		[]string{
			s.String(see.Name("Server")), // server name
			p.name, // reflex' unique materialization name
			s.String(see.Name("Name")), // (dynamic) execution name
		})
	//
	container, err := anchor.MakeDocker(*p.cmd)
	if err != nil {
		log.Fatalf("container spawn error (%v)", err)
	}
	defer anchor.Scrub() // Anchor will be scrubbed before the exit meme is sent out
	g := Image{
		see.Name("Spawn"):  spwn,
		see.Name("Stdin"):  container.Stdin(),
		see.Name("Stdout"): container.Stdout(),
		see.Name("Stderr"): container.Stderr(),
	}
	log.Printf("circuit docker io (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	p.eye.Show("IO", g)
	log.Printf("circuit docker waiting (%v)", Linearize(fmt.Sprintf("%v", spwn)))
	stat, err := container.Wait()
	if err != nil {
		panic("circuit container wait aborted by user")
	}
	log.Printf("circuit container (%v) exited", Linearize(fmt.Sprintf("%v", spwn)))
	var exit error
	if stat.State.ExitCode != 0 {
		exit = errors.New(fmt.Sprintf("circuit container exit code: %d", stat.State.ExitCode))
		log.Println(exit.Error())
	}
	return exit
}
