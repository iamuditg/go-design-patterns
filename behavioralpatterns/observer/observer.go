package main

import (
	"fmt"
)

// Subject represents the subject being observed.
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObserver()
}

// Observer represents the observer that receives updates.
type Observer interface {
	Update(value int)
}

// ConcreteSubject is a concrete implementation of the subject interface.
type ConcreteSubject struct {
	observers []Observer
	value     int
}

// RegisterObserver adds an observer to the subject's observer list.
func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver removes an observer from the subject's observer list.
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifies all registered observers about the state change.
func (s *ConcreteSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.value)
	}
}

// SetValue updates the subject's value and notifies all observers.
func (s *ConcreteSubject) SetValue(value int) {
	s.value = value
	s.NotifyObservers()
}

// ConcreteObserver is a concrete implementation of the Observer interface
type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(value int) {
	fmt.Printf("Observer %s received update: %d\n", o.name, value)
}

func main() {
	// Create the subject
	subject := &ConcreteSubject{}

	// Create observers
	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	// Register observers to the subject
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// set a new value on the subject , which triggers notification to the observers
	for i := 0; i < 10; i++ {
		subject.SetValue(i)
	}

	// Output :
	// Observer Observer 1 received update: 10
	// Observer Observer 2 received update: 10

	subject.RemoveObserver(observer1)
	subject.RemoveObserver(observer2)

	fmt.Println("observer removed")
}
