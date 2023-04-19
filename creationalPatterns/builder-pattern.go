package creationalPatterns

import "fmt"

/*
The Builder Pattern is a creational design pattern that allows you to create complex objects step by step.
It separates the construction of a complex object from its representation,
so that the same construction process can create different representations.
*/

/*
In this example, we define a Pizza struct, which represents the product we want to build.
We then define a PizzaBuilder interface, which provides methods to set the dough, sauce, and toppings of the pizza,
and a method to get the final pizza.
We then define two concrete builders: MargheritaPizzaBuilder and PepperoniPizzaBuilder, which implement `
*/

// Product

type Pizza struct {
	dough    string
	sauce    string
	toppings []string
}

func (p *Pizza) String() string {
	return fmt.Sprintf("Pizza with %s dough, %s sauce, and toppings: %v", p.dough, p.sauce, p.toppings)
}

// Builder interface

type PizzaBuilder interface {
	SetDough() PizzaBuilder
	SetSauce() PizzaBuilder
	SetToppings() PizzaBuilder
	GetPizza() *Pizza
}

// Concrete builders

type MargheritaPizzaBuilder struct {
	pizza *Pizza
}

func NewMargheritaPizzaBuilder() *MargheritaPizzaBuilder {
	return &MargheritaPizzaBuilder{pizza: &Pizza{}}
}

func (b *MargheritaPizzaBuilder) SetDough() PizzaBuilder {
	b.pizza.dough = "thin"
	return b
}

func (b *MargheritaPizzaBuilder) SetSauce() PizzaBuilder {
	b.pizza.sauce = "tomato"
	return b
}

func (b *MargheritaPizzaBuilder) SetToppings() PizzaBuilder {
	b.pizza.toppings = []string{"cheese", "basil"}
	return b
}

func (b *MargheritaPizzaBuilder) GetPizza() *Pizza {
	return b.pizza
}

type PepperoniPizzaBuilder struct {
	pizza *Pizza
}

func NewPepperoniPizzaBuilder() *PepperoniPizzaBuilder {
	return &PepperoniPizzaBuilder{pizza: &Pizza{}}
}

func (b *PepperoniPizzaBuilder) SetDough() PizzaBuilder {
	b.pizza.dough = "thick"
	return b
}

func (b *PepperoniPizzaBuilder) SetSauce() PizzaBuilder {
	b.pizza.sauce = "tomato"
	return b
}

func (b *PepperoniPizzaBuilder) SetToppings() PizzaBuilder {
	b.pizza.toppings = []string{"cheese", "pepperoni", "mushrooms", "onions"}
	return b
}

func (b *PepperoniPizzaBuilder) GetPizza() *Pizza {
	return b.pizza
}

// Director
type PizzaDirector struct {
	Builder PizzaBuilder
}

func (d *PizzaDirector) Construct() *Pizza {
	return d.Builder.SetDough().SetSauce().SetToppings().GetPizza()
}
