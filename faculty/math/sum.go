// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package math

import (
	"sync"

	// "github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
	// "github.com/gocircuit/escher/kit/plumb"
)

// IntSum
type IntSum struct{
	sync.Mutex
	x, y, sum int
}

func (s *IntSum) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return &IntSum{}
}

func (s *IntSum) save(valve string, value int) (x, y, sum int) {
	s.Lock()
	defer s.Unlock()
	switch valve {
	case "X":
		s.x = value
		s.y, s.sum = s.sum - s.x, s.x + s.y
	case "Y":
		s.y = value
		s.x, s.sum = s.sum - s.y, s.x + s.y
	case "Sum":
		s.sum = value
		s.x, s.y = s.sum - s.y, s.sum - s.x
	}
	return s.x, s.y, s.sum
}

func (s *IntSum) CognizeX(eye *be.Eye, v interface{}) {
	_, y, sum := s.save("X", v.(int))
	s.fire(eye, "Y", "Sum", y, sum)
}

func (s *IntSum) CognizeY(eye *be.Eye, v interface{}) {
	x, _, sum := s.save("Y", v.(int))
	s.fire(eye, "X", "Sum", x, sum)
}

func (s *IntSum) CognizeSum(eye *be.Eye, v interface{}) {
	x, y, _ := s.save("Sum", v.(int))
	s.fire(eye, "X", "Y", x, y)
}

func (s *IntSum) fire(eye *be.Eye, v1, v2 string, u1, u2 int) {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(2)
	go func() {
		defer func() {
			recover()
		}()
		defer wg.Done()
		eye.Show(v1, u1)
	}()
	go func() {
		defer func() {
			recover()
		}()
		defer wg.Done()
		eye.Show(v2, u2)
	}()
}
