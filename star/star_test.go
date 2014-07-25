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
	s := Make() // singleton
	s.Show(1)
	fmt.Println(s.Print("", "\t"))
	s = s.Traverse("fwd1", "rev1")
	fmt.Println(s.Print("", "\t"))
	s = s.Traverse("fwd2", "rev2")
	fmt.Println(s.Print("", "\t"))
	s1, s2 := s.Split("rev2", "fwd2")
	fmt.Println("split")
	fmt.Println(s1.Print("", "\t"))
	fmt.Println(s2.Print("", "\t"))
	r := s2.Copy()
	fmt.Println(r.Print("", "\t"))
	r = r.Merge("ha", "ha", s1)
	fmt.Println(r.Print("", "\t"))
}
