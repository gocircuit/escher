// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package browser

import (
	"encoding/json"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"

	"github.com/petar/maymounkov.io/code.google.com/p/go.net/websocket"
)

// Session is a websocket session abstraction
type Session struct {
	s, r    sync.Mutex
	carrier *websocket.Conn
	conn    struct {
		sync.Mutex
		id map[string]*Conn
	}
}

// NewSession creates a new message-passing session over the given websocket
func NewSession(carrier *websocket.Conn) *Session {
	ssn := &Session{carrier: carrier}
	ssn.conn.id = make(map[string]*Conn)
	runtime.SetFinalizer(ssn, func(x *Session) {
		log.Println("bye, collecting websocket session")
		x.carrier.Close()
	})
	go ssn.loop()
	return ssn
}

func (ssn *Session) send(id string, pay string) {
	ssn.s.Lock()
	defer ssn.s.Unlock()
	msg := &Msg{
		ID:  id,
		Pay: pay,
	}
	if err := websocket.JSON.Send(ssn.carrier, msg); err != nil {
		panic(err) // panic into Conn.Send and into the user code (reflex circuit, e.g.)
	}
}

func (ssn *Session) recv() (id, pay string) {
	ssn.r.Lock()
	defer ssn.r.Unlock()
	var msg = &Msg{}
	if err := websocket.JSON.Receive(ssn.carrier, msg); err != nil {
		panic(err)
	}
	return msg.ID, msg.Pay
}

// loop demultiplexes incoming messages to connections
func (ssn *Session) loop() {
	defer func() {
		recover() // Kill the loop quietly (on carrier death) and cleanup
		ssn.clunk()
	}()
	for {
		id, pay := ssn.recv()
		u := ssn.lookup(id)
		if u == nil {
			log.Printf("no connection [%s] for message [%s]", id, pay)
			continue
		}
		u.accept(pay)
	}
}

func (ssn *Session) lookup(id string) *Conn {
	ssn.conn.Lock()
	defer ssn.conn.Unlock()
	return ssn.conn.id[id]
}

func (ssn *Session) save(conn *Conn) {
	ssn.conn.Lock()
	defer ssn.conn.Unlock()
	ssn.conn.id[conn.id] = conn
}

func (ssn *Session) scrub(id string) {
	ssn.conn.Lock()
	defer ssn.conn.Unlock()
	delete(ssn.conn.id, id)
}

func (ssn *Session) clunk() { // called only from loop
	ssn.carrier.Close()
	ssn.conn.Lock()
	defer ssn.conn.Unlock()
	for _, u := range ssn.conn.id {
		u.clunk() // Notify conns that they're dead
	}
	ssn.conn.id = nil // Cause post-mortem invocations to New to panic
}

func marshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (ssn *Session) new(id string) *Conn {
	var (
		ch = make(chan string, 1)
		u  = &Conn{
			id:   id,
			ssn:  ssn,
			recv: ch, // channel for JSON-encoded strings
		}
	)
	u.send.ch = ch
	ssn.save(u)
	return u
}

func (ssn *Session) Dial() *Conn {
	return ssn.new(strconv.FormatInt(rand.Int63(), 21))
}
