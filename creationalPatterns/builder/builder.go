package main

import "fmt"

// Product represents the object being built.
type Product struct {
	part1 string
	part2 int
	part3 bool
}

// Builder represents the builder interface.
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() Product
}

// ConcreteBuilder represents a concrete builder that implements the builder interface.
type ConcreteBuilder struct {
	product Product
}

// BuildPart1 builds the first part of the product.
func (b *ConcreteBuilder) BuildPart1() {
	b.product.part1 = "Part 1"
}

// BuildPart2 builds the second part of the product.
func (b *ConcreteBuilder) BuildPart2() {
	b.product.part2 = 42
}

// BuildPart3 builds the third part of the product.
func (b *ConcreteBuilder) BuildPart3() {
	b.product.part3 = true
}

// GetProduct returns the built product.
func (b *ConcreteBuilder) GetProduct() Product {
	return b.product
}

// Director represents the director that controls the building process.
type Director struct {
	builder Builder
}

// SetBuilder sets the builder to use.
func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

// Construct builds the product using the builder.
func (d *Director) Construct() Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
	return d.builder.GetProduct()
}

func main() {
	// Create a director
	director := Director{}

	// Create a builder
	builder := &ConcreteBuilder{}

	// Set the builder for the director
	director.SetBuilder(builder)

	// Construct the product
	product := director.Construct()

	// Print the product
	fmt.Printf("Product: %+v\n", product)

	// Output:
	// Product: {part1:Part 1 part2:42 part3:true}
}
