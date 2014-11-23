// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package a

func isNotVerbatim(r rune) bool {
	return r != '`'
}

func IsCommaOrSemicolonOrNewline(r rune) bool {
	return isNewline(r) || r == ';' || r == ','
}

func isNewline(r rune) bool {
	return r == '\n' || r == '\r'
}

func isNotNewline(r rune) bool {
	return !isNewline(r)
}

const RefineSymbolString = "."
const RefineSymbolRune = '.'

func IsIdentifier(r rune) bool {
	if r == RefineSymbolRune {
		return false
	}
	switch {
	case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
		return true
	case r >= '0' && r <= '9':
		return true
	case r == '_', r == '?', r == '@', r == '*', r == '-', r == '+', r == '$':
		return true
	}
	return false
}

func IsName(s string) bool {
	return NewSrcString(s).Consume(IsIdentifier) == s
}

func IsIdentifierOrRefineSymbol(r rune) bool {
	return r == RefineSymbolRune || IsIdentifier(r)
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
	src.Match(keyword)
	// make sure this is the end of the keyword
	if src.Len() > 0 && IsIdentifier(src.RuneAt(0)) {
		panic("no keyword")
	}
	Whitespace(src)
}
