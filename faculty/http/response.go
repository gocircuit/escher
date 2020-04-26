// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	cir "github.com/hoijui/escher/circuit"
)

// cognizeResponse reads the circuit response u and fills in the http header object, returning the status and body.
func (s *Server) cognizeResponse(header http.Header, u cir.Circuit) (status int, body io.ReadCloser, ok bool) {

	// Status
	if status, ok = u.IntOptionAt("Status"); !ok {
		return
	}

	// Header
	var h cir.Circuit
	if h, ok = u.CircuitOptionAt("Header"); !ok {
		return
	}
	for _, k := range h.SortedLetters() {
		g, ok := h.CircuitOptionAt(k)
		if !ok {
			continue
		}
		header[k] = circuitSlice(g)
	}

	// Body gate should be convertible to string
	var v cir.Value
	if v, ok = u.OptionAt("Body"); !ok {
		return
	}
	// var m be.Materializer
	// if m, ok = v.(be.Materializer); ok { // extract body from a noun reflex, if a materializer for one is given
	// 	x, _ := be.MaterializeReflex(be.NewIndex(), m, s.matter)
	// 	synapse, ok := x[DefaultValve]
	// 	if !ok {
	// 		panic("expecting reflex with one default valve")
	// 	}
	// 	ch := make(chan interface{}, 1)
	// 	synapse.Connect(func(w interface{}) { ch <- w })
	// 	v = <-ch
	// }
	switch t := v.(type) {
	case string:
		body = ioutil.NopCloser(bytes.NewBufferString(t))
	case []byte:
		body = ioutil.NopCloser(bytes.NewBuffer(t))
	case io.Reader:
		body = ioutil.NopCloser(t)
	case io.ReadCloser:
		body = t
	default:
		panic("unrecognized http response body type")
	}
	ok = true
	return
}

func circuitSlice(u cir.Circuit) []string {
	var ss []string
	for _, j := range u.SortedNumbers() {
		ss = append(ss, fmt.Sprintf("%v", u.At(j)))
	}
	return ss
}
