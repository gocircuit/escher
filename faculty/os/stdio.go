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

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	kitio "github.com/gocircuit/escher/kit/io"
)

type Stdin struct{}

func (Stdin) Materialize(*be.Matter) (be.Reflex, Value) {
	return MaterializeReadFrom(os.Stdin)
}

type Stdout struct{}

func (Stdout) Materialize(*be.Matter) (be.Reflex, Value) {
	return MaterializeWriteTo(os.Stdout)
}

type Stderr struct{}

func (Stderr) Materialize(*be.Matter) (be.Reflex, Value) {
	return MaterializeWriteTo(os.Stderr)
}

func MaterializeWriteTo(w io.Writer) (be.Reflex, Value) {
	x := &writerTo{
		WriteCloser: kitio.SovereignWriter(w),
	}
	reflex, eye := be.NewEyeCognizer(x.cognize, DefaultValve)
	go eye.Show(DefaultValve, x.WriteCloser)
	return reflex, nil
}

type writerTo struct {
	io.WriteCloser // sovereign writer
}

func (x *writerTo) cognize(eye *be.Eye, valve Name, value interface{}) {
	switch t := value.(type) {
	case io.Reader:
		go CopyClose(x.WriteCloser, t, false, true)
	case string:
		x.WriteCloser.Write([]byte(t))
	default:
		log.Printf("unexpected type at writer origin (%T)", t)
	}
}

func CopyClose(w io.Writer, r io.Reader, closeWriter, closeReader bool) {
	_, err := io.Copy(w, r)
	if err != nil {
		log.Printf("draining problem (%s)", err)
	}
	if tt, ok := w.(*os.File); ok {
		tt.Sync()
	}
	if closeReader {
		if tt, ok := r.(io.Closer); ok {
			tt.Close()
		}
	}
	if closeWriter {
		if tt, ok := w.(io.Closer); ok {
			tt.Close()
		}
	}
}

func MaterializeReadFrom(w io.Reader) (be.Reflex, Value) {
	x := &readFrom{
		ReadCloser: kitio.SovereignReader(w),
	}
	reflex, eye := be.NewEyeCognizer(x.cognize, DefaultValve)
	go eye.Show(DefaultValve, x.ReadCloser)
	return reflex, nil
}

type readFrom struct {
	io.ReadCloser // sovereign writer
}

func (x *readFrom) cognize(eye *be.Eye, valve Name, value interface{}) {
	switch t := value.(type) {
	case io.Writer:
		go CopyClose(t, x.ReadCloser, true, false)
	default:
		log.Printf("unexpected type at reader origin (%T)", t)
	}
}
