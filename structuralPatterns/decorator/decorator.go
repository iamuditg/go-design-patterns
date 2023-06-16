package main

import "fmt"

// Component defines the interface for the component.
type Component interface {
	Operation() string
}

// ConcreteComponent represents the concrete component.
type ConcreteComponent struct{}

// Operation implements the Operation method of the Component interface.
func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent operation"
}

// Decorator represents the base decorator struct.
type Decorator struct {
	component Component
}

// Operation implements the Operation method of the Component interface.
// It delegates the operation to the underlying component.
func (d *Decorator) Operation() string {
	return d.component.Operation()
}

// ConcreteDecoratorA represents a concrete decorator.
type ConcreteDecoratorA struct {
	Decorator
}

// Operation decorates the operation of the component.
func (d *ConcreteDecoratorA) Operation() string {
	// Perform additional operations before calling the component's operation
	result := "ConcreteDecoratorA operation\n"
	result += d.component.Operation()
	// Perform additional operations after calling the component's operation
	return result
}

// ConcreteDecoratorB represents another concrete decorator.
type ConcreteDecoratorB struct {
	Decorator
}

// Operation decorates the operation of the component.
func (d *ConcreteDecoratorB) Operation() string {
	// Perform additional operations before calling the component's operation
	result := "ConcreteDecoratorB operation\n"
	result += d.component.Operation()
	// Perform additional operations after calling the component's operation
	return result
}

func main() {
	// Create a concrete component
	component := &ConcreteComponent{}

	// Create decorators and wrap the component
	decoratorA := &ConcreteDecoratorA{Decorator: Decorator{component: component}}
	decoratorB := &ConcreteDecoratorB{Decorator: Decorator{component: decoratorA}}

	// Call the operation on the decorator chain
	result := decoratorB.Operation()

	// Print the result
	fmt.Println(result)

	// Output:
	// ConcreteDecoratorB operation
	// ConcreteDecoratorA operation
	// ConcreteComponent operation
}
