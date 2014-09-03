// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	"fmt"
	"hash/fnv"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ChainKey(superKey string, fullName []string) string {
	h := fnv.New32a()
	h.Write([]byte(superKey))
	h.Write([]byte("·"))
	h.Write([]byte(strings.Join(fullName, ".")))
	return strconv.FormatUint(uint64(h.Sum32()), 36) + ":" + fullName[len(fullName)-1]
}

func stk() {
	prof := pprof.Lookup("goroutine")
	prof.WriteTo(os.Stderr, 1)
}

// panicf is a quick/lazy way to report errors with their reason stacks
func panicf(format string, arg ...interface{}) {
	panic(fmt.Sprintf(format, arg...))
}

// Expire is a device that invokes an expiration function,
// if a termination condition is not met in time.
type Expire struct {
	sync.Mutex
	done bool
}

func (w *Expire) Init(late func()) {
	go func() {
		time.Sleep(time.Second / 3)
		w.ifNotDone(late)
	}()
}

func (w *Expire) InitString(s string) {
	go func() {
		time.Sleep(time.Second / 3)
		w.ifNotDone(func() {
			println(s)
		})
	}()
}

func (w *Expire) Done() {
	w.Lock()
	defer w.Unlock()
	w.done = true
}

func (w *Expire) ifNotDone(g func()) {
	w.Lock()
	defer w.Unlock()
	if w.done {
		return
	}
	g()
}
