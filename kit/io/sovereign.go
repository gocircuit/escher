// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

// Package io...
package io

import (
	"io"
	"sync"
)

// SovereignReader returns a synchronized version of the argument reader.
func SovereignReader(x io.Reader) io.ReadCloser {
	switch t := x.(type) {
	case io.ReadCloser:
		return &sovereignReader{
			ReadCloser: t,
		}
	case io.Reader:
		return &sovereignReader{
			ReadCloser: NopReadCloser(t),
		}
	}
	panic(1)
}

type sovereignReader struct {
	sync.Mutex
	io.ReadCloser
}

func (x *sovereignReader) Read(p []byte) (int, error) {
	x.Lock()
	defer x.Unlock()
	return x.ReadCloser.Read(p)
}

func (x *sovereignReader) Close() error {
	x.Lock()
	defer x.Unlock()
	return x.ReadCloser.Close()
}

func NopReadCloser(x io.Reader) io.ReadCloser {
	return &nopReadCloser{x}
}

type nopReadCloser struct {
	io.Reader
}

func (x *nopReadCloser) Close() error {
	return nil
}

// SovereignWriter returns a synchronized version of the argument writer.
func SovereignWriter(x io.Writer) io.WriteCloser {
	switch t := x.(type) {
	case io.WriteCloser:
		return &sovereignWriter{
			WriteCloser: t,
		}
	case io.Writer:
		return &sovereignWriter{
			WriteCloser: NopWriteCloser(t),
		}
	}
	panic(1)
}

type sovereignWriter struct {
	sync.Mutex
	io.WriteCloser
}

func (x *sovereignWriter) Write(p []byte) (int, error) {
	x.Lock()
	defer x.Unlock()
	return x.WriteCloser.Write(p)
}

func (x *sovereignWriter) Close() error {
	x.Lock()
	defer x.Unlock()
	return x.WriteCloser.Close()
}

func NopWriteCloser(x io.Writer) io.WriteCloser {
	return &nopWriteCloser{x}
}

type nopWriteCloser struct {
	io.Writer
}

func (x *nopWriteCloser) Close() error {
	return nil
}
