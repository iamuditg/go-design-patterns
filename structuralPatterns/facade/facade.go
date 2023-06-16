package main

import "fmt"

// SubsystemA represents a subsystem of complex functionality.
type SubsystemA struct{}

// OperationA1 represents a specific operation of SubsystemA.
func (s *SubsystemA) OperationA1() string {
	return "SubsystemA: Operation A1\n"
}

// OperationA2 represents another operation of SubsystemA.
func (s *SubsystemA) OperationA2() string {
	return "SubsystemA: Operation A2\n"
}

// SubsystemB represents another subsystem of complex functionality.
type SubsystemB struct{}

// OperationB1 represents a specific operation of SubsystemB.
func (s *SubsystemB) OperationB1() string {
	return "SubsystemB: Operation B1\n"
}

// OperationB2 represents another operation of SubsystemB.
func (s *SubsystemB) OperationB2() string {
	return "SubsystemB: Operation B2\n"
}

// Facade represents the facade that simplifies the interface of the subsystems.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

// NewFacade creates a new instance of the Facade.
func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

// Operation wraps the complex operations of the subsystems and provides a simplified interface.
func (f *Facade) Operation() string {
	result := "Facade operation:\n"
	result += f.subsystemA.OperationA1()
	result += f.subsystemB.OperationB1()
	result += f.subsystemA.OperationA2()
	result += f.subsystemB.OperationB2()
	return result
}

func main() {
	// Create a facade
	facade := NewFacade()

	// Call the simplified operation
	result := facade.Operation()

	// Print the result
	fmt.Println(result)

	// Output:
	// Facade operation:
	// SubsystemA: Operation A1
	// SubsystemB: Operation B1
	// SubsystemA: Operation A2
	// SubsystemB: Operation B2
}
