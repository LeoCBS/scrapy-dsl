// Package responsible to define all tokens present on sdsl
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // variavel names
	INT   = "INT"   // 12345

	// Operators
	//ASSIGN = "="
	//PLUS   = "+"

	// Delimiters
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	// Keywords
	xpath   = "xpath"
	GET     = "get"
	EXTRACT = "extract"
	PARSE   = "parse"
	ON      = "on"
)
