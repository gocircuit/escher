// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so symbols of other times and backgrounds can
// see history clearly.

package circuit

// Name is one of: int or string
type Name interface{}

// Meaning is one of: string, int, float64, complex128, Circuit
type Meaning interface{}

// Super is a placeholder meaning for the super symbol
type Super struct{}

func (Super) String() string {
	return "*"
}

// circuit ...
type circuit struct {
	symbol map[Name]Meaning
	match map[Name]map[Name]Matching // symbol -> valve -> opposing symbol and valve
}

type Circuit struct {
	*circuit
}

// Matching ...
type Matching struct {
	Symbol [2]Name
	Valve [2]Name
}

func (x Matching) Reverse() Matching {
	x.Symbol[0], x.Symbol[1] = x.Symbol[1], x.Symbol[0]
	x.Valve[0], x.Valve[1] = x.Valve[1], x.Valve[0]
	return x
}

// New ...
func New() Circuit {
	return Circuit{
		&circuit{
			symbol: make(map[Name]Meaning),
			match: make(map[Name]map[Name]Matching),
		},
	}
}

var Nil Circuit // the nil circuit

func (u *circuit) IsNil() bool {
	return u == nil
}

func (u *circuit) IsEmpty() bool {
	return len(u.symbol) == 0 && len(u.match) == 0
}

// Change adds a symbol to this circuit.
func (u *circuit) Change(name Name, meaning Meaning) (before Meaning, overwrite bool) {
	before, overwrite = u.symbol[name]
	u.symbol[name] = meaning
	return
}

func (u *circuit) ChangeExclusive(name Name, meaning Meaning) {
	if _, over := u.Change(name, meaning); over {
		panic(1)
	}
}

// At ...
func (c *circuit) At(name Name) (Meaning, bool) {
	v, ok := c.symbol[name]
	return v, ok
}

func (u *circuit) Forget(name Name) Meaning {
	forgotten := u.symbol[name]
	delete(u.symbol, name)
	return forgotten
}

// Match ...
func (c *circuit) Match(x Matching) {
	if x.Symbol[0] == x.Symbol[1] && x.Valve[0] == x.Valve[1] {
		panic("mismatch")
	}
	p := []map[Name]Matching{
		c.valves(x.Symbol[0]), 
		c.valves(x.Symbol[1]),
	}
	v := x.Valve
	if _, ok := p[0][v[0]]; ok {
		panic("dup")
	}
	if _, ok := p[1][v[1]]; ok {
		panic("dup")
	}
	p[0][v[0]], p[1][v[1]] = x, x.Reverse()
}

func (c *circuit) valves(p Name) map[Name]Matching {
	if _, ok := c.symbol[p]; !ok {
		c.symbol[p] = nil // placeholder
	}
	if _, ok := c.match[p]; !ok {
		c.match[p] = make(map[Name]Matching)
	}
	return c.match[p]
}

func (u *circuit) Valves(symbol Name) map[Name]Matching {
	return u.match[symbol]
}

// Follow ...
func (c *circuit) Follow(p, v Name) (q, u Name) {
	x, ok := c.valves(p)[v]
	if !ok {
		return nil, nil
	}
	return x.Symbol[1], x.Valve[1]
}

func (c *circuit) Letters() []string {
	var l []string
	for key, _ := range c.symbol {
		if s, ok := key.(string); ok {
			l = append(l, s)
		}
	}
	return l
}

func (c *circuit) Numbers() []int {
	var l []int
	for key, _ := range c.symbol {
		if i, ok := key.(int); ok {
			l = append(l, i)
		}
	}
	return l
}

func (u *circuit) Symbols() map[Name]Meaning {
	return u.symbol
}

func (u *circuit) String() string {
	return u.Print(nil, "", "\t")
}
