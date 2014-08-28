// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package os

import (
	"io"
	"log"
	"os"

	// "github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/be"
	"github.com/gocircuit/escher/kit/plumb"
	kitio "github.com/gocircuit/escher/kit/io"
)

type Stdin struct{}

func (Stdin) Materialize() be.Reflex {
	return MaterializeReadFrom(os.Stdin)
}

type Stdout struct{}

func (Stdout) Materialize() be.Reflex {
	return MaterializeWriteTo(os.Stdout)
}

type Stderr struct{}

func (Stderr) Materialize() be.Reflex {
	return MaterializeWriteTo(os.Stderr)
}

func MaterializeWriteTo(w io.Writer) be.Reflex {
	x := &writerTo{
		WriteCloser: kitio.SovereignWriter(w),
	}
	reflex, eye := plumb.NewEyeCognizer(x.cognize, "_")
	go eye.Show("_", x.WriteCloser)
	return reflex
}

type writerTo struct{
	io.WriteCloser // sovereign writer
}

func (x *writerTo) cognize(eye *plumb.Eye, valve string, value interface{}) {
	switch t := value.(type) {
	case io.Reader:
		go func() {
			_, err := io.Copy(x.WriteCloser, t)
			if err != nil {
				log.Printf("writer origin drain problem (%s)", err)
			}
			if tt, ok := t.(*os.File); ok {
				tt.Sync()
			}
			if tt, ok := t.(io.Closer); ok {
				tt.Close()
			}
			// We don't close the writer so it can be reused.
		}()
	default:
		log.Printf("unexpected type at writer origin (%T)", t)
	}
}

func MaterializeReadFrom(w io.Reader) be.Reflex {
	x := &readFrom{
		ReadCloser: kitio.SovereignReader(w),
	}
	reflex, eye := plumb.NewEyeCognizer(x.cognize, "_")
	go eye.Show("_", x.ReadCloser)
	return reflex
}

type readFrom struct{
	io.ReadCloser // sovereign writer
}

func (x *readFrom) cognize(eye *plumb.Eye, valve string, value interface{}) {
	switch t := value.(type) {
	case io.Writer:
		go func() {
			if _, err := io.Copy(t, x.ReadCloser); err != nil {
				log.Printf("reader origin drain problem (%s)", err)
			}
			if tt, ok := t.(io.Closer); ok {
				tt.Close()
			}
			// We don't close the reader so it can be reused.
		}()
	default:
		log.Printf("unexpected type at reader origin (%T)", t)
	}
}
