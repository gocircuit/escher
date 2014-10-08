{	
	"name": "Escher",
	"scopeName": "source.escher",
	"fileTypes": ["escher"],
	"patterns": [
		{
			"begin": "//",
			"end": "\\z",
			"name": "comment.line.double-slash.escher"
		},
		{
			"begin": "\"",
			"beginCaptures": {
				"0": {
					"name": "punctuation.definition.string.begin.escher"
				}
			},
			"end": "\"",
			"endCaptures": {
				"0": {
					"name": "punctuation.definition.string.end.escher"
				}
			},
			"name": "string.quoted.double.escher",
			"patterns": [
				{
					"include": "#string_placeholder"
				},
				{
					"include": "#string_escaped_char"
				}
			]
		},
		{
			"begin": "`",
			"beginCaptures": {
				"0": {
					"name": "punctuation.definition.string.begin.escher"
				}
			},
			"end": "`",
			"endCaptures": {
				"0": {
					"name": "punctuation.definition.string.end.escher"
				}
			},
			"name": "string.quoted.raw.escher",
			"patterns": [
				{
					"include": "#string_placeholder"
				},
				{
					"include": "source.gotemplate"
				}
			]
		},
		{
			"match": "\\b((\\d+\\.(\\d+)?([eE][+-]?\\d+)?|\\d+[eE][+-]?\\d+|\\.\\d+([eE][+-]?\\d+)?)i?)\\b",
			"name": "constant.numeric.floating-point.escher"
		},
		{
			"match": "\\b(\\d+i|0[xX][0-9A-Fa-f]+|0[0-7]*|[1-9][0-9]*)\\b",
			"name": "constant.numeric.integer.escher"
		},
		{
			"name": "constant.other.rune.escher",
			"match": "'(?:[^'\\\\]|\\\\(?:\\\\|[abfnrtv']|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8}|[0-7]{3}))'"
		},
		{
			"captures": {
				"0": {
					"name": "variable.other.escher"
				},
				"1": {
					"name": "keyword.operator.initialize.escher"
				}
			},
			"comment": "This matches the 'x := 0' style of variable declaration.",
			"match": "(?:[[:alpha:]_][[:alnum:]_]*)(?:,\\s+[[:alpha:]_][[:alnum:]_]*)*\\s*(=)",
			"name": "meta.initialization.short.escher"
		},
		{
			"match": "(=|(?:[+]|-|[|]|^|[*]|/|%|<<|>>|&|&^)=)",
			"name": "keyword.operator.assignment.escher"
		},
		{
			"match": "(;)",
			"name": "keyword.operator.semi-colon.escher"
		},
		{
			"match": "(,)",
			"name": "punctuation.definition.comma.escher"
		},
	],
	"uuid": "ca03e751-04ef-4330-9a6b-9b99aae1c418"
}
