package structuralPatterns

/*
The Adapter Pattern is a structural design pattern that allows objects with incompatible interfaces to work together.
It acts as a wrapper around an existing class,
allowing it to work with other classes that have different interfaces.
*/

// Target interface
type Target interface {
	Request() string
}

// Adaptee interface
type Adaptee interface {
	SpecificRequest() string
}

// concrete adaptee
type ConcreteAdaptee struct{}

func (a *ConcreteAdaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter
type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Request() string {
	return "adapter " + a.adaptee.SpecificRequest()
}
