// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package browser

import (
	"encoding/json"
	"sync"
)

// Conn is a logical sub-connection within a physical websocket session.
type Conn struct {
	id   string
	ssn  *Session
	recv <-chan string
	send struct {
		sync.Mutex
		ch chan<- string
	}
}

func (u *Conn) ID() string {
	return u.id
}

// Clunk removes the connection from the session dispatcher
func (u *Conn) Clunk() {
	u.ssn.scrub(u.id) // scrub will panic if the session is already broken
}

// Inject will panic if the underlying session carrier is broken.
func (u *Conn) Inject(funcjs string, arg interface{}) {
	inj := Inject{
		Func: funcjs,
		Arg:  arg,
	}
	u.ssn.send(u.id, marshal(inj))
}

// accept, invoked by the session loop, notifies the Conn user of new feature messages
func (u *Conn) accept(pay string) {
	u.send.Lock()
	defer u.send.Unlock()
	u.send.ch <- pay
}

// clunk, invoked by the session loop, notifies the Conn user that the carrier has been closed
func (u *Conn) clunk() {
	u.send.Lock()
	defer u.send.Unlock()
	close(u.send.ch)
}

// Eject will retrieve data into the given data structure.
func (u *Conn) Eject(into interface{}) {
	pay, ok := <-u.recv
	if !ok {
		panic("end of session") // panic if the underlying session is broken
	}
	if into == nil {
		return
	}
	if err := json.Unmarshal([]byte(pay), into); err != nil {
		panic(err)
	}
}
