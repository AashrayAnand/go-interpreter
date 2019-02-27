package lexer

import (
	"fmt"
	"testing"

	"go-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "let five = 4;" +
		"let ten = 10;" +
		"let add = fn(x, y) {" +
		"x + y;" +
		"};" +
		"let result = add(five, ten);"

	tests := []struct {
		expextedType    token.TokenType
		expextedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		// get next token from character stream
		tok := l.NextToken()
		fmt.Println(tok.Literal)

		if tok.Type != tt.expextedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, actual=%q", i, tt.expextedType, tok.Type)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, actual=%q", i, tt.expextedLiteral, tok.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5; if else if else
	`

	tests := []struct {
		expextedType    token.TokenType
		expextedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		//
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.TIMES, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		// get next token from character stream
		tok := l.NextToken()
		fmt.Println(tok.Literal)

		if tok.Type != tt.expextedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, actual=%q", i, tt.expextedType, tok.Type)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, actual=%q", i, tt.expextedLiteral, tok.Literal)
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `if (5 != 10) {
		return true;
	} else if (5 == 10){
		return false;
	}
	`

	tests := []struct {
		expextedType    token.TokenType
		expextedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.NEQ, "!="},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}
	l := New(input)

	for i, tt := range tests {
		// get next token from character stream
		tok := l.NextToken()
		fmt.Println(tok.Literal)

		if tok.Type != tt.expextedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, actual=%q", i, tt.expextedType, tok.Type)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, actual=%q", i, tt.expextedLiteral, tok.Literal)
		}
	}
}
