# Mediator Pattern
The [mediator pattern](https://en.wikipedia.org/wiki/Mediator_pattern) allows to objects no longer communicate directly with each other, but instead communicate through the mediator. This reduces the dependencies between communicating objects, thereby reducing coupling. In plain words, your object(s) must know about "mediator" to communicate with another object across a "mediator". So, that allows to "mediator" implements cooperative behavior by sending a request to one or more.

## Implementation
To implement you will need:
- ``Mediator interface`` - an intermediary describing the organization of the process of information exchange between objects
- ``The Concrete Mediator``, which implements the Mediator interface
- ``Colleague interface`` - describing the organization of the process of interaction of collaborative objects with an object of the Mediator type
- ``The Concrete Colleague`` that implements the Colleague interface. Each object-colleague knows only about the object-mediator. All objects-colleagues exchange information only through an intermediary.

## Usage
### Mediator interface and Concrete Mediator to communicate between objects
```go
type Mediator interface {
	AddCollegue(c Colleague)
	Communicate(c Colleague)
}

type ConcreteMediator struct {
	Colleague *list.List
}
```

### Colleague interface and Concrete Colleague, all communications must be possible through ``mediator.Communicate()`` in ``ConcreteColleague``
```go
type Colleague interface {
	GetData() interface{}
}

type ConcreteColleague struct {
	mediator ConcreteMediator
}
```

For better explanation watch short code of chat room example: ([on playground](https://play.golang.org/p/lsBkEEfkCv)) or see local example [mediator/main.go](mediator/main.go)

## Rules of Thumb 
GoF design patterns recommends use when:
- A set of objects communicate in well-defined but a complex way. So resulting are unstructured and complex to understand.
- Reusing an object is difficult because it refers and communicates with many others objects.
- A behavior that is distributed between objects should be customizable without a subclassing.

Also, the mediator can be implemented with using the observer pattern, colleague classes act as Subjects, sending notifications to the mediator whenever they change state.