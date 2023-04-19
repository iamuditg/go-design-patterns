package main

import (
	"fmt"
	"github.com/iamuditg/creationalPatterns"
	"github.com/iamuditg/structuralPatterns"
)

func main() {

	// abstractFactoryPattern()
	// builderPattern()
	// prototypePattern()
	// singletonPattern()
	bridgePattern()

}

func bridgePattern() {
	// Creating a refined abstraction with concrete implementor A
	implementorA := &structuralPatterns.ConcreteImplementorA{}
	refinedAbstraction := &structuralPatterns.RefinedAbstraction{}
	refinedAbstraction.SetImplementor(implementorA)

	// Using the refined abstraction with concrete implementor A
	resultA := refinedAbstraction.Operation()
	fmt.Println(resultA)

	// Creating a refined abstraction with concrete implementor B
	implementorB := &structuralPatterns.ConcreteImplementorB{}
	refinedAbstraction.SetImplementor(implementorB)

	// Using the refined abstraction with concrete implementor B
	resultB := refinedAbstraction.Operation()
	fmt.Println(resultB)
}

func singletonPattern() {
	// Getting the singleton instance
	singleton := creationalPatterns.GetInstance()

	// Setting the data of the singleton
	singleton.SetData("new data")

	// Getting the data of the singleton
	data := singleton.GetData()
	fmt.Println(data)
}

/*
PROTOTYPE PATTERN
*/
func prototypePattern() {
	// Creating a circle with radius 2.5
	circle1 := &creationalPatterns.Circle{Radius: 2.5}
	fmt.Printf("Circle 1 : %p\n", circle1)

	// Cloning the circle to create a circle with radius 5.0
	circle2 := circle1.Clone()
	circle2.(*creationalPatterns.Circle).Radius = 5.0

	fmt.Printf("Circle 2: %p\n", circle2)

	// Drawing the circles
	circle1.Draw()
	circle2.Draw()
}

/*
BUILDER PATTERN
*/
func builderPattern() {
	// creating a margherita pizza
	margheritaBuilder := creationalPatterns.NewMargheritaPizzaBuilder()
	director := &creationalPatterns.PizzaDirector{
		Builder: margheritaBuilder,
	}
	margherita := director.Construct()
	fmt.Println(margherita)

	// Creating a Pepperoni pizza
	pepperoniBuilder := creationalPatterns.NewPepperoniPizzaBuilder()
	director.Builder = pepperoniBuilder
	pepperoni := director.Construct()
	fmt.Println(pepperoni)
}

/*
ABSTRACT FACTORY PATTERNS
*/
func abstractFactoryPattern() {
	carFactory := &creationalPatterns.CarFactory{}
	car := carFactory.CreateVehicle()
	engine := carFactory.CreateEngine()

	car.Drive()
	engine.Start()

	// Creating a bike with electric engine
	bikeFactory := &creationalPatterns.BikeFactory{}
	bike := bikeFactory.CreateVehicle()
	engine = bikeFactory.CreateEngine()

	bike.Drive()
	engine.Start()
}
