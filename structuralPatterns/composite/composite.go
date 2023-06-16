package main

import "fmt"

// Component defines the common interface for both leaf and composite components.
type Component interface {
	Operation() string
}

// Leaf represents the leaf component that doesn't have any children.
type Leaf struct {
	Name string
}

// Operation performs the operation specific to the leaf component.
func (l *Leaf) Operation() string {
	return "Leaf: " + l.Name
}

// Composite represents the composite component that can have child components.
type Composite struct {
	Name       string
	Components []Component
}

// Operation performs the operation specific to the composite component.
func (c *Composite) Operation() string {
	result := fmt.Sprintf("Composite: %s\n", c.Name)
	for _, component := range c.Components {
		result += component.Operation() + "\n"
	}
	return result
}

// AddComponent adds a new child component to the composite.
func (c *Composite) AddComponent(component Component) {
	c.Components = append(c.Components, component)
}

// RemoveComponent removes a child component from the composite.
func (c *Composite) RemoveComponent(component Component) {
	var updatedComponents []Component
	for _, comp := range c.Components {
		if comp != component {
			updatedComponents = append(updatedComponents, comp)
		}
	}
	c.Components = updatedComponents
}

func main() {
	// Create leaf components
	leaf1 := &Leaf{Name: "Leaf 1"}
	leaf2 := &Leaf{Name: "Leaf 2"}
	leaf3 := &Leaf{Name: "Leaf 3"}

	// Create composite components
	composite1 := &Composite{Name: "Composite 1"}
	composite2 := &Composite{Name: "Composite 2"}

	// Add leaf components to composite1
	composite1.AddComponent(leaf1)
	composite1.AddComponent(leaf2)

	// Add leaf components to composite2
	composite2.AddComponent(leaf3)

	// Add composite2 to composite1
	composite1.AddComponent(composite2)

	// Perform operations on composite1
	result := composite1.Operation()

	// Print the result
	fmt.Println(result)

	// Output:
	// Composite: Composite 1
	// Leaf: Leaf 1
	// Leaf: Leaf 2
	// Composite: Composite 2
	// Leaf: Leaf 3
}
