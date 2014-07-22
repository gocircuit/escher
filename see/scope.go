// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

// “[v, u, w, x]”
func Scope(src *Src) (scope []interface{}, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			scope, ok = nil, false
		}
	}()
	t := src.Copy()
	t.Match("[")
	Space(t)
	for {
		q := t.Copy()
		Space(q)
		id, ok := SeeDesign(q)
		if !ok {
			break
		}
		Space(q)
		q.TryMatch(",")
		Space(q)
		scope = append(scope, id)
		t.Become(q)
	}
	Space(t)
	t.TryMatch(",")
	Space(t)
	t.Match("]")
	Space(t)
	src.Become(t)
	return scope, true
}
