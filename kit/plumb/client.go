// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package plumb

import "sync"

type Client struct {
	req chan chan<- interface{}
	recognize func(interface{})
	sync.Mutex
}

func (c *Client) Init(recognize func(interface{})) {
	c.req = make(chan chan<- interface{}, 1)
	c.recognize = recognize
}

func (c *Client) Cognize(v interface{}) {
	select {
	case ch := <-c.req:
		ch <- v
	default:
		panic("received change without outstanding fetch")
	}
}

func (c *Client) Fetch(req interface{}) interface{} {
	c.Lock()
	defer c.Unlock()
	ch := make(chan interface{}, 1)
	c.req <- ch
	c.recognize(req)
	return <-ch
}
