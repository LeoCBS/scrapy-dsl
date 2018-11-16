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
	ASSIGN = "="
	//PLUS   = "+"

	// Delimiters
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	// Keywords
	LET     = "LET"
	XPATH   = "XPATH"
	GET     = "GET"
	EXTRACT = "EXTRACT"
	ON      = "ON"
)

var keywords = map[string]TokenType{
	"let":     LET,
	"xpath":   XPATH,
	"get":     GET,
	"extract": EXTRACT,
	"on":      ON,
}

func LookupType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
