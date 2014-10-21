// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package memory

import (
	// "container/list"

	. "github.com/gocircuit/escher/circuit"
)

// I see forward. I think back. I see that I think. I think that I see. Thinking and seeing are not apart.

type Memory Circuit

func (m Memory) Print(prefix, indent string) string {
	return "(Memory)"
}

func (m Memory) Get(addr Address) (v Value) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
		}
	}()
	v = Circuit(m)
	for _, name := range addr.Path {
		v = v.(Circuit).At(name)
	}
	return DeepCopy(v)
}

func (m Memory) Put(addr Address, value Value) Value {
	if len(addr.Path) == 0 {
		panic("no path")
	}
	x := m
	for i, g := range addr.Path {
		if i+1 == len(addr.Path) {
			break
		}
		x = x.Refine(g)
	}
	return x.Include(addr.Path[len(addr.Path)-1], DeepCopy(value))
}

func (m Memory) Forget(addr Address) Value {
	if len(addr.Path) == 0 {
		panic("no path")
	}
	x := Circuit(m)
	for i, g := range addr.Path {
		if i+1 == len(addr.Path) {
			break
		}
		u, ok := x.OptionAt(g)
		if !ok {
			return nil
		}
		x, ok = u.(Circuit)
		if !ok {
			return nil
		}
	}
	return x.Exclude(addr.Path[len(addr.Path)-1])
}

func (m Memory) Goto(gate ...Name) Memory {
	for _, name := range gate {
		m = Memory(Circuit(m).CircuitAt(name))
	}
	return m
}

func (m Memory) Refine(name Name) Memory {
	v, ok := Circuit(m).OptionAt(name)
	if ok {
		if w, ok := v.(Circuit); ok {
			return Memory(w)
		}
	}
	k := New()
	Circuit(m).Include(name, k)
	return Memory(k)
}

func (m Memory) Include(n Name, v Value) Value {
	return Circuit(m).Include(n, Copy(v))
}

func (m Memory) Exclude(n Name) Value {
	return Circuit(m).Exclude(n)
}

func (m Memory) Link(u, v Vector) {
	Circuit(m).Link(u, v)
}

func (m Memory) Unlink(u, v Vector) {
	Circuit(m).Unlink(u, v)
}

func (m Memory) Len() int {
	return Circuit(m).Len()
}
