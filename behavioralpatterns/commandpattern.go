package behavioralpatterns

import "fmt"

/*
The Command Pattern is a behavioral design pattern that allows you to encapsulate a request as an object,
thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.
*/

// Command interface
type Command interface {
	Execute()
}

// Receiver
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver Action executing")
}

// Concrete command
type ConcreteCommand struct {
	receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	fmt.Println("ConcreteCommand executing")
	c.receiver.Action()
}

// Invoker
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	fmt.Println("Executing command:")
	i.command.Execute()
}

func main() {
	// Creating receiver
	receiver := &Receiver{}

	// Creating concrete command
	command := &ConcreteCommand{receiver: receiver}

	// Creating invoker
	invoker := &Invoker{}

	// Setting command on invoker
	invoker.SetCommand(command)

	// Executing command
	invoker.ExecuteCommand()
}
