package lexer_test

import (
	"testing"

	"github.com/LeoCBS/scrapy-dsl/lexer"
	"github.com/LeoCBS/scrapy-dsl/token"
)

func TestNextTokenOnlyWithDelimiters(t *testing.T) {
	input := "();"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := lexer.New(input)
	for i, ts := range tests {
		tok := l.NextToken()
		if tok.Type != ts.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, ts.expectedType, tok.Type)
		}
		if tok.Literal != ts.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, ts.expectedLiteral, tok.Literal)
		}
	}
}
