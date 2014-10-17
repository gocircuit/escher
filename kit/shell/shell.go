// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package shell

import (
	"bufio"
	"fmt"
	"log"
	"io"
	"path"
	"sort"
	"strings"

	. "github.com/gocircuit/escher/circuit"
	// . "github.com/gocircuit/escher/be"
	// . "github.com/gocircuit/escher/kit/fs"
	. "github.com/gocircuit/escher/kit/memory"
	"github.com/gocircuit/escher/see"
)

type Shell struct {
	in io.Reader // io channels for interaction with user
	out io.WriteCloser
	err io.WriteCloser
	scan *bufio.Scanner
	foci map[string]*Focus // list of foci under management
	current string // current focus
}

type Focus struct {
	Name string
	Path []Name
	Memory Memory
}

func NewShell(in io.Reader, out, err io.WriteCloser) *Shell {
	sh := &Shell{
		in: in,
		out: out,
		err: err,
		scan: bufio.NewScanner(in),
		foci: make(map[string]*Focus),
		current: "",
	}
	return sh
}

func (sh *Shell) StartSession(name string, memory Circuit) {
	sh.attach(name, memory)
	sh.jump([]string{name})
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
		case "v", "view":
			sh.view(words[1:])
		case "jump", "jmp", "j":
			sh.jump(words[1:])
		case "l", "link":
			sh.link(words[1:])
		case "u", "unlink":
			sh.unlink(words[1:])
		case "recycle", "r":
			// exit the shell loop, thereby blocking the user experience until
			// the next session is started
			return
		case "peek":
			sh.peek(words[1:])
		default:
			fmt.Fprintf(sh.err, "command not recognized; try help\n")
		}
		sh.prompt()
	}
}

func (sh *Shell) memory() Memory {
	return sh.foci[sh.current].Memory
}

func (sh *Shell) focus() *Focus {
	return sh.foci[sh.current]
}

func (sh *Shell) prompt() {
	fmt.Fprintf(sh.err, "%s·%s ", sh.current, printPath(sh.focus().Path))
}

func (sh *Shell) attach(name string, value Circuit) {
	f, ok := sh.foci[name]
	if ok {
		log.Fatalf("shell already has focus named %s", name)
	}
	f = &Focus{
		Name: name,
		Path: []Name{},
		Memory: Memory(value),
	}
	sh.foci[name] = f
}

func (sh *Shell) jump(w []string) {
	if len(w) != 1 {
		fmt.Fprintf(sh.err, "jump requires one argument\n")
		return
	}
	f, ok := sh.foci[w[0]]
	if !ok {
		f = &Focus{
			Name: w[0],
			Path: sh.focus().Path,
			Memory: Memory(New()),
		}
		sh.foci[w[0]] = f
	}
	sh.current = w[0]
}

func (sh *Shell) view(w []string) {
	var ord []string
	for f, _ := range sh.foci {
		ord = append(ord, f)
	}
	sort.Strings(ord)
	for _, f := range ord {
		x := sh.foci[f]
		if f == sh.current {
			fmt.Fprintf(sh.err, "* %s:%s\n", f, printPath(x.Path))
		} else {
			fmt.Fprintf(sh.err, "  %s:%s\n", f, printPath(x.Path))
		}
	}
}

func (sh *Shell) cd(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "path not found\n")
		}
	}()
	switch {
	case len(w) == 0:
		sh.focus().Path = []Name{}
	case len(w) == 1:
		pov, _ := sh.glob(w[0])
		sh.memory().Goto(pov...)
		sh.focus().Path = pov
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
		x := sh.memory().Lookup(Address{pov})
		if p, ok := x.(interface{ Peek() Circuit }); ok {
			if _, ok := sh.foci[w[1]]; ok {
				fmt.Fprintf(sh.err, "focus name %q already in use", w[1])
				return
			}
			sh.attach(w[1], p.Peek())
		} else {
			fmt.Fprintf(sh.err, "object of type %T does not have a peek method", x)
		}
	default:
		fmt.Fprintf(sh.err, "peek accepts two arguments: a path and a new unique focus name\n")
	}
}

func (sh *Shell) at() Memory {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "current path is disconnected from root\n")
		}
	}()
	return sh.memory().Goto(sh.focus().Path...)
}

func (sh *Shell) ls(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "path not found\n")
		}
	}()
	switch {
	case len(w) == 0:
		fmt.Fprintf(sh.err, "%v\n", Circuit(sh.at()).Print("", "   ", 1))
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
	fmt.Fprintf(sh.err, "%s\n", printPath(sh.focus().Path))
}

func (sh *Shell) unlink(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "unlink error\n")
		}
	}()
	sh.at().Unlink(parseLink(w))
}

func (sh *Shell) link(w []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(sh.err, "link error\n")
		}
	}()
	sh.at().Link(parseLink(w))
}

func (sh *Shell) mk(w []string) {
	switch {
	case len(w) == 0:
		fmt.Fprintf(sh.err, "mk command needs arguments\n")
	case len(w) == 1:
		for _, b := range w[0] {
			if !see.IsIdentifier(rune(b)) {
				fmt.Fprintf(sh.err, "name must be an identifier\n")
				return
			}
		}
		dir, file := path.Split(w[0])
		if dir != "" {
			fmt.Fprintf(sh.err, "mk argument cannot be a path\n")
			return
		}
		sh.at().Refine(file)
	case len(w) == 2:
		dir, file := path.Split(w[0])
		if dir != "" {
			fmt.Fprintf(sh.err, "mk argument cannot be a path\n")
			return
		}
		x := see.SeeValueOrNil(see.NewSrcString(w[1]))
		if x == nil {
			fmt.Fprintf(sh.err, "Value not recognized\n")
			return
		}
		sh.at().Include(file, x)
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
r             Recycle the current view
p             Show current path
v             Show all foci
j b           Change current focus to "b"
ls            Show circuit in current focus
ls ../ef/     Show circuit at path relative to current
peek x/y  Peek addressed object if peeking supported
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
