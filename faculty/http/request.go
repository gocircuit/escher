// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package http

import (
	"net/http"
	"strings"

	cir "github.com/hoijui/escher/circuit"
)

// requestCircuit converts an http.Request object into a data circuit representation
func requestCircuit(req *http.Request) cir.Circuit {
	x := cir.New()

	// HTTP method
	x.Gate["Method"] = req.Method

	// URL path
	var nn []cir.Name
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
	x.Gate["Path"] = cir.NewAddress(nn...)

	// URL query
	v := cir.New()
	for k, ss := range req.URL.Query() {
		v.Gate[k] = sliceCircuit(ss)
	}
	x.Gate["Query"] = v

	return x
}

func sliceCircuit(ss []string) cir.Circuit {
	x := cir.New()
	for i, v := range ss {
		x.Gate[i] = v
	}
	return x
}
