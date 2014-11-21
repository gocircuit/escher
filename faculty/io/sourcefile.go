// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package io

import (
	"log"
	"os"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	// "github.com/gocircuit/escher/faculty"
)

func NewSourceFileMaterializer(name string) be.Materializer {
	return be.NewNativeMaterializer(SourceFile{}, name)
}

type SourceFile struct{}

func (SourceFile) Spark(eye *be.Eye, _ *be.Matter, aux ...interface{}) Value {
	go func() {
		name := aux[0].(string)
		file, err := os.Open(name)
		if err != nil {
			log.Printf("Problem opening file %q (%v)", name, err)
			panic("open file")
		}
		eye.Show(DefaultValve, file)
	}()
	return nil
}

func (SourceFile) Cognize(eye *be.Eye, v interface{}) {}
