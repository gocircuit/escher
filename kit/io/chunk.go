// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package io

import (
	"bufio"
	"bytes"
	"io"
)

type ChunkReader struct {
	*bufio.Scanner
}

func NewChunkReader(r io.Reader) *ChunkReader {
	return &ChunkReader{Scanner: bufio.NewScanner(r)}
}

func (r *ChunkReader) Read() (chunk []byte, err error) {
	var n int
	var w bytes.Buffer
	for r.Scanner.Scan() {
		t := r.Scanner.Text()
		w.WriteString(t)
		if t == "\n" {
			n++
		} else {
			n = 0
		}
		if n == 2 {
			return w.Bytes(), nil
		}
	}
	return nil, io.ErrUnexpectedEOF
}
