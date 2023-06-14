package main

import (
	"fmt"
	"strconv"
)

// Expression represents the expression interface.
type Expression interface {
	Interpret() int
}

// NumberExpression represents a number expression.
type NumberExpression struct {
	number int
}

// Interpret interprets the number expression by returning its value.
func (n *NumberExpression) Interpret() int {
	return n.number
}

// PlusExpression represents an addition expression.
type PlusExpression struct {
	left  Expression
	right Expression
}

// Interpret interprets the addition expression by evaluating the left and right expressions and returning their sum.
func (p *PlusExpression) Interpret() int {
	return p.left.Interpret() + p.right.Interpret()
}

// MinusExpression represents a subtraction expression.
type MinusExpression struct {
	left  Expression
	right Expression
}

// Interpret interprets the subtraction expression by evaluating the left and right expressions and returning their difference.
func (m *MinusExpression) Interpret() int {
	return m.left.Interpret() - m.right.Interpret()
}

// Parser represents the parser that parses and evaluates expressions.
type Parser struct {
	tokens []string
	index  int
}

// ParseExpression parses the expression and returns the corresponding expression object.
func (p *Parser) ParseExpression() Expression {
	token := p.getNextToken()

	switch token {
	case "+":
		return &PlusExpression{
			left:  p.ParseExpression(),
			right: p.ParseExpression(),
		}
	case "-":
		return &MinusExpression{
			left:  p.ParseExpression(),
			right: p.ParseExpression(),
		}
	default:
		number, _ := strconv.Atoi(token)
		return &NumberExpression{
			number: number,
		}
	}
}

// getNextToken returns the next token in the token list.
func (p *Parser) getNextToken() string {
	token := p.tokens[p.index]
	p.index++
	return token
}

func main() {
	// Create the parser
	parser := &Parser{
		tokens: []string{"5", "2", "1", "+", "-"},
		index:  0,
	}

	// Parse the expression
	expression := parser.ParseExpression()

	// Interpret the expression
	result := expression.Interpret()
	fmt.Println("Result:", result)

	// Output:
	// Result: 4
}
