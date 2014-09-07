// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package understand

import (
	// "fmt"

	"github.com/gocircuit/escher/see"
	. "github.com/gocircuit/escher/image"
)

// Circuit ...
type Circuit struct {
	name  string // Name of the circuit design
	sourceDir string // Host source directory where this circuit's source implementation was found
	sourceFile string // File name
	genus []*see.Circuit // Union of syntactic circuits embodied in this semantic circuit
	peer Image // Peer name to peer structure; this circuit's super-peer participates with its name
}

func (c *Circuit) Name() string {
	return c.name
}

func (c *Circuit) SourceDir() string {
	return c.sourceDir
}

func (c *Circuit) SourceFile() string {
	return c.sourceFile
}

func (c *Circuit) PeerNames() []string {
	return c.peer.Letters()
}

func (c *Circuit) PeerByName(name string) *Peer {
	p, _ := c.peer.OptionalInterface(name).(*Peer)
	return p
}

// Peer ...
type Peer struct {
	name string
	index int
	design interface{}
	valve Image // Valve name to valve structure
}

func (p *Peer) Copy() *Peer {
	var q Peer = *p
	return &q
}

func (p *Peer) Name() string {
	return p.name
}

// Index returns the ordinal index of the clause containing the definition of this peer
// within the circuit's syntactic implementation
func (p *Peer) Index() int {
	return p.index
}

func (p *Peer) Design() interface{} {
	return p.design
}

func (p *Peer) ValveNames() []string {
	return p.valve.Letters()
}

func (p *Peer) ValveByName(name string) *Valve {
	v, _ := p.valve.OptionalInterface(name).(*Valve)
	return v
}

// Valve ...
type Valve struct {
	Of *Peer
	Name string
	// Ordinal index of the clause containing the valve's first occurence (in a matching)
	// within the circuit's syntactic implementation
	Index int
	Matching *Valve
}
