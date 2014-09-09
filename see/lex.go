// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

func isNotVerbatim(r rune) bool {
	return r != '`'
}

func isCommaOrSemicolon(r rune) bool {
	return r == '\n' || r == '\r' || r == ';' || r == ','
}

func isNewline(r rune) bool {
	return r == '\n' || r == '\r'
}

func isNotNewline(r rune) bool {
	return !isNewline(r)
}

func isIdentifier(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
		return true
	case r >= '0' && r <= '9':
		return true
	case r == '_', r == '?':
		return true
	}
	return false
}

func isIdentifierFirst(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
		return true
	case r == '_', r == '?':
		return true
	}
	return false
}

func Identifier(src *Src) string {
	if src.Len() == 0 {
		return ""
	}
	if !isIdentifierFirst(src.RuneAt(0)) {
		return ""
	}
	return src.Consume(isIdentifier)
}

// Identifier + Delimiter + Operator < Literal

func isOperator(r rune) bool {
	switch r {
	case '=':
		return true
	}
	return false
}

const ValveSelector = ':'

// isLiteral returns true iff r is not a whitespace and not a newline and not a comment character.
func isLiteral(r rune) bool {
	return r != ValveSelector && !isWhitespace(r) && !isNewline(r) && r != '/' && !isDelimiter(r)
}

func isDelimiter(r rune) bool {
	switch r {
	case '{', '}', '[', ']', '(', ')', ',', ';':
		return true
	}
	return false
}

func Literal(src *Src) string {
	return src.Consume(isLiteral)
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func Whitespace(src *Src) string {
	return src.Consume(isWhitespace)
}

func Newline(src *Src) int {
	return len(src.Consume(isNewline))
}

func Keyword(keyword string, src *Src) {
	src.Form(keyword)
	// make sure this is the end of the keyword
	if src.Len() > 0 && isIdentifier(src.RuneAt(0)) {
		panic("no keyword")
	}
	Whitespace(src)
}
