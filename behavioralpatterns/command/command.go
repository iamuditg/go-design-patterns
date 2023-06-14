package main

import "fmt"

// Command represents the command interface.
type Command interface {
	Execute()
}

// ConcreteCommandA is a concrete implementation of the Command interface.
type ConcreteCommandA struct {
	receiver *Receiver
}

// Execute executes the command by invoking the corresponding action on the receiver.
func (c *ConcreteCommandA) Execute() {
	c.receiver.ActionA()
}

// ConcreteCommandB is a concrete implementation of the Command interface.
type ConcreteCommandB struct {
	receiver *Receiver
}

// Execute executes the command by invoking the corresponding action on the receiver.
func (c *ConcreteCommandB) Execute() {
	c.receiver.ActionB()
}

// Receiver represents the receiver that executes the actions.
type Receiver struct{}

// ActionA represents an action performed by the receiver.
func (r *Receiver) ActionA() {
	fmt.Println("Receiver executing Action A")
}

// ActionB represents another action performed by the receiver.
func (r *Receiver) ActionB() {
	fmt.Println("Receiver executing Action B")
}

// Invoker represents the invoker that triggers the execution of commands.
type Invoker struct {
	command Command
}

// SetCommand sets the command to be executed by the invoker.
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand executes the command by invoking its Execute method.
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	// Create the receiver
	receiver := &Receiver{}

	// Create concrete commands
	commandA := &ConcreteCommandA{receiver: receiver}
	commandB := &ConcreteCommandB{receiver: receiver}

	// Create the invoker
	invoker := &Invoker{}

	// Set commandA and execute
	invoker.SetCommand(commandA)
	invoker.ExecuteCommand()

	// Set commandB and execute
	invoker.SetCommand(commandB)
	invoker.ExecuteCommand()

	// Output:
	// Receiver executing Action A
	// Receiver executing Action B
}
