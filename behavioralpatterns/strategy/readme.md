The Strategy pattern is a behavioral design pattern that enables an object (context) to select a particular algorithm or behavior at runtime from a family of interchangeable strategies. This pattern allows for encapsulating individual strategies, making them independent and interchangeable without affecting the context that uses them.

## Components
The Strategy pattern consists of the following components:

## Strategy: 
Defines the interface that encapsulates the algorithm or behavior to be performed. It typically declares a single method that represents the strategy's behavior.

## Concrete Strategies: 
Concrete implementations of the strategy interface. Each concrete strategy encapsulates a specific algorithm or behavior.

# Context: 
Represents the context or object that uses the strategy. It maintains a reference to the current strategy and provides methods to set the strategy and execute it.

# Example
In our example, we have implemented the Strategy pattern in Go. Here's how it works:

1. We define the Strategy interface that declares the method(s) to be implemented by the concrete strategies. In our example, the Execute method represents the behavior that can vary across strategies.
2. We implement two concrete strategies, ConcreteStrategyA and ConcreteStrategyB, which implement the Strategy interface. Each concrete strategy provides its own implementation of the Execute method.
3. The Context struct represents the context or object that uses the strategy. It maintains a reference to the current strategy and provides methods to set the strategy and execute it.
4. In the main function, we create an instance of the Context as the context that will use the strategy.
5. We create concrete instances of the strategies, strategyA and strategyB.
6. We set strategyA as the strategy for the context using the SetStrategy method.
7. We execute the current strategy using the `Execute