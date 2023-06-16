package main

import "fmt"

// Target represents the target interface that the client expects to work with.
type Target interface {
	Request() string
}

// Adaptee represents the adaptee that needs to be adapted to the target interface.
type Adaptee struct {
	Data string
}

// SpecificRequest is the specific method of the adaptee.
func (a *Adaptee) SpecificRequest() string {
	return "Adaptee method: " + a.Data
}

// Adapter represents the adapter that adapts the adaptee to the target interface.
type Adapter struct {
	Adaptee *Adaptee
}

// Request adapts the specific request of the adaptee to the target interface.
func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

func main() {
	// Create an instance of the adaptee
	adaptee := &Adaptee{Data: "Adaptee Data"}

	// Create an instance of the adapter
	adapter := &Adapter{Adaptee: adaptee}

	// Call the target method through the adapter
	result := adapter.Request()

	// Print the result
	fmt.Println(result)

	// Output:
	// Adaptee method: Adaptee Data
}
