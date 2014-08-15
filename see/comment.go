// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

func SpaceNoNewline(src *Src) {
	if len(Whitespace(src)) > 0 {
		return
	}
	panic("whitespace")
}

func Space(src *Src) (newLine bool) {
	for commentAndEndOfLine(src) {
		newLine = true
	}
	if src.Len() == 0 {
		newLine = true
	}
	return
}

func commentAndEndOfLine(src *Src) bool {
	Whitespace(src)
	if len(src.Consume(isCommaOrSemicolon)) > 0 {
		return true
	}
	comment(src)
	return len(src.Consume(isNewline)) > 0
}

func comment(src *Src) {
	defer func() {
		recover()
	}()
	t := src.Copy()
	t.Match("//")
	t.Consume(isNotNewline) // comment body
	src.Become(t)
}
