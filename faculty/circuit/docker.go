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

	dkr "github.com/hoijui/circuit/client/docker"
	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
	"github.com/hoijui/escher/kit/plumb"
)

// Docker

// docker is the materialized docker reflex
type Docker struct {
	sync.Once                  // start backloop once
	spawn     chan interface{} // notify loop of spawn memes
}

func (p *Docker) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	p.spawn = make(chan interface{})
	return nil
}

func (p *Docker) CognizeCommand(eye *be.Eye, dvalue interface{}) {
	p.Once.Do(
		func() {
			back := &dockerBack{
				eye:   eye,
				cmd:   cognizeDockerCommand(dvalue),
				spawn: p.spawn,
			}
			go back.loop()
		},
	)
}

func (p *Docker) CognizeSpawn(eye *be.Eye, dvalue interface{}) {
	p.spawn <- dvalue
	log.Printf("circuit container spawning (%v)", cir.Linearize(fmt.Sprintf("%v", dvalue)))
}

func (p *Docker) CognizeExit(eye *be.Eye, dvalue interface{}) {}

func (p *Docker) CognizeIO(eye *be.Eye, dvalue interface{}) {}

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
	img, ok := v.(cir.Circuit)
	if !ok {
		panic(fmt.Sprintf("non-image sent as circuit container command (%v)", v))
	}
	cmd := &dkr.Run{}
	cmd.Image = img.StringAt("Image") // mandatory
	if mem, ok := img.IntOptionAt("Memory"); ok {
		cmd.Memory = int64(plumb.AsInt(mem))
	}
	if cpu, ok := img.IntOptionAt("CpuShares"); ok {
		cmd.CpuShares = int64(plumb.AsInt(cpu))
	}
	if lxc, ok := img.CircuitOptionAt("Lxc"); ok {
		for _, key := range lxc.Numbers() {
			cmd.Lxc = append(cmd.Lxc, lxc.StringAt(key))
		}
	}
	if vol, ok := img.CircuitOptionAt("Volume"); ok {
		for _, key := range vol.Numbers() {
			cmd.Volume = append(cmd.Volume, vol.StringAt(key))
		}
	}
	if entry, ok := img.StringOptionAt("Entry"); ok {
		cmd.Entry = entry
	}
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
	log.Printf("circuit docker command %v", cir.QuickPrint("", "t", -1, img))
	return cmd
}

type dockerBack struct {
	eye   *be.Eye
	cmd   *dkr.Run
	spawn <-chan interface{}
}

func (p *dockerBack) loop() {
	for {
		spwn := <-p.spawn
		x := cir.New().Grow("Spawn", spwn)
		if exit := p.spawnDocker(spwn); exit != nil {
			x.Grow("Exit", 1)
			p.eye.Show("Exit", x)
		} else {
			x.Grow("Exit", 0)
			p.eye.Show("Exit", x)
		}
		log.Printf("circuit container exit meme sent (%v)", cir.Linearize(fmt.Sprintf("%v", x)))
	}
}

func (p *dockerBack) spawnDocker(spwn interface{}) error {
	// anchor determination
	s := spwn.(cir.Circuit)
	anchor := program.Client.Walk(
		[]string{
			s.StringAt("Server"), // server name
			s.StringAt("Name"),   // (dynamic) execution name
		})
	//
	container, err := anchor.MakeDocker(*p.cmd)
	if err != nil {
		log.Fatalf("container spawn error (%v)", err)
	}
	defer anchor.Scrub() // Anchor will be scrubbed before the exit meme is sent out
	g := cir.New().
		Grow("Spawn", spwn).
		Grow("Stdin", container.Stdin()).
		Grow("Stdout", container.Stdout()).
		Grow("Stderr", container.Stderr())
	log.Printf("circuit docker io (%v)", cir.Linearize(fmt.Sprintf("%v", spwn)))
	p.eye.Show("IO", g)
	log.Printf("circuit docker waiting (%v)", cir.Linearize(fmt.Sprintf("%v", spwn)))
	stat, err := container.Wait()
	if err != nil {
		panic("circuit container wait aborted by user")
	}
	log.Printf("circuit container (%v) exited", cir.Linearize(fmt.Sprintf("%v", spwn)))
	var exit error
	if stat.State.ExitCode != 0 {
		exit = errors.New(fmt.Sprintf("circuit container exit code: %d", stat.State.ExitCode))
		log.Println(exit.Error())
	}
	return exit
}
