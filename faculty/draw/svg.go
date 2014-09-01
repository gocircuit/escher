// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package draw

import (
	"bytes"
	"text/template"

	"github.com/gocircuit/escher/understand"
)

const svgImg = `
<?xml version="1.0" standalone="no"?>
<svg width="700px" height="700px" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
	viewBox="-1 -1 2 2">
	<defs><style type="text/css">@import url(http://fonts.googleapis.com/css?family=Lato);</style></defs>
	<circle cx="0" cy="0" r="0.01" stroke="none" fill="red" stroke-width="0"/>
	???
</svg>
`

func Draw(uc *understand.Circuit) string {
	var err error
	t := template.New("")
	if t, err = t.Parse(svgImg); err != nil {
		panic(err)
	}
	var w bytes.Buffer
	if err = t.Execute(&w, Compute(uc)); err != nil {
		panic(err)
	}
	return w.String()
}
