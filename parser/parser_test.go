package parser

import (
	"go-interpreter/ast"
	"go-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	// create lexer from input
	l := lexer.New(input)
	// create parser from lexer of input
	p := New(l)

	program := p.ParseProgram()
	if program != nil {
		t.Fatalf("ParseProgram() returns nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expextedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expextedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not let, got=%q", s.TokenLiteral())
		return false
	}

	// s.(*ast.Statement) returns a tuple, which consists of the
	// statement, and an "ok" variable that tells us if the value
	// s matches the struct ast.Statement
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not type ast.Statement, got=%T", s)
		return false
	}

	//
	if letStmt.Name.Value != name {
		t.Errorf("let statement name value not '%s', got '%s'", name, letStmt.Name.Value)
		return false
	}
	return true
}
