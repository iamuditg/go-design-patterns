package structuralPatterns

import "fmt"

/*
The Facade Pattern is a structural design pattern that provides a simplified interface to a complex system of classes,
interfaces, and objects. It encapsulates a group of interfaces and classes into a single interface,
making it easier to use and reducing the complexity of the underlying system.
*/

/*
In this example, we define two subsystems: Subsystem1 and Subsystem2,
which have their own separate interfaces and classes.
We define a Facade struct, which encapsulates both Subsystem1 and Subsystem2 into a single interface.
The Facade struct has a NewFacade() method that creates instances of both subsystems,
and an Operation() method that calls the Operation1() and Operation2() methods of both subsystems.
Finally, we demonstrate the use of the Facade Pattern by creating a Facade object and calling its Operation() method,
which initializes and uses both subsystems to perform some operation.
Overall, the Facade Pattern provides a simplified interface to a complex system of classes, interfaces, and objects,
making it easier to use and reducing the complexity of the underlying system.
It can be useful in scenarios where you have a complex system with many components,
and you want to provide a simple and unified interface for using that system.
*/

// Subsystem1

type Subsystem1 struct{}

func (s *Subsystem1) Operation1() string {
	return "Subsystem1 operation"
}

// Subsystem2

type Subsystem2 struct{}

func (s *Subsystem2) Operation2() string {
	return "Subsystem2 operation"
}

// Facade
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

func (f *Facade) Operation() string {
	result := "Facade initializes subsystems:\n"
	result += f.subsystem1.Operation1() + "\n"
	result += f.subsystem2.Operation2() + "\n"
	return result
}

func main() {
	// Using the facade
	facade := NewFacade()
	result := facade.Operation()
	fmt.Println(result)
}
