The Command pattern is a behavioral design pattern that encapsulates a request as an object, thereby allowing parameterizing clients with different requests, queue or log requests, and support undoable operations. It decouples the sender of a request from its receiver, enabling greater flexibility and extensibility in handling requests.

## Components
The Command pattern consists of the following components:

## Command: 
Defines the interface for executing a command. It declares a single method, such as Execute(), that encapsulates the actions associated with the command.

## Concrete Commands: 
Concrete implementations of the command interface. Each concrete command encapsulates a specific action and contains a reference to the receiver that performs the action.

## Receiver: 
Represents the object that performs the actions associated with the commands. It defines the actual behavior for each action.

## Invoker: 
Triggers the execution of commands by invoking their Execute() method. It contains a reference to the command and controls when and how the command is executed.

## Example
In our example, we have implemented the Command pattern in Go. Here's how it works:

1. We define the Command interface that declares a single method, Execute(), which encapsulates the actions associated with the command.
2. We implement two concrete commands, ConcreteCommandA and ConcreteCommandB, which encapsulate specific actions and contain a reference to the receiver.
3. The Receiver struct represents the object that performs the actions associated with the commands. It defines the actual behavior for each action, such as ActionA and ActionB.
4. The Invoker struct represents the invoker that triggers the execution of commands. It contains a reference to the command and provides methods to set the command and execute it.
5. In the main function, we create an instance of the Receiver as the receiver that will perform the actions.
6. We create concrete instances of the commands, commandA and commandB, passing the receiver as a parameter.
7. We create an instance of the Invoker as the invoker that will execute the commands.
8. We set commandA as the command for the invoker using the SetCommand method and execute it using the ExecuteCommand method.
9. We then set commandB as the command for the invoker and execute it.
10. As a result, the invoker triggers the execution of the respective commands, which in turn invoke the corresponding actions on the receiver.
11. By using the Command pattern, we achieve decoupling between the sender of a request and its receiver. The sender doesn't need to know the specific actions performed by the receiver, and the receiver doesn't need to be aware of the sender. This promotes flexibility, extensibility, and maintainability in the codebase, as new commands can be added without modifying existing classes.