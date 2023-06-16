package main

import "fmt"

// Implementor defines the interface for the implementors.
type Implementor interface {
	OperationImpl() string
}

// ConcreteImplementorA represents a concrete implementation of the Implementor interface.
type ConcreteImplementorA struct{}

// OperationImpl implements the OperationImpl method of the Implementor interface for ConcreteImplementorA.
func (imp *ConcreteImplementorA) OperationImpl() string {
	return "ConcreteImplementorA operation"
}

// ConcreteImplementorB represents another concrete implementation of the Implementor interface.
type ConcreteImplementorB struct{}

// OperationImpl implements the OperationImpl method of the Implementor interface for ConcreteImplementorB.
func (imp *ConcreteImplementorB) OperationImpl() string {
	return "ConcreteImplementorB operation"
}

// Abstraction defines the abstraction interface.
type Abstraction interface {
	Operation() string
}

// RefinedAbstraction represents the refined abstraction that extends the abstraction.
type RefinedAbstraction struct {
	implementor Implementor
}

// Operation delegates the operation to the implementor.
func (ra *RefinedAbstraction) Operation() string {
	return ra.implementor.OperationImpl()
}

func main() {
	// Create instances of the implementors
	impA := &ConcreteImplementorA{}
	impB := &ConcreteImplementorB{}

	// Create instances of the refined abstractions, each with a different implementor
	abstractionA := &RefinedAbstraction{implementor: impA}
	abstractionB := &RefinedAbstraction{implementor: impB}

	// Call the operations on the abstractions
	resultA := abstractionA.Operation()
	resultB := abstractionB.Operation()

	// Print the results
	fmt.Println(resultA)
	fmt.Println(resultB)

	// Output:
	// ConcreteImplementorA operation
	// ConcreteImplementorB operation
}
