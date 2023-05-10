package creationalPatterns

/*
The Abstract Factory Pattern is a creational design pattern that allows you to create families of related objects without specifying their concrete classes.
It provides an interface for creating families of related objects, but allows subclasses to decide which concrete classes to instantiate.
In this way, the Abstract Factory Pattern decouples the client code from the implementation details of the objects being created.
*/

// Product Interface
type Product interface {
	Use() string
}

// ConcreteProduct
type ConcreteProduct struct {
}

// Use returns a string indicating how concreteProduct is used
func (p *ConcreteProduct) Use() string {
	return "Using ConcreteProduct"
}

// Factory is an interface that defines the behaviour of a factory
type Factory interface {
	Create() Product
}

// ConcreteFactory is a factory that creates ConcreteProduct
type ConcreteFactory struct {
}

// Create returns a new instance of ConcreteProduct
func (f *ConcreteFactory) Create() Product {
	return &ConcreteProduct{}
}
