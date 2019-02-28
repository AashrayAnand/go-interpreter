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
	if program == nil {
		t.Fatalf("ParseProgram() returns nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if letStmt == nil {
		t.Errorf("nil let statement")
		return false
	}
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
				return x;
				return 10;
				return 5 + 4;
				`
	// create lexer from input
	l := lexer.New(input)
	// create parser from lexer of input
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returns nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"return"},
		{"return"},
		{"return"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testReturnStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "return" {
		t.Errorf("s.TokenLiteral not 'return'. got=%q", s.TokenLiteral())
		return false
	}
	return true
}
