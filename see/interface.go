// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

// “(v, u, w, x)”
func SeeInterface(src *Src) (valves []string) {
	t := src.Copy()
	t.Match("(")
	Space(t)
	for {
		q := t.Copy()
		Space(q)
		id := Identifier(q)
		if id == "" {
			break
		}
		Space(q)
		q.TryMatch(",")
		Space(q)
		valves = append(valves, id)
		t.Become(q)
	}
	Space(t)
	t.TryMatch(",")
	Space(t)
	t.Match(")")
	Space(t)
	src.Become(t)
	return
}
