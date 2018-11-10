package lexer_test

import (
	"testing"

	"github.com/LeoCBS/scrapy-dsl/token"
)

func TestNextToken(t *testing.T) {
	input := "();"
	tests := []struct {
		expectedType    toke.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lexer = New(input)
	for i, testCase := range tests {
		tok := l.NextToken()
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
