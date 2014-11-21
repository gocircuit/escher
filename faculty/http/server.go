// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package http

import (
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
)

func init() {
	faculty.Register("http.Server", be.NewNativeMaterializer(&Server{}))
}

type Server struct {
	eye *be.Eye
	matter *be.Matter
	sync.Mutex
	server *http.Server
	throttle chan struct{}
}

func (s *Server) Spark(eye *be.Eye, matter *be.Matter, aux ...interface{}) Value {
	s.eye, s.matter = eye, matter
	const throttle = 50
	s.throttle = make(chan struct{}, throttle)
	for i := 0; i < throttle; i++ {
		s.throttle <- struct{}{}
	}
	return nil
}

func (s *Server) CognizeRequestResponse(eye *be.Eye, value interface{}) {}

func (s *Server) CognizeStart(eye *be.Eye, value interface{}) {
	s.Lock()
	defer s.Unlock()
	//
	u := value.(Circuit)
	if s.server != nil {
		panic("server running")
	}
	s.server = &http.Server{
		Addr: u.StringAt("Address"),
		Handler: s,
	}
	go func() {
		s.server.ListenAndServe()
	}()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	<-s.throttle
	defer func() {
		s.throttle <- struct{}{}
	}()
	//
	mx, my := be.NewEntanglement()
	ch := make(chan struct{}, 1)
	go mx.Synapse().Focus(
		func (v interface{}) {
			defer func() {
				ch <- struct{}{}
			}()
			status, body, ok := s.cognizeResponse(w.Header(), v.(Circuit))
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Escher web server: App error."))
				return
			}
			defer body.Close()
			w.WriteHeader(status)
			io.Copy(w, body)
		},
	)
	s.eye.Show(
		"RequestResponse", 
		New().
			Grow("Request", requestCircuit(req)).
			Grow("Respond", my),
	)
	<-ch
}

// body is either string, reader or a materializer whose default valve returns one of those
func (s *Server) cognizeResponse(header http.Header, u Circuit) (status int, body io.ReadCloser, ok bool) {
	
	// Status
	if status, ok = u.IntOptionAt("Status"); !ok {
		return
	}

	// Header
	var h Circuit
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

	// Body
	var v Value
	if v, ok = u.OptionAt("Body"); !ok {
		return 
	}
	var m be.Materializer
	if m, ok = v.(be.Materializer); ok { // extract body from a noun reflex, if a materializer for one is given
		x, _ := be.MaterializeReflex(be.NewIndex(), m, s.matter)
		synapse, ok := x[DefaultValve]
		if !ok {
			panic("expecting reflex with one default valve")
		}
		ch := make(chan interface{}, 1)
		synapse.Focus(func(w interface{}) { ch <- w })
		v = <-ch
	}
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

func requestCircuit(req *http.Request) Circuit {
	x := New()

	// HTTP method
	x.Gate["Method"] = req.Method

	// URL path
	var nn []Name
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}
	if len(parts) == 1 && parts[0] == "" {
		parts = []string{}
	}
	for _, n := range parts {
		nn = append(nn, n)
	}
	x.Gate["Path"] = NewAddress(nn...)

	// URL query
	v := New()
	for k, ss := range req.URL.Query() {
		v.Gate[k] = sliceCircuit(ss)
	}
	x.Gate["Query"] = v

	return x
}

func sliceCircuit(ss []string) Circuit {
	x := New()
	for i, v := range ss {
		x.Gate[i] = v
	}
	return x
}

func circuitSlice(u Circuit) []string {
	var ss []string
	for _, j := range u.SortedNumbers() {
		ss = append(ss, fmt.Sprintf("%v", u.At(j)))
	}
	return ss
}
