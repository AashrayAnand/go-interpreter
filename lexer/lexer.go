package lexer

import (
	"go-interpreter/token"
)

// Lexer : lexer struct
type Lexer struct {
	input        string
	position     int // current position in input (points to char)
	readPosition int // current reading position in input (after current char)
	// this is used to peek to the next char, if needed e.g.
	// to see if a second '=' comes after an '=' char
	ch byte // current char under examination

	// position is what we just read, readPosition is what we will read next
}

// New : create new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken : get the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// read characters until we have reached a non-whitespace character
	l.skipWhiteSpace()

	// different cases for character we encounter
	switch l.ch {
	case '=':
		// check if next 2 tokens form equals operator
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.TIMES, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	// if it does not match any special character
	// check if it is a literal
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// lookup the identifier, return special types for
			// keywords (if,else..), otherwise IDENT
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	// read next character
	l.readChar()
	return tok
}

// newToken : create new token of type t, and with literal value v
func newToken(t token.TokenType, v byte) token.Token {
	return token.Token{Type: t, Literal: string(v)}
}

// returns whether a character is a letter, or underscore
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peekChar : used to peek one character ahead, for cases of 2 char
// patterns, e.g. ==, !=
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// readIdentifier : gets all of the characters of an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	// increment position pointer while the
	// character is still a letter
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhiteSpace : lexer method which skips whitespace
func (l *Lexer) skipWhiteSpace() {
	// keep reading characters until we skip through white space
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readChar : read next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
