package main

import (
	"github.com/LeoCBS/scrapy-dsl/ast"
	"github.com/LeoCBS/scrapy-dsl/lexer"
	"github.com/LeoCBS/scrapy-dsl/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  toke.Token
	peekToken token.Token
}

func New() *Parser {

	p := &Parser{l: l}

	//read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
}
