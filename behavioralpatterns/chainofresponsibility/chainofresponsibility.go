package main

import "fmt"

// Request represents a request that needs to be handled.
type Request struct {
	content string
}

// Handler represents the handler interface.
type Handler interface {
	SetNext(handler Handler)
	Handle(request Request)
}

// ConcreteHandlerA represents a concrete handler that handles requests of a certain type.
type ConcreteHandlerA struct {
	next Handler
}

// SetNext sets the next handler in the chain.
func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

// Handle handles the request. If the request can be handled by this handler, it does so. Otherwise, it passes the request to the next handler in the chain.
func (h *ConcreteHandlerA) Handle(request Request) {
	if request.content == "A" {
		fmt.Println("ConcreteHandlerA: Handling request", request.content)
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerA: Unable to handle request", request.content)
	}
}

// ConcreteHandlerB represents a concrete handler that handles requests of a certain type.
type ConcreteHandlerB struct {
	next Handler
}

// SetNext sets the next handler in the chain.
func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

// Handle handles the request. If the request can be handled by this handler, it does so. Otherwise, it passes the request to the next handler in the chain.
func (h *ConcreteHandlerB) Handle(request Request) {
	if request.content == "B" {
		fmt.Println("ConcreteHandlerB: Handling request", request.content)
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerB: Unable to handle request", request.content)
	}
}

// ConcreteHandlerC represents a concrete handler that handles requests of a certain type.
type ConcreteHandlerC struct {
	next Handler
}

// SetNext sets the next handler in the chain.
func (h *ConcreteHandlerC) SetNext(handler Handler) {
	h.next = handler
}

// Handle handles the request. If the request can be handled by this handler, it does so. Otherwise, it passes the request to the next handler in the chain.
func (h *ConcreteHandlerC) Handle(request Request) {
	if request.content == "C" {
		fmt.Println("ConcreteHandlerC: Handling request", request.content)
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("ConcreteHandlerC: Unable to handle request", request.content)
	}
}

func main() {
	// Create the handlers
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	// Set the chain of responsibility
	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	// Create requests
	request1 := Request{content: "A"}
	request2 := Request{content: "B"}
	request3 := Request{content: "C"}
	request4 := Request{content: "D"}

	// Process the requests
	handlerA.Handle(request1)
	handlerA.Handle(request2)
	handlerA.Handle(request3)
	handlerA.Handle(request4)

	// Output:
	// ConcreteHandlerA: Handling request A
	// ConcreteHandlerB: Handling request B
	// ConcreteHandlerC: Handling request C
	// ConcreteHandlerA: Unable to handle request D
}
