package main

import "fmt"

// AbstractClass represents the abstract class with a template method.
type AbstractClass interface {
	TemplateMethod()
}

// ConcreteClassA is a concrete implementation of the AbstractClass.
type ConcreteClassA struct {
}

// TemplateMethod implements the template method of the AbstractClass for ConcreteClassA.
func (c *ConcreteClassA) TemplateMethod() {
	fmt.Println("Executing steps for ConcreteClassA:")
	c.Step1()
	c.Step2()
}

// Step1 is a step implementation for ConcreteClassA.
func (c *ConcreteClassA) Step1() {
	fmt.Println("Step 1 for ConcreteClassA")
}

// Step2 is a step implementation for ConcreteClassA.
func (c *ConcreteClassA) Step2() {
	fmt.Println("Step 2 for ConcreteClassA")
}

// ConcreteClassB is a concrete implementation of the AbstractClass.
type ConcreteClassB struct{}

// TemplateMethod implements the template method of the AbstractClass for ConcreteClassB.
func (c *ConcreteClassB) TemplateMethod() {
	fmt.Println("Executing steps for ConcreteClassB:")
	c.Step1()
	c.Step3()
}

// Step1 is a step implementation for ConcreteClassB.
func (c *ConcreteClassB) Step1() {
	fmt.Println("Step 1 for ConcreteClassB")
}

// Step3 is a step implementation for ConcreteClassB.
func (c *ConcreteClassB) Step3() {
	fmt.Println("Step 3 for ConcreteClassB")
}

func main() {
	// Create an instance of ConcreteClassA
	classA := &ConcreteClassA{}
	// Execute the template method of ConcreteClassA
	classA.TemplateMethod()

	fmt.Println()

	// Create an instance of ConcreteClassB
	classB := &ConcreteClassB{}
	// Execute the template method of ConcreteClassB
	classB.TemplateMethod()

	// Output:
	// Executing steps for ConcreteClassA:
	// Step 1 for ConcreteClassA
	// Step 2 for ConcreteClassA
	//
	// Executing steps for ConcreteClassB:
	// Step 1 for ConcreteClassB
	// Step 3 for ConcreteClassB
}
