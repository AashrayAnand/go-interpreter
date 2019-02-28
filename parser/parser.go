package parser

import (
	"fmt"
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	// case of current token being beginning
	// of a let statement
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
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

	// get first two tokens, place them into curr
	// and peek token positions
	p.nextToken()
	p.nextToken()

	return p
}
