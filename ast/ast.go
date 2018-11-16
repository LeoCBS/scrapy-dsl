// Package responsible to define abstract syntax tree (AST)
// SDSL choose model top-down to parser AST using strategy recursive descent parser
//
// Statements examples:
//
// let lastPage = 10;
//
// let <identifier> = <expression>;
//
// PS: Basically the difference between idetifier and expression is that
// expression produce some value and identifier don't

package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Node interface implementation, this is the roof element from us AST,
// Our AST always start with Program Node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

// represent let statement, implements func statementNode and TokenLiteral as
// like any statement node
type LetStatement struct {
	Token token.Token // the token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// This struct represents variable names (identifier).
// Identifier is expression type just to make possible
// assign identifier to another identifier and produce values,
//e.g:
// let count = length
type Identifier struct {
	Token token.Token // the token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
