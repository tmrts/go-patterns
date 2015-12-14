# awesome-go-patterns
 A collection of Go design patterns/idioms
 
Current Patterns:

__Creational Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [abstract_factory](abstract_factory.go) | use a generic function with specific factories |
| [borg](borg.go) | a singleton with shared-state among instances |
| [builder](builder.go) | instead of using multiple constructors, builder object receives parameters and returns constructed objects |
| [factory_method](factory_method.go) | delegate a specialized function/method to create instances |
| [lazy_evaluation](lazy_evaluation.go) | lazily-evaluated property pattern in Go |
| [pool](pool.go) | preinstantiate and maintain a group of instances of the same type |
| [prototype](prototype.go) | use a factory and clones of a prototype for new instances (if instantiation is expensive) |

__Structural Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [3-tier](3-tier.go) | data<->business logic<->presentation separation (strict relationships) |
| [adapter](adapter.go) | adapt one interface to another using a white-list |
| [bridge](bridge.go) | a client-provider middleman to soften interface changes |
| [composite](composite.go) | encapsulate and provide access to a number of different objects |
| [decorator](decorator.go) | wrap functionality with other functionality in order to affect outputs |
| [facade](facade.go) | use one class as an API to a number of others |
| [flyweight](flyweight.go) | transparently reuse existing instances of objects with similar/identical state |
| [front_controller](front_controller.go) | single handler requests coming to the application |
| [mvc](mvc.go) | model<->view<->controller (non-strict relationships) |
| [proxy](proxy.go) | an object funnels operations to something else |

__Behavioral Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [chain](chain.go) | apply a chain of successive handlers to try and process the data |
| [catalog](catalog.go) | general methods will call different specialized methods based on construction parameter |
| [chaining_method](chaining_method.go) | continue callback next object method |
| [command](command.go) | bundle a command and arguments to call later |
| [mediator](mediator.go) | an object that knows how to connect other objects and act as a proxy |
| [memento](memento.go) | generate an opaque token that can be used to go back to a previous state |
| [observer](observer.go) | provide a callback for notification of events/changes to data |
| [publish_subscribe](publish_subscribe.go) | a source syndicates events/data to 0+ registered listeners |
| [registry](registry.go) | keep track of all subclasses of a given class |
| [specification](specification.go) |  business rules can be recombined by chaining the business rules together using boolean logic |
| [state](state.go) | logic is organized into a discrete number of potential states and the next state that can be transitioned to |
| [strategy](strategy.go) | selectable operations over the same data |
| [template](template.go) | an object imposes a structure but takes pluggable components |
| [visitor](visitor.go) | invoke a callback for all items of a collection |
