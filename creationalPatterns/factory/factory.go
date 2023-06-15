package main

import "fmt"

// Product represents the product created by the factory.
type Product interface {
	Use()
}

// ConcreteProductA represents a concrete implementation of the product.
type ConcreteProductA struct{}

// Use uses the product A.
func (p *ConcreteProductA) Use() {
	fmt.Println("Using Product A")
}

// ConcreteProductB represents a concrete implementation of the product.
type ConcreteProductB struct{}

// Use uses the product B.
func (p *ConcreteProductB) Use() {
	fmt.Println("Using Product B")
}

// Factory represents the factory interface.
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA represents a concrete implementation of the factory.
type ConcreteFactoryA struct{}

// CreateProduct creates a new instance of Product A.
func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB represents a concrete implementation of the factory.
type ConcreteFactoryB struct{}

// CreateProduct creates a new instance of Product B.
func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Create an instance of Factory A
	factoryA := &ConcreteFactoryA{}

	// Create a product using Factory A
	productA := factoryA.CreateProduct()
	productA.Use()

	// Create an instance of Factory B
	factoryB := &ConcreteFactoryB{}

	// Create a product using Factory B
	productB := factoryB.CreateProduct()
	productB.Use()

	// Output:
	// Using Product A
	// Using Product B
}
