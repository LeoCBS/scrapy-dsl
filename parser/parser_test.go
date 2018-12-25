package parser_test

import (
	"testing"

	//	"github.com/LeoCBS/scrapy-dsl/ast"
	"github.com/LeoCBS/scrapy-dsl/lexer"
	"github.com/LeoCBS/scrapy-dsl/parser"
)

func TestLetStatements(t *testing.T) {

	input := "let leo = 10;"
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Error("parserProgram() return nil")
	}
}
