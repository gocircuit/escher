// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package circuit

func (u Circuit) Reflow(s, t Name) {
	if _, ok := u.Flow[t]; ok {
		panic("reflow overwrite")
	}
	fs, ok := u.Flow[s]
	if !ok {
		return
	}
	u.Flow[t] = fs
	delete(u.Flow, s)
	for vlv, vec := range fs {
		u.Flow[vec.Gate][vec.Valve] = Vector{t, vlv}
	}
}
