// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package see

// “reflex Name(A, B, C) // comment”
func SeeReflex(src *Src) (reflex *Reflex) {
	reflex = &Reflex{}
	t := src.Copy()
	Space(t)
	Keyword("reflex", t)
	Space(t)
	if reflex.Name = Identifier(t); reflex.Name == "" {
		return nil
	}
	Space(t)
	reflex.Valve = SeeInterface(t)
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return
}
