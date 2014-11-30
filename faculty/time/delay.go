// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package time

import (
	"sync"
	"time"

	"github.com/gocircuit/escher/be"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/faculty"
	"github.com/gocircuit/escher/kit/plumb"
)

func init() {
	faculty.Register(be.NewMaterializer(&Ticker{}), "time", "Ticker")
	faculty.Register(be.NewMaterializer(&Delay{}), "time", "Delay")
}

// Delay…
type Delay struct {
	sync.Mutex
	dur time.Duration
}

func (t *Delay) Spark(*be.Eye, Circuit, ...interface{}) Value {
	return nil
}

func (t *Delay) delay() time.Duration {
	t.Lock()
	defer t.Unlock()
	return t.dur
}

func (t *Delay) CognizeDuration(eye *be.Eye, value interface{}) {
	t.dur = time.Duration(plumb.AsInt(value))
}

func (t *Delay) CognizeNorth(eye *be.Eye, value interface{}) {
	time.Sleep(t.delay())
	eye.Show("South", value)
}

func (t *Delay) CognizeSouth(eye *be.Eye, value interface{}) {
	time.Sleep(t.delay())
	eye.Show("North", value)
}
