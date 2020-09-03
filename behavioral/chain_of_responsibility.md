# Observer Pattern

The [chain of responsibility pattern](https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern) is a design pattern consisting in a series of independent processes that will run in in sequence. It gives power to rearange the sequence or add new processes at the end of it without changing its structure.

## Implementation

Each process in the chain must implement a function to be executed and call the next process at the end.
It will call each process recursivelly until there is no more porcesses in the chain of responsibility.

## Usage

For usage, see [chain_of_responsibility/main.go](chain_of_responsibility/main.go) or [view in the Playground](https://play.golang.org/p/XA2v0XlenAi).
