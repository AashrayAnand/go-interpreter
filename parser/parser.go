package parser

import (
	"fmt"
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"go-interpreter/token"
	"strconv"
)

const (
	_ int = iota
	// iota gives the following constants incrementing
	// numbers from zero, this is used to establish a
	// precedence of operators, when evaluating expressions
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string

	// used to check if there is an associated function with
	// an operator, either in the prefix map, or infix map
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	// parsing function for prefix operators
	prefixParseFn func() ast.Expression
	// parsing function for infix operators
	// argument is left side of infix operator e.g.
	// for 5 + 9, 5 is ast.Expression argument
	infixParseFn func(ast.Expression) ast.Expression
)

// helper functions for parser to add entries to prefix and infix
// parser function maps
func (p *Parser) addPrefixParseFn(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) addInfixParseFn(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// advance current and peeking token, pointing curToken
// at peekToken, and calling NextToken on our lexer to
// get the peek token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// if next token is expected type, get next token, otherwise
// return false
func (p *Parser) expectPeek(t token.TokenType) bool {
	// get next token if it matches expected
	// token type, return true, else return false
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// check if current token is expected type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// check if next token is expected type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type)
	// add new syntax error to list of errors
	p.errors = append(p.errors, msg)

}

// parse a let statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// p.curToken will be a let token
	stmt := &ast.LetStatement{Token: p.curToken}

	// return nil if let is not followed by an
	// identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// add identifier as name of the let statement node
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// return nil if identifier is not followed by equals sign
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// iterate through tokens until semicolon is reached
	// TODO: evaluate right hand expression
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	// get prefix parsing function for current type
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	// create ExpressionStatement with current token
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	// set the Expression to be the left side of the expression
	stmt.Expression = p.parseExpression(LOWEST)

	// go to next token if we reach a semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// create identifier node for current token (token.IDENT)
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	il := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	il.Value = value
	return il
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	// case of current token being beginning
	// of a let statement
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func New(l *lexer.Lexer) Parser {
	p := Parser{
		l:      l,
		errors: []string{},
	}

	// make prefix parse functions map
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	// bind identifier parse function to identifier token
	p.addPrefixParseFn(token.IDENT, p.parseIdentifier)
	p.addPrefixParseFn(token.INT, p.parseIntegerLiteral)
	// get first two tokens, place them into curr
	// and peek token positions
	p.nextToken()
	p.nextToken()

	return p
}
