// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package faculty

// Root is a global variable where packages can add gates as side-effect of being imported.
var Root = NewFaculty("root")
