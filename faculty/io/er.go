// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package io

import (
	"io"
	"log"
	"os"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
	kitio "github.com/hoijui/escher/kit/io"
)

// Writer is a gate that reads from values sent to it and writes to an underlying writer.
type Writer struct {
	io.WriteCloser // sovereign writer
}

func NewWriterMaterializer(w io.Writer) be.Materializer {
	return be.NewMaterializer(&Writer{}, w)
}

func (x *Writer) Spark(eye *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	x.WriteCloser = kitio.SovereignWriter(aux[0].(io.Writer))
	for vlv := range matter.CircuitAt("View").Gate {
		go eye.Show(vlv, x.WriteCloser)
	}
	return nil
}

func (x *Writer) OverCognize(eye *be.Eye, _ cir.Name, value interface{}) {
	switch t := value.(type) {
	case io.Reader:
		go CopyClose(x.WriteCloser, t, false, true)
	case string:
		x.WriteCloser.Write([]byte(t))
	default:
		be.Panicf("unexpected type at writer origin (%T)", t)
	}
}

// Reader is a gate that reads from an underlying reader and writes to every object sent to it.
type Reader struct {
	io.ReadCloser // sovereign writer
}

func NewReaderMaterializer(r io.Reader) be.Materializer {
	return be.NewMaterializer(&Reader{}, r)
}

func (x *Reader) Spark(eye *be.Eye, matter cir.Circuit, aux ...interface{}) cir.Value {
	x.ReadCloser = kitio.SovereignReader(aux[0].(io.Reader))
	for vlv := range matter.CircuitAt("View").Gate {
		go eye.Show(vlv, x.ReadCloser)
	}
	return nil
}

func (x *Reader) OverCognize(_ *be.Eye, _ cir.Name, value interface{}) {
	switch t := value.(type) {
	case io.Writer:
		go CopyClose(t, x.ReadCloser, true, false)
	default:
		be.Panicf("reader gate doesn't know how to write to (%T)", t)
	}
}

// Util

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
