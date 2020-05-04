package ast

// AST is an interface of an
// Abstract syntax tree for the this Shell
type AST interface{}

// TokenMeta contains meta information about
// a Token that's been parsed into the AST
type TokenMeta struct {
	// LineNumber holds the number
	// that the token was parsed from
	LineNumber int

	// CharacterCount holds the number
	// of character from the start of the current
	// line that this token started at
	CharacterCount int
}

// Token represents an individual token
// including its string value and meta
// information but not it's actual grammer token
type Token struct {

	// Value holds the actual string
	// value of the parsed token
	Value string

	// Meta holds the meta values
	// associated with this token
	Meta TokenMeta
}
