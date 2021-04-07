package ast

import (
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{},
	}
	expected := ""
	if program.String() != expected {
		t.Errorf("program.String() wrong. want=%q, got=%q", expected, program.String())
	}
}
