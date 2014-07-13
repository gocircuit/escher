// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

// “min: [v, u, w, x],  // comment”
func SeeField(src *Src) (name string, scope []interface{}, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			name, scope, ok = "", nil, false
		}
	}()
	t := src.Copy()
	Space(t)
	if name = Identifier(t); name == "" { // Name
		return "", nil, false
	}
	Space(t)
	t.Match(":")
	Space(t)
	scope, ok = Scope(t)
	if !ok {
		return "", nil, false
	}
	Space(t)
	t.TryMatch(",")
	Space(t)
	src.Become(t)
	return name, scope, true
}
