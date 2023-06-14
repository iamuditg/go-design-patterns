The Interpreter pattern is a behavioral design pattern that defines a language and its grammar representation, as well as provides an interpreter to interpret sentences in the language. It allows the interpretation of sentences or expressions and provides a way to evaluate or execute them.

## Components
The Interpreter pattern consists of the following components:

## Expression: 
Defines the interface for interpreting an expression or sentence. It typically includes a method, such as Interpret(), that evaluates or executes the expression.

## Terminal Expression: 
Represents a basic or atomic expression in the language. It implements the expression interface and provides the implementation for interpreting a terminal expression.

## Non-Terminal Expression: 
Represents a composite or complex expression in the language. It implements the expression interface and provides the implementation for interpreting a non-terminal expression by evaluating or executing its subexpressions.

## Context: 
Contains the information or context required for interpreting an expression. It can store variables, provide access to external data, or maintain state during interpretation.

## Parser: 
Parses and interprets the language grammar. It takes the input expression or sentence, breaks it down into individual tokens, and constructs the corresponding expression objects to interpret the input.

## Example
In our example, we have implemented the Interpreter pattern in

1. We define the Expression interface that declares a method, Interpret(), which represents the interpretation of an expression.
2. We implement a NumberExpression struct that represents a number expression. It implements the Expression interface and returns the value of the number when interpreted.
3. We implement PlusExpression and MinusExpression structs that represent addition and subtraction expressions, respectively. They also implement the Expression interface and evaluate the left and right expressions to perform the corresponding operation when interpreted.
4. We define a Parser struct that takes a list of tokens and parses them into expressions. The ParseExpression() method of the parser recursively builds the expression object based on the tokens and returns the corresponding expression.
5. In the main function, we create an instance of the parser with a list of tokens representing the expression "5 - 2 + 1".
6. We call ParseExpression() on the parser to construct the expression object.
7. We then call Interpret() on the expression to evaluate the expression and obtain the result.
8. The result is printed to the console.
9. By using the Interpreter pattern, we can define a language or grammar and interpret sentences or expressions in that language. It provides a way to evaluate or execute expressions based on the grammar rules and the provided input. This pattern is especially useful in scenarios where expressions need to be dynamically interpreted and executed, such as in rule engines, query languages, or mathematical calculations.