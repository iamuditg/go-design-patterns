package structuralPatterns

import "fmt"

/*
The Decorator Pattern is a structural design pattern that allows behavior to be added to an individual object,
either statically or dynamically, without affecting the behavior of other objects from the same class.
It is useful when you want to add new functionality to an object without modifying its original code.
*/

/*
In this example, we define a Component interface, which provides a method Operation() that returns a string.
We define a concrete component: ConcreteComponent, which implements the Component interface and returns a string indicating that it is a concrete component.
We define a Decorator struct, which has a component attribute of type Component.
The Decorator struct implements the Component interface, and its Operation() method calls the Operation()
method of the component attribute, while adding some additional behavior.
We define a concrete decorator: ConcreteDecorator, which embeds the Decorator struct and adds some additional behavior to the Operation() method.
Finally, we demonstrate the use of the Decorator Pattern by creating a ConcreteComponent object,
and then creating a Decorator object and a ConcreteDecorator object that wrap the ConcreteComponent.
We then call the Operation() method on the ConcreteDecorator object, which internally calls the Operation()
method of the ConcreteComponent object, and observe the output.
Overall, the Decorator Pattern allows you to add behavior to an individual object,
either statically or dynamically, without affecting the behavior of other objects from the same class.
This can be useful in scenarios where you want to add new functionality to an object without modifying its original code,
or when you want to add or remove functionality at runtime.
*/

// ComponentP interface
type ComponentP interface {
	Operation() string
}

// Concrete Component struct

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "concrete component"
}

// Decorator struct
type Decorator struct {
	component ComponentP
}

func (d *Decorator) Operation() string {
	return "decorator " + d.component.Operation()
}

// Concrete decorator
type ConcreteDecorator struct {
	decorator Decorator
}

func (d *ConcreteDecorator) Operation() string {
	return "concrete decorator " + d.decorator.Operation()
}

func main() {
	// Creating a concrete component
	component := &ConcreteComponent{}

	// Creating a decorator for the concrete component
	decorator := &Decorator{component: component}

	// Creating a concrete decorator for the decorator
	concreteDecorator := &ConcreteDecorator{decorator: *decorator}

	// Using the concrete decorator
	result := concreteDecorator.Operation()
	fmt.Println(result)
}
