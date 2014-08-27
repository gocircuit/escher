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
	return MaterializeWriterOrigin(os.Stdin)
}

type Stdout struct{}

func (Stdout) Materialize() be.Reflex {
	return MaterializeReaderOrigin(os.Stdout)
}

type Stderr struct{}

func (Stderr) Materialize() be.Reflex {
	return MaterializeReaderOrigin(os.Stderr)
}

func MaterializeWriterOrigin(w io.Writer) be.Reflex {
	x := &writerOrigin{
		WriteCloser: kitio.SovereignWriter(w),
	}
	reflex, eye := plumb.NewEyeCognizer(x.cognize, "_")
	go eye.Show("_", x.WriteCloser)
	return reflex
}

type writerOrigin struct{
	io.WriteCloser // sovereign writer
}

func (x *writerOrigin) cognize(eye *plumb.Eye, valve string, value interface{}) {
	switch t := value.(type) {
	case io.Reader:
		go func() {
			if _, err := io.Copy(x.WriteCloser, t); err != nil {
				log.Printf("writer origin drain problem (%s)", err)
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

func MaterializeReaderOrigin(w io.Reader) be.Reflex {
	x := &readerOrigin{
		ReadCloser: kitio.SovereignReader(w),
	}
	reflex, eye := plumb.NewEyeCognizer(x.cognize, "_")
	go eye.Show("_", x.ReadCloser)
	return reflex
}

type readerOrigin struct{
	io.ReadCloser // sovereign writer
}

func (x *readerOrigin) cognize(eye *plumb.Eye, valve string, value interface{}) {
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
