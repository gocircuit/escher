// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package shell

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/kit/memory"
	"github.com/gocircuit/escher/see"
)

// TODO:
// * mk take path arguments
// * peek's second (destination) argument can be path

type Shell struct {
	name string
	in io.Reader // io channels for interaction with user
	out io.WriteCloser
	err io.WriteCloser
	scan *bufio.Scanner
	space Memory
	at []Name
}

func NewShell(name string, in io.Reader, out, err io.WriteCloser) *Shell {
	sh := &Shell{
		name: name,
		in: in,
		out: out,
		err: err,
		scan: bufio.NewScanner(in),
		at: []Name{},
	}
	return sh
}

func (sh *Shell) Start(value Circuit) {
	sh.space = Memory(value)
	sh.prompt()
	for sh.scan.Scan() {
		words := split(sh.scan.Text())
		if len(words) == 0 {
			sh.prompt()
			continue
		}
		switch strings.ToLower(words[0]) {
		case "help", "h", "?":
			sh.help(words[1:])
		case "ls":
			sh.ls(words[1:])
		case "cd":
			sh.cd(words[1:])
		case "mk":
			sh.mk(words[1:])
		case "rm":
			sh.rm(words[1:])
		case "p", "pwd":
			sh.path(words[1:])
		case "l", "link":
			sh.link(words[1:])
		case "u", "unlink":
			sh.unlink(words[1:])
		case "peek":
			sh.peek(words[1:])
		default:
			fmt.Fprintf(sh.err, "command not recognized; try help\n")
		}
		sh.prompt()
	}
}

func (sh *Shell) memory() Memory {
	return sh.space
}

func (sh *Shell) prompt() {
	fmt.Fprintf(sh.err, "%s·%s ", sh.name, printPath(sh.at))
}

func (sh *Shell) cd(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "path not found\n")
		}
	}()
	switch {
	case len(w) == 0:
		sh.at = []Name{}
	case len(w) == 1:
		pov, _ := sh.glob(w[0])
		sh.memory().Goto(pov...)
		sh.at = pov
	default:
		fmt.Fprintf(sh.err, "cd accepts at most one argument\n")
	}
}

func (sh *Shell) peek(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "path not found\n")
		}
	}()
	switch {
	case len(w) == 2:
		pov, _ := sh.glob(w[0])
		to, _ := sh.glob(w[1])
		//
		x := sh.memory().Lookup(Address{pov})
		if p, ok := x.(interface{ Peek() Circuit }); ok {
			sh.space.Goto(to[:len(to)-1]...).Include(to[len(to)-1], p.Peek())
		} else {
			fmt.Fprintf(sh.err, "object of type %T does not have a peek method\n", x)
		}
	default:
		fmt.Fprintf(sh.err, "peek accepts two arguments: a path and a new unique focus name\n")
	}
}

func (sh *Shell) memoryAt() Memory {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "current path is disconnected from root\n")
		}
	}()
	return sh.memory().Goto(sh.at...)
}

func (sh *Shell) ls(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "path not found\n")
		}
	}()
	switch {
	case len(w) == 0:
		fmt.Fprintf(sh.err, "%v\n", Circuit(sh.memoryAt()).Print("", "   ", 1))
	case len(w) == 1:
		pov, ell := sh.glob(w[0])
		recurse := 1
		if ell {
			recurse = -1
		}
		fmt.Fprintf(sh.err, "%v\n", Circuit(sh.memory().Goto(pov...)).Print("", "   ", recurse))
	default:
		fmt.Fprintf(sh.err, "ls accepts at most one argument\n")
	}
}

func (sh *Shell) path(w []string) {
	fmt.Fprintf(sh.err, "%s\n", printPath(sh.at))
}

func (sh *Shell) unlink(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "unlink error\n")
		}
	}()
	sh.memoryAt().Unlink(parseLink(w))
}

func (sh *Shell) link(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "link error\n")
		}
	}()
	sh.memoryAt().Link(parseLink(w))
}

func (sh *Shell) mk(w []string) {
	switch {
	case len(w) == 0:
		fmt.Fprintf(sh.err, "mk command needs arguments\n")
	case len(w) == 1:
		t, _ := sh.glob(w[0])
		dir := sh.memory().Goto(t[:len(t)-1]...)
		if old := Circuit(dir).Include(t[len(t)-1], New()); old != nil {
			fmt.Fprintf(sh.err, "Displaced: %v\n", old)
		}
	case len(w) == 2:
		t, _ := sh.glob(w[0])
		x := see.SeeValueOrNil(see.NewSrcString(w[1]))
		if x == nil {
			fmt.Fprintf(sh.err, "Value not recognized\n")
			return
		}
		dir := sh.memory().Goto(t[:len(t)-1]...)
		if old := Circuit(dir).Include(t[len(t)-1], x); old != nil {
			fmt.Fprintf(sh.err, "Displaced: %v\n", old)
		}
	default:
		fmt.Fprintf(sh.err, "mk accepts at most two arguments\n")
	}
}

func (sh *Shell) rm(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "remove issue\n")
		}
	}()
	if len(w) != 1 {
		fmt.Fprintf(sh.err, "rm accepts one argument\n")
		return
	}
	pov, _ := sh.glob(w[0])
	if len(pov) == 0 {
		fmt.Fprintf(sh.err, "cannot remove root\n")
		return
	}
	sh.memory().Goto(pov[:len(pov)-1]...).Exclude(pov[len(pov)-1])
}

func (sh *Shell) help(w []string) {
	const help = `
a             Accept next value
p             Show current path
ls            Show circuit in current focus
ls ../ef/     Show circuit at path relative to current
peek a/b c/d  Peek addressed object if peeking supported
cd            Move current focus to root memory circuit
cd /          "
cd ef/gh      Move current focus relative to itself
cd ..         Move current focus to parent memory circuit
mk xyz        Make a memory gate named "xyz"
mk xyz "abc"  Make a gate named "xyz"
mk xyz 123    "
mk xyz {a 1}  "
rm abc        Remove gate named "abc"
l x:a = y:b   Link gate valves
u x:a = y:b   Unlink gate valves
`
	fmt.Fprintf(sh.err, "%s\n", help)
}
