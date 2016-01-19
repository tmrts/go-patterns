<p align="center">
  <img src="/gopher.jpg" height="400">
</p>

# Go Patterns [![Travis Widget]][Travis] [![Awesome Widget]][Awesome] [![License Widget]][License]
[Awesome Widget]: https://img.shields.io/badge/awesome-%E2%9C%93-ff69b4.svg?style=flat-square
[Awesome]: https://github.com/sindresorhus/awesome
[Travis Widget]: https://img.shields.io/travis/tmrts/go-patterns.svg?style=flat-square
[Travis]: http://travis-ci.org/tmrts/go-patterns
[License Widget]: https://img.shields.io/badge/license-Creative%20Commons%204.0-E91E63.svg?style=flat-square
[License]: http://creativecommons.org/licenses/by/4.0/
A curated collection of idiomatic design & application patterns for Go language.

For use cases please see the test files (files suffixed with _test).

__Creational Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Abstract Factory](abstract_factory.go) | use a generic function with specific factories |
| [Singleton](singleton/singleton.go) | Restricts instantiation of a class to one object |
| [Builder](builder/builder.go) | instead of using multiple constructors, builder object receives parameters and returns constructed objects |
| [Factory Method](factory_method.go) | delegate a specialized function/method to create instances |
| [Lazy Evaluation](lazy_evaluation.go) | lazily-evaluated property pattern in Go |
| [Object Pool](object_pool/pool.go) | Instantiates and maintains a group of objects instances of the same type |

__Structural Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Adapter](adapter.go) | adapt one interface to another using a white-list |
| [Bridge](bridge.go) | a client-provider middleman to soften interface changes |
| [Composite](composite.go) | encapsulate and provide access to a number of different objects |
| [Decorator](decorator.go) | Adds behavior to an object, statically or dynamically |
| [Facade](facade.go) | use one class as an API to a number of others |
| [Flyweight](flyweight.go) | transparently reuse existing instances of objects with similar/identical state |
| [Model View Controller](mvc.go) | model<->view<->controller (non-strict relationships) |
| [Proxy](proxy.go) | an object funnels operations to something else |

__Behavioral Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Chain](chain.go) | apply a chain of successive handlers to try and process the data |
| [Catalog](catalog.go) | general methods will call different specialized methods based on construction parameter |
| [Chaining Method](chaining_method.go) | continue callback next object method |
| [Command](command.go) | bundle a command and arguments to call later |
| [Mediator](mediator.go) | an object that knows how to connect other objects and act as a proxy |
| [Memento](memento.go) | generate an opaque token that can be used to go back to a previous state |
| [Observer](observer.go) | provide a callback for notification of events/changes to data |
| [Registry](registry.go) | keep track of all subclasses of a given class |
| [Specification](specification.go) |  business rules can be recombined by chaining the business rules together using boolean logic |
| [State](state.go) | logic is organized into a discrete number of potential states and the next state that can be transitioned to |
| [Strategy](strategy/strategy.go) | Encapsulates an algorithm inside a struct |
| [Template](template.go) | an object imposes a structure but takes pluggable components |
| [Visitor](visitor.go) | invoke a callback for all items of a collection |
 
__Synchronization Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Lock/Mutex](mutex/mutex.go) | Enforces mutual exclusion limit on accessing a resource |
| [Read-Write Lock](read_write_lock.go) | |
| [Condition Variable](condition_variable.go) | |
| [Monitor](monitor.go) | Combination of mutex and condition variable patterns |
| [Semaphore](semaphore/semaphore.go) | Allows controlling access to a common resource |

__Concurrency Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Scheduler](scheduler.go) | |
| [Barrier](barrier.go) | |
| [Producer Consumer](producer_consumer.go) | |
| [Futures](future.go) | |
| [Broadcast](broadcast.go) | |
| [Multiplex](multiplex.go) | |
| [Generators](generator.go) | |
| [Coroutines](coroutine/coroutine.go) | |
| [Parallelism](parallelism/md5.go) | Completes large number of indenpendent tasks |
| [Bounded Parallelism](bounded_parallelism/md5.go) | Completes large number of indenpendent tasks with resource limits |

__Messaging Patterns__:
| [Fan-In](fan/fan_in.go) | Funnels tasks to a work sink (e.g. server) |
| [Fan-Out](fan/fan_out.go) | Distributes tasks amongs workers |
| [Publish/Subscribe](publish_subscribe.go) | Passes information to a collection of recipients who subscribed to a topic |
| [Request & Reply](fan) | |
| [Push & Pull](fan) | |


__Stability Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Bulkheads](bulkhead.go) | |
| [Circuit Breaker](circuitbreaker/circuit_breaker.go) | Stops the flow of the requests when requests are likely to fail |
| [Deadline](deadline.go) | |
| [Fail Fast](fail_fast.go) | |
| [Handshaking](handshaking.go) | |
| [Steady State](steady_state.go) | |

__Profiling Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Timing Functions](timing.go) | Wraps a function and logs the execution |

__Idioms__:

| Pattern | Description |
|:-------:| ----------- |
| [Functional Options](functional_options.go) | Allows creating clean APIs with sane defaults and idiomatic overrides |

__Anti-Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Cascading Failures]() | |

__Other Patterns__:

| Pattern | Description |
|:-------:| ----------- |

# License

[![Creative Commons License](http://i.creativecommons.org/l/by/4.0/88x31.png)](http://creativecommons.org/licenses/by/4.0/)

This work is licensed under a [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/).
