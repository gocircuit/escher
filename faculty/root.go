// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

import (
	. "github.com/gocircuit/escher/circuit"
)

// Root is the global faculties memory where Go packages add gate designs as side-effect of being imported.
var Root = New()
