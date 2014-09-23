// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package shell

import (
	"os"
	"testing"
)

func TestShell(t *testing.T) {
	NewShell(os.Stdin, os.Stdout, os.Stderr, nil)
	// select{}
}
