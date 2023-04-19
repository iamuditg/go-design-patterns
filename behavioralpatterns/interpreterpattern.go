package behavioralpatterns

import (
	"fmt"
	"strings"
)

/*
The Interpreter Pattern is a behavioral design pattern
that allows you to define a grammar for a language and implement an interpreter that interprets sentences in that language.
*/

// Expression interface
type Expression interface {
	Interpret() bool
}

// Terminal expression
type TerminalExpression struct {
	data string
}

func (t *TerminalExpression) Interpret() bool {
	return strings.Contains(t.data, "go")
}

// Or expression
type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) Interpret() bool {
	return o.expr1.Interpret() || o.expr2.Interpret()
}

// And expression
type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) Interpret() bool {
	return a.expr1.Interpret() && a.expr2.Interpret()
}

// Interpreter
type Interpreter struct {
	expression Expression
}

func (i *Interpreter) Interpret() bool {
	return i.expression.Interpret()
}

func main() {
	// Creating expressions
	expression1 := &TerminalExpression{"golang"}
	expression2 := &TerminalExpression{"java"}
	expression3 := &TerminalExpression{"python"}

	// Creating a parse tree
	orExpression := &OrExpression{expr1: expression1, expr2: expression2}
	andExpression := &AndExpression{expr1: orExpression, expr2: expression3}

	// Creating interpreter
	interpreter := &Interpreter{expression: andExpression}

	// Interpreting
	fmt.Printf("Interpreter interprets '%s' as '%t'", "golang and python or java", interpreter.Interpret())
}
