package main

import "fmt"

// Cloneable represents the interface for cloning objects.
type Cloneable interface {
	Clone() Cloneable
}

// ConcretePrototype represents a concrete implementation of the prototype.
type ConcretePrototype struct {
	data string
}

// Clone creates a new instance of the concrete prototype by performing a deep copy.
func (p *ConcretePrototype) Clone() Cloneable {
	// Perform a deep copy of the data field
	newData := p.data
	return &ConcretePrototype{data: newData}
}

func main() {
	// Create an instance of the prototype
	prototype := &ConcretePrototype{data: "Prototype Data"}

	// Clone the prototype
	clone := prototype.Clone().(*ConcretePrototype)

	// Modify the cloned object
	clone.data = "Modified Data"

	// Print the original and cloned objects
	fmt.Println("Original Prototype:", prototype.data)
	fmt.Println("Cloned Object:", clone.data)

	// Output:
	// Original Prototype: Prototype Data
	// Cloned Object: Modified Data
}
