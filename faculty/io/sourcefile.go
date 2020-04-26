// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package io

import (
	"log"
	"os"

	"github.com/hoijui/escher/be"
	cir "github.com/hoijui/escher/circuit"
)

func NewSourceFile(name string) be.Materializer {
	return be.NewMaterializer(SourceFile{}, name)
}

type SourceFile struct{}

func (SourceFile) Spark(eye *be.Eye, _ cir.Circuit, aux ...interface{}) cir.Value {
	go func() {
		name := aux[0].(string)
		file, err := os.Open(name)
		if err != nil {
			log.Printf("Problem opening file %q (%v)", name, err)
			panic("open file")
		}
		eye.Show(cir.DefaultValve, file)
	}()
	return nil
}

func (SourceFile) Cognize(eye *be.Eye, v interface{}) {}
