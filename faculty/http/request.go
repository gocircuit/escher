// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package http

import (
	// "fmt"
	"net/http"
	"strings"

	. "github.com/gocircuit/escher/circuit"
)

// requestCircuit converts an http.Request object into a data circuit representation
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
