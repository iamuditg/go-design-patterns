package main

import "fmt"

func main() {

	// abstractFactoryPattern()
	builderPattern()
}

/*
BUILDER PATTERN
*/
func builderPattern() {
	// creating a margherita pizza
	margheritaBuilder := NewMargheritaPizzaBuilder()
	director := &PizzaDirector{
		builder: margheritaBuilder,
	}
	margherita := director.Construct()
	fmt.Println(margherita)

	// Creating a Pepperoni pizza
	pepperoniBuilder := NewPepperoniPizzaBuilder()
	director.builder = pepperoniBuilder
	pepperoni := director.Construct()
	fmt.Println(pepperoni)
}

/*
ABSTRACT FACTORY PATTERNS
*/
func abstractFactoryPattern() {
	carFactory := &CarFactory{}
	car := carFactory.CreateVehicle()
	engine := carFactory.CreateEngine()

	car.Drive()
	engine.Start()

	// Creating a bike with electric engine
	bikeFactory := &BikeFactory{}
	bike := bikeFactory.CreateVehicle()
	engine = bikeFactory.CreateEngine()

	bike.Drive()
	engine.Start()
}
