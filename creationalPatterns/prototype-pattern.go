package creationalPatterns

import "fmt"

/*
The Prototype Pattern is a creational design pattern that allows you to create new objects by cloning existing objects.
It allows for efficient object creation and reduces the need for subclassing.
*/

/*

In this example, we define a Shape interface, which provides methods to clone a shape and draw it. We then define a concrete product: Circle, which implements the Shape interface and has a radius attribute.

Next, we demonstrate the use of the Prototype Pattern by creating a Circle object with a radius of 2.5, and then cloning it to create another Circle object with a radius of 5.0. We can see that the two Circle objects have different memory addresses.

Finally, we demonstrate that the Draw() method of both Circle objects works correctly.

Overall, the Prototype Pattern allows us to create new objects by cloning existing objects, rather than creating new objects from scratch. This can improve the efficiency of object creation and reduce the need for subclassing.

*/

// Product interface

type Shape interface {
	Clone() Shape
	Draw()
}

// Concrete product

type Circle struct {
	Radius float64
}

func (c *Circle) Clone() Shape {
	return &Circle{Radius: c.Radius}
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a circle with radius %.2f\n", c.Radius)
}
