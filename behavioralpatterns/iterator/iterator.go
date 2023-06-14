package main

import "fmt"

// Iterator represents the iterator interface.
type Iterator interface {
	HasNext() bool
	Next() string
}

// ConcreteIterator is a concrete implementation of the Iterator interface.
type ConcreteIterator struct {
	collection []string
	index      int
}

// HasNext checks if there are more elements to iterate over.
func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.collection)
}

// Next returns the next element in the collection.
func (i *ConcreteIterator) Next() string {
	if i.HasNext() {
		element := i.collection[i.index]
		i.index++
		return element
	}
	return ""
}

// Aggregate represents the aggregate interface.
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteAggregate is a concrete implementation of the Aggregate interface.
type ConcreteAggregate struct {
	collection []string
}

// CreateIterator returns a new iterator for the collection.
func (a *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{
		collection: a.collection,
		index:      0,
	}
}

func main() {
	// Create the collection
	collection := []string{"Apple", "Banana", "Orange", "Grapes"}

	// Create the aggregate
	aggregate := &ConcreteAggregate{
		collection: collection,
	}

	// Create the iterator
	iterator := aggregate.CreateIterator()

	// Iterate over the collection using the iterator
	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println(element)
	}

	// Output:
	// Apple
	// Banana
	// Orange
	// Grapes
}
