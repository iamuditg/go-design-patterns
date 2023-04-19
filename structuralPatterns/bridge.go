package structuralPatterns

import "fmt"

/*
The Bridge Pattern is a structural design pattern that decouples an abstraction from its implementation,
allowing them to vary independently.
It allows you to create a bridge between the interface of an abstraction and the implementation of that interface,
enabling the two to vary independently.
*/

/*
In this example, we define an Implementor interface, which provides a method Operation() that returns a string.
We define two concrete implementors: ConcreteImplementorA and ConcreteImplementorB, which implement the Implementor interface and return different strings from their Operation() methods.
We define an Abstraction interface, which provides a method SetImplementor() to set the implementor, and a method Operation() to use the implementor to perform an operation.
We define a RefinedAbstraction struct, which has an implementor attribute of type Implementor. The RefinedAbstraction struct implements the Abstraction interface,
and its Operation() method uses the implementor attribute to perform an operation.
Finally, we demonstrate the use of the Bridge Pattern by creating a RefinedAbstraction object with ConcreteImplementorA as the implementor,
and then calling its Operation() method. We then create a new RefinedAbstraction object with ConcreteImplementorB as the implementor, and call its Operation() method.
Overall, the Bridge Pattern allows you to decouple an abstraction from its implementation, allowing them to vary independently.
This can be useful in scenarios where you want to switch implementations at runtime, or when you have multiple implementations that need to work with a common abstraction.
*/

// Implementor interface
type Implementor interface {
	Operation() string
}

// Concrete implementors
type ConcreteImplementorA struct{}

func (a *ConcreteImplementorA) Operation() string {
	return "concrete implementor A"
}

type ConcreteImplementorB struct{}

func (b *ConcreteImplementorB) Operation() string {
	return "concrete implementor B"
}

// Abstraction interface
type Abstraction interface {
	SetImplementor(Implementor)
	Operation() string
}

// Refined abstraction
type RefinedAbstraction struct {
	implementor Implementor
}

func (r *RefinedAbstraction) SetImplementor(implementor Implementor) {
	r.implementor = implementor
}

func (r *RefinedAbstraction) Operation() string {
	return "refined abstraction with " + r.implementor.Operation()
}

func main() {
	// Creating a refined abstraction with concrete implementor A
	implementorA := &ConcreteImplementorA{}
	refinedAbstraction := &RefinedAbstraction{}
	refinedAbstraction.SetImplementor(implementorA)

	// Using the refined abstraction with concrete implementor A
	resultA := refinedAbstraction.Operation()
	fmt.Println(resultA)

	// Creating a refined abstraction with concrete implementor B
	implementorB := &ConcreteImplementorB{}
	refinedAbstraction.SetImplementor(implementorB)

	// Using the refined abstraction with concrete implementor B
	resultB := refinedAbstraction.Operation()
	fmt.Println(resultB)
}
