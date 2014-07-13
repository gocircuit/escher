// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package think

func DontCognize(interface{}) {}

func NewNounReflex(v interface{}) Reflex {
	s, t := NewSynapse()
	go func() {
		s.Attach(DontCognize).ReCognize(v)
	}()
	return Reflex{"": t}
}
