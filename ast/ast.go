package ast

import "go-interpreter/token"

// LetStatement : represents a let statement, comprised of
// a token, the identifier that is binded to, and the right
// side expression that is binded to the identifier
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier // name of the token, is pointer to an identifer
	Value Expression  // right side of the let statement, can be any expression
}

type ReturnStatement struct {
	Token token.Token // token.RETURN token
	Value Expression  // expression that is being returned
}

// dummy methods which will result in these structs
// implementing the Statement interface
func (ls *LetStatement) statementNode()    {}
func (ls *ReturnStatement) statementNode() {}

// TokenLiteral : get literal value of statement from token
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// TokenLiteral : get literal value of statement from token
func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier : represents an identifier
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// dummy method that will result in an identifier
// implementing the Expression interace
func (i *Identifier) expressionNode() {}

// TokenLiteral : get literal value of an identifier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Node : implementers of the Node interface must
// include a TokenLiteral method that
// returns the literal value of a token
type Node interface {
	TokenLiteral() string
}

// Statements do not return a value
// must implement the Node interface
type Statement interface {
	Node
	statementNode()
}

// Expressions return a value
// must implement the Node interface
type Expression interface {
	Node
	expressionNode()
}

// Program node will be the root of every AST produced
// by the parser, every valid monkey program is a series
// of statements. This is just a slice of AST nodes that
// implement the Statement interface
type Program struct {
	Statements []Statement
}

// TokenLiteral : method for Program node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
