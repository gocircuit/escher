// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

func SeePeer(src *Src) (p *Peer) {
	defer func() {
		if r := recover(); r != nil {
			p = nil
		}
	}()
	t := src.Copy()
	Space(t)
	p = &Peer{}
	p.Name = Identifier(t)
	if p.Name == "" {
		return nil
	}
	SpaceNoNewline(t)
	var ok bool
	if p.Design, ok = SeeDesign(t); !ok {
		return nil
	}
	if !Space(t) { // require newline at end
		return nil
	}
	src.Become(t)
	return
}
