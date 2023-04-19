package behavioralpatterns

import "fmt"

/*
The Chain of Responsibility Pattern is a behavioral design pattern that allows you to pass requests through a chain of handlers.
Each handler in the chain can decide to handle the request or pass it on to the next handler in the chain.
*/

/*
In this example, we define a Handler interface, which provides methods for setting the next handler in the chain (SetNext()) and for handling a request (Handle()).
We define a BaseHandler struct, which implements the Handler interface and has a next attribute of type Handler. The Handle() method of BaseHandler checks
if there is a next handler in the chain and if so, passes the request to it.
We define three concrete handlers: ConcreteHandler1, ConcreteHandler2, and ConcreteHandler3,
which embed the BaseHandler struct and implement their own Handle() methods.
Each concrete handler checks if it can handle the request and if so, returns a result.
If not, it calls the Handle() method of the next handler in the chain.
Finally, we demonstrate the use of the Chain of Responsibility Pattern by creating three concrete handlers and setting up a chain of handlers
by calling SetNext() on each handler in the chain.
We then send requests to the first handler in the chain, and observe the output as the requests are passed through the chain of handlers.
Overall, the Chain of Responsibility Pattern allows you to pass requests through a chain of handlers, each of which can decide to handle
*/

// Handler interface
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request int) string
}

// BaseHandler struct
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(request int) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

// ConcreteHandler1 struct
type ConcreteHandler1 struct {
	BaseHandler
}

func (h *ConcreteHandler1) Handle(request int) string {
	if request >= 0 && request < 10 {
		return fmt.Sprintf("Handled by ConcreteHandler1: %d", request)
	}
	return h.BaseHandler.Handle(request)
}

// Concrete handler 2
type ConcreteHandler2 struct {
	BaseHandler
}

func (h *ConcreteHandler2) Handle(request int) string {
	if request >= 10 && request < 20 {
		return fmt.Sprintf("Handled by ConcreteHandler2: %d", request)
	}
	return h.BaseHandler.Handle(request)
}

// Concrete handler 3
type ConcreteHandler3 struct {
	BaseHandler
}

func (h *ConcreteHandler3) Handle(request int) string {
	if request >= 20 && request < 30 {
		return fmt.Sprintf("Handled by ConcreteHandler3: %d", request)
	}
	return h.BaseHandler.Handle(request)
}

func main() {
	// Creating handlers
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}

	// Setting up the chain of handlers
	handler1.SetNext(handler2).SetNext(handler3)

	// Sending requests to the first handler in the chain
	for _, request := range []int{5, 15, 25, 35} {
		result := handler1.Handle(request)
		if result != "" {
			fmt.Println(result)
		} else {
			fmt.Println("No handler found for:", request)
		}
	}
}
