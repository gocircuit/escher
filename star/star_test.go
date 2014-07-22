// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package star

import (
	"testing"
	"fmt"
)

func TestStar(t *testing.T) {
	r := Make()
	r.Grow("hello").Grow("there").Grow("dolly")
	fmt.Println(r.Print("", "\t"))
	fmt.Println(r.Traverse("there").Print("", "\t"))
}
