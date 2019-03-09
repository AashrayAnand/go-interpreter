package ast

import (
	"bytes"
	"go-interpreter/token"
)

// LetStatement : represents a let statement, comprised of
// a token, the identifier that is binded to, and the right
// side expression that is binded to the identifier
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier // name of the token, is pointer to an identifer
	Value Expression  // right side of the let statement, can be any expression
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN token
	ReturnValue Expression  // expression that is being returned
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression  //
}

type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int64
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// dummy methods which will result in these structs
// implementing the Statement interface
func (ls *LetStatement) statementNode()        {}
func (rs *ReturnStatement) statementNode()     {}
func (es *ExpressionStatement) statementNode() {}

// dummy methods which will result in these structs
// implementing the statement interface
func (il *IntegerLiteral) expressionNode()   {}
func (pe *PrefixExpression) expressionNode() {}
func (ie *InfixExpression) expressionNode()  {}

// TokenLiteral functions to satisfy Node interface
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String functions to satisfy node interface

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	// returns a concatenation of the token literal
	// value ("let") + Name (identifier value e.g. "x") and
	// the value (expression to string) e.g. "let x = ..."
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	// returns a concatenation of token literal
	// value ("return") + ReturnValue (expression to string)
	// e.g. "return ..."
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// returns expression as a string value
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(ie.Operator)
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (i *Identifier) String() string { return i.Value }

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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
	String() string
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
