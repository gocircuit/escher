// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package shell

import (
	"bytes"
	"errors"
	"fmt"
	// "path"
	"strings"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/see"
)

func printPath(p []Name) string {
	if len(p) == 0 {
		return "/"
	}
	var w bytes.Buffer
	for _, x := range p {
		fmt.Fprintf(&w, "/%v", x)
	}
	return w.String()
}

func split(src string) (r []string) {
	var s int
	for i, b := range src {
		if b != ' ' && b != '\t' {
			continue
		}
		switch {
		case i > s:
			r = append(r, src[s:i])
			s = i + 1
		case i == s:
			s++
		}
	}
	if len(src) > s {
		r = append(r, src[s:])
	}
	return
}

func glob(s string) (walk []string, ellipses bool, err error) {
	for _, b := range s {
		if !see.IsIdentifier(rune(b)) && b != '/' && b != '.' {
			return nil, false, errors.New("glob characters")
		}
	}
	if strings.HasSuffix(s, "...") {
		ellipses = true
		s = s[:len(s) - len("...")]
	}
	walk = strings.Split(s, "/")
	return
}

func derelativize(walk []string, pov []Name) ([]Name, bool) {
	if len(walk) > 1 && walk[0] == "" {
		pov = []Name{}
	}
	for _, w := range walk {
		switch w {
		case "..":
			if len(pov) == 0 {
				return nil, false
			}
			pov = pov[:len(pov)-1]
		case ".", "":
		default:
			pov = append(pov, w)
		}
	}
	return pov, true
}

func (sh *Shell) glob(w string) (pov []Name, ell bool) {
	walk, ell, err := glob(w)
	if err != nil {
		fmt.Fprintf(sh.err, "glob not recognized (%s)\n", err)
		panic(err)
	}
	pov, ok := derelativize(walk, sh.focus().Path)
	if !ok {
		fmt.Fprintf(sh.err, "path not valid\n")
		panic(0)
	}
	return pov, ell
}

func parseLink(w []string) (x, y Vector) {
	v, cry := see.SeeLink(see.NewSrcString(strings.Join(w, " ")), 0)
	if v == nil {
		panic("link not recognized")
	}
	if cry[0] != nil || cry[1] != nil {
		panic("link is not simple")
	}
	return v[0], v[1]
}
