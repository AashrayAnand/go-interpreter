package ast

import (
	"go-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	// creates a program with a single statement, that is the
	// equivalent of let myVar = anotherVar;
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program string wrong, got... %s", program.String())
	}
}
