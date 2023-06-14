package main

import "fmt"

// Strategy represents the strategy interface that defines the method(s) to
// be implemented by concrete strategies.
type Strategy interface {
	Execute()
}

// ConcreteStrategyA is a concrete implementation of the strategy interface.
type ConcreteStrategyA struct {
}

// Execute implements the execute method of the strategy interface for ConcreteStrategyA.
func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing ConcreteStrategyA")
}

// ConcreteStrategyB is a concrete implementation of the strategy interface.
type ConcreteStrategyB struct {
}

// Execute implements the execute method of the strategy interface for ConcreteStrategyA.
func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing ConcreteStrategyB")
}

// Context represents the context that uses the strategy.
type Context struct {
	strategy Strategy
}

// SetStrategy sets the strategy to be used by the context
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy executes the current strategy.
func (c *Context) ExecuteStrategy() {
	if c.strategy == nil {
		fmt.Println("No strategy set")
		return
	}
	c.strategy.Execute()
}

func main() {
	// Create a context
	context := &Context{}

	// Create concrete strategies
	strategyA := &ConcreteStrategyA{}
	strategyB := &ConcreteStrategyB{}

	// Set ConcreteStrategyA as the strategy
	context.SetStrategy(strategyA)
	// Execute the current strategy
	context.ExecuteStrategy()

	// Set ConcreteStrategyB as the strategy
	context.SetStrategy(strategyB)
	// Execute the current strategy
	context.ExecuteStrategy()

	// Output:
	// Executing ConcreteStrategyA
	// Executing ConcreteStrategyB
}
