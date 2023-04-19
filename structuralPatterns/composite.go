package structuralPatterns

import "fmt"

/*
The Composite Pattern is a structural design pattern that allows you to compose objects into tree structures to represent part-whole hierarchies.
It allows you to treat individual objects and compositions of objects uniformly, allowing you to work with complex structures as if they were simple objects.
*/

/*
In this example, we define a Component interface, which provides a method Operation() that returns a string.
We define a leaf component: Leaf, which has a name attribute. The Leaf struct implements the Component interface,
and its Operation() method returns a string indicating that it is a leaf component.
We define a composite component: Composite, which has a children attribute of type []Component.
The Composite struct implements the Component interface, and its Operation() method returns a string indicating that it is a composite component,
along with the results of calling the Operation() method on all of its children.
Finally, we demonstrate the use of the Composite Pattern by creating a hierarchy of components,
including two leaf components and one composite component. We call the Operation()
method on both the composite component and the child composite component, and observe the output.
Overall, the Composite Pattern allows you to compose objects into tree structures to represent part-whole hierarchies,
and to treat individual objects and compositions of objects uniformly.
This can be useful in scenarios where you need to work with complex structures as if they were simple objects,
or when you need to represent part-whole hierarchies in your code.
*/

// Component interface

type Component interface {
	Operation() string
}

// Leaf Component

type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return "leaf" + l.name
}

// Composite component

type Composite struct {
	children []Component
}

func (c *Composite) Add(child Component) {
	c.children = append(c.children, child)
}

func (c *Composite) Remove(child Component) {
	for i, cmp := range c.children {
		if cmp == child {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}
	}
}

func (c *Composite) Operation() string {
	var result string
	for _, child := range c.children {
		result += child.Operation()
	}
	return "composite" + result
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
