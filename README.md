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

__Creational Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Abstract Factory](abstract_factory.go) | Provides an interface for creating families of releated objects |
| [Builder](builder/builder.go) | Builds a complex object using simple objects |
| [Factory Method](factory_method.go) | Defers instantiation of an object to a specialized function for creating instances |
| [Object Pool](object_pool/pool.go) | Instantiates and maintains a group of objects instances of the same type |
| [Singleton](singleton/singleton.go) | Restricts instantiation of a class to one object |

__Structural Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Adapter](adapter.go) | Adapts otherwise incompatible interfaces to work together by adapting one to the other |
| [Bridge](bridge.go) | Decouples an interface from its implementation so that the two can vary independently |
| [Composite](composite.go) | Encapsulates and provides access to a number of different objects |
| [Decorator](decorator.go) | Adds behavior to an object, statically or dynamically |
| [Facade](facade.go) | Uses one class as an API to a number of others |
| [Flyweight](flyweight.go) | Reuses existing instances of objects with similar/identical state to minimize resource usage |
| [Model View Controller](mvc.go) | Divides an app into three interconnected parts to separate internal representation from presentation to user |
| [Proxy](proxy.go) | Provides a surrogate for an object to control it's actions |

__Behavioral Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Chain of Responsibility](chain_of_responsibility.go) | Avoids coupling a sender to receiver by giving more than object a chance to handle the request |
| [Command](command.go) | Bundles a command and arguments to call later |
| [Mediator](mediator.go) | Connects objects and acts as a proxy |
| [Memento](memento.go) | Generate an opaque token that can be used to go back to a previous state |
| [Observer](observer.go) | Provide a callback for notification of events/changes to data |
| [Registry](registry.go) | Keep track of all subclasses of a given class |
| [State](state.go) | Encapsulates varying behavior for the same object based on its internal state |
| [Strategy](strategy/strategy.go) | Enables an algorithm's behavior to be selected at runtime |
| [Template](template.go) | Defines a skeleton class which defers some methods to subclasses |
| [Visitor](visitor.go) | Separates an algorithm from an object on which it operates |
 
__Synchronization Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Condition Variable](condition_variable.go) | Provides a mechanism for threads to temporarily give up access in order to wait for some condition |
| [Lock/Mutex](mutex/mutex.go) | Enforces mutual exclusion limit on a resource to gain exclusive access |
| [Monitor](monitor.go) | Combination of mutex and condition variable patterns |
| [Read-Write Lock](read_write_lock.go) | Allows parallel read access, but only exclusive access on write operations to a resource |
| [Semaphore](semaphore/semaphore.go) | Allows controlling access to a common resource |

__Concurrency Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [N-Barrier](barrier.go) | Prevents a process from proceeding until all N processes reach to the barrier |
| [Bounded Parallelism](bounded_parallelism/md5.go) | Completes large number of indenpendent tasks with resource limits |
| [Broadcast](broadcast.go) | Transfers a message to all recipients simultaneously |
| [Coroutines](coroutine/coroutine.go) | Subroutines that allow suspending and resuming execution at certain locations |
| [Generators](generator.go) | Yields a sequence of values one at a time |
| [Reactor](reactor.go) | Demultiplexes service requests delivered concurrently to a service handler and dispatches them syncronously to the associated request handlers |
| [Parallelism](parallelism/md5.go) | Completes large number of indenpendent tasks |
| [Producer Consumer](producer_consumer.go) | Separates tasks from task executions |
| [Scheduler](scheduler.go) | Orchestrates steps to be performed as part of a task |

__Messaging Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Fan-In](fan/fan_in.go) | Funnels tasks to a work sink (e.g. server) |
| [Fan-Out](fan/fan_out.go) | Distributes tasks amongs workers |
| [Futures & Promises](futures_promises.go) | Acts as a place-holder of a result that is initally unknown for synchronization purposes |
| [Publish/Subscribe](messaging/publish_subscribe.md) | Passes information to a collection of recipients who subscribed to a topic |
| [Push & Pull](push_pull.go) | Distributes messages to multiple workers, arranged in a pipeline |

__Stability Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Bulkheads](bulkhead.go)  | Enforces a principle of failure containment (i.e. prevents cascading failures) |
| [Circuit-Breaker](circuitbreaker/circuit_breaker.go) | Stops the flow of the requests when requests are likely to fail |
| [Deadline](deadline.go) | Allows clients to stop waiting for a response once the probability of response becomes low (e.g. after waiting 10 seconds for a page refresh)|
| [Fail-Fast](fail_fast.go) | Checks the availability of required resources at the start of a request and fails if the requirements are not satisfied |
| [Handshaking](handshaking.go) | Asks a component if it can take any more load, if it can't the request is declined |
| [Steady-State](steady_state.go) | For every service that accumulates a resource, some other service must recycle that resource |

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
| [Cascading Failures]() | A failure in a system of interconnected parts in which the failure of a part causes a domino effect |

__Other Patterns__:

| Pattern | Description |
|:-------:| ----------- |

# License

[![Creative Commons License](http://i.creativecommons.org/l/by/4.0/88x31.png)](http://creativecommons.org/licenses/by/4.0/)

This work is licensed under a [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/).
