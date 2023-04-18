package main

import "fmt"

/*
The Abstract Factory Pattern is a creational design pattern that allows you to create families of related objects without specifying their concrete classes.
It provides an interface for creating families of related objects, but allows subclasses to decide which concrete classes to instantiate.
In this way, the Abstract Factory Pattern decouples the client code from the implementation details of the objects being created.
*/

/*
In this example, we define two product interfaces: Vehicle and Engine.
We then define concrete products that implement these interfaces: Car, Bike, GasEngine, and ElectricEngine.
Next, we define an abstract factory interface VehicleFactory, which provides methods to create a Vehicle and an Engine.
We then define two concrete factories: CarFactory and BikeFactory, which implement VehicleFactory.
Finally, we demonstrate the use of the Abstract Factory Pattern by creating a car with a gas engine using the CarFactory,
and a bike with an electric engine using the BikeFactory.
Overall, the Abstract Factory Pattern allows us to create families of related objects without specifying their concrete classes.
This makes our code more flexible and easier to maintain, as we can easily swap out the concrete factories to create different families of objects.
*/

// Product Interface

type Vehicle interface {
	Drive()
}

type Engine interface {
	Start()
}

// Concrete products

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving Car")
}

type Bike struct{}

func (b *Bike) Drive() {
	fmt.Println("Riding bike")
}

type GasEngine struct{}

func (g *GasEngine) Start() {
	fmt.Println("Starting gas Engine")
}

type ElectricEngine struct{}

func (e *ElectricEngine) Start() {
	fmt.Println("Starting electric engine")
}

// Abstract Factory Interface

type VehicleFactory interface {
	CreateVehicle() Vehicle
	CreateEngine() Engine
}

// Concrete Factories

type CarFactory struct{}

func (f *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

func (f *CarFactory) CreateEngine() Engine {
	return &GasEngine{}
}

type BikeFactory struct{}

func (f *BikeFactory) CreateVehicle() Vehicle {
	return &Bike{}
}

func (f *BikeFactory) CreateEngine() Engine {
	return &ElectricEngine{}
}
