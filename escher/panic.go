// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"time"
)

func init() {
	//InstallCtrlCPanic()
	//InstallGoroutineWatch()
}

// InstallCtrlCPanic installs a Ctrl-C signal handler that panics
func InstallCtrlCPanic() {
	go func() {
		//defer SavePanicTrace()
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		for _ = range ch {
			println("ctrl/c")
			// prof := pprof.Lookup("goroutine")
			// prof.WriteTo(os.Stderr, 2)
			os.Exit(1)
		}
	}()
}

func InstallGoroutineWatch() {
	go func() {
		for {
			time.Sleep(3 * time.Second)
			println(fmt.Sprintf("goroutines=%d", runtime.NumGoroutine()))
		}
	}()
}

func fatalf(format string, arg ...interface{}) {
	println(fmt.Sprintf(format, arg...))
	os.Exit(1)
}

func stack() {
	var w bytes.Buffer
	p := pprof.Lookup("goroutine")
	p.WriteTo(&w, 1)
	println(w.String())
}
