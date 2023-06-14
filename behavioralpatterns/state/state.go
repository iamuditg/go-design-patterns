package main

import "fmt"

// State represents the state interface.
type State interface {
	Handle(context *Context)
}

// ConcreteStateA is a concrete implementation of the State interface.
type ConcreteStateA struct{}

// Handle handles the state in ConcreteStateA.
func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("Handling state in ConcreteStateA")
	context.SetState(&ConcreteStateB{})
}

// ConcreteStateB is a concrete implementation of the State interface.
type ConcreteStateB struct{}

// Handle handles the state in ConcreteStateB.
func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Handling state in ConcreteStateB")
	context.SetState(&ConcreteStateA{})
}

// Context represents the context that uses the state.
type Context struct {
	state State
}

// SetState sets the current state of the context.
func (c *Context) SetState(state State) {
	c.state = state
}

// Request triggers the handling of the state by invoking the current state's Handle method.
func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	// Create the context
	context := &Context{}

	// Set initial state to ConcreteStateA
	context.SetState(&ConcreteStateA{})

	// Perform requests and observe state changes
	context.Request()
	context.Request()
	context.Request()

	// Output:
	// Handling state in ConcreteStateA
	// Handling state in ConcreteStateB
	// Handling state in ConcreteStateA
}
