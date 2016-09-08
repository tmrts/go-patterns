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
| TODO: [Abstract Factory](creational/abstract_factory.md) | Provides an interface for creating families of releated objects |
| TODO: [Builder](creational/builder.md) | Builds a complex object using simple objects |
| TODO: [Factory Method](creational/factory.md) | Defers instantiation of an object to a specialized function for creating instances |
| [Object Pool](creational/object_pool.md) | Instantiates and maintains a group of objects instances of the same type |
| [Singleton](creational/singleton.md) | Restricts instantiation of a type to one object |

__Structural Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [Adapter](structural/.md) | Adapts otherwise incompatible interfaces to work together by adapting one to the other |
| TODO: [Bridge](structural/bridge.md) | Decouples an interface from its implementation so that the two can vary independently |
| TODO: [Composite](structural/composite.md) | Encapsulates and provides access to a number of different objects |
| [Decorator](structural/decorator.md) | Adds behavior to an object, statically or dynamically |
| TODO: [Facade](structural/facade.md) | Uses one type as an API to a number of others |
| TODO: [Flyweight](structural/flyweight.md) | Reuses existing instances of objects with similar/identical state to minimize resource usage |
| TODO: [Model View Controller](structural/model_view_controller.md) | Divides an app into three interconnected parts to separate internal representation from presentation to user |
| TODO: [Proxy](structural/proxy.md) | Provides a surrogate for an object to control it's actions |

__Behavioral Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [Chain of Responsibility](behavioral/chain_of_responsibility.md) | Avoids coupling a sender to receiver by giving more than object a chance to handle the request |
| TODO: [Command](behavioral/command.md) | Bundles a command and arguments to call later |
| TODO: [Mediator](behavioral/mediator.md) | Connects objects and acts as a proxy |
| TODO: [Memento](behavioral/memento.md) | Generate an opaque token that can be used to go back to a previous state |
| [Observer](behavioral/observer.md) | Provide a callback for notification of events/changes to data |
| TODO: [Registry](behavioral/registry.md) | Keep track of all subclasses of a given class |
| TODO: [State](behavioral/state.md) | Encapsulates varying behavior for the same object based on its internal state |
| [Strategy](behavioral/strategy.md) | Enables an algorithm's behavior to be selected at runtime |
| TODO: [Template](behavioral/template.md) | Defines a skeleton class which defers some methods to subclasses |
| TODO: [Visitor](behavioral/visitor.md) | Separates an algorithm from an object on which it operates |
 
__Synchronization Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [Condition Variable](synchronization/condition_variable.md) | Provides a mechanism for threads to temporarily give up access in order to wait for some condition |
| TODO: [Lock/Mutex](synchronization/mutex.md) | Enforces mutual exclusion limit on a resource to gain exclusive access |
| TODO: [Monitor](synchronization/monitor.md) | Combination of mutex and condition variable patterns |
| TODO: [Read-Write Lock](synchronization/read_write_lock.md) | Allows parallel read access, but only exclusive access on write operations to a resource |
| [Semaphore](synchronization/semaphore.md) | Allows controlling access to a common resource |

__Concurrency Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [N-Barrier](concurrency/barrier.md) | Prevents a process from proceeding until all N processes reach to the barrier |
| [Bounded Parallelism](concurrency/bounded_parallelism.md) | Completes large number of independent tasks with resource limits |
| TODO: [Broadcast](concurrency/broadcast.md) | Transfers a message to all recipients simultaneously |
| TODO: [Coroutines](concurrency/coroutine.md) | Subroutines that allow suspending and resuming execution at certain locations |
| [Generators](concurrency/generator.md) | Yields a sequence of values one at a time |
| TODO: [Reactor](concurrency/reactor.md) | Demultiplexes service requests delivered concurrently to a service handler and dispatches them syncronously to the associated request handlers |
| [Parallelism](concurrency/parallelism.md) | Completes large number of independent tasks |
| TODO: [Producer Consumer](concurrency/producer_consumer.md) | Separates tasks from task executions |
| TODO: [Scheduler](concurrency/scheduler.md) | Orchestrates steps to be performed as part of a task |

__Messaging Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| [Fan-In](messaging/fan_in.md) | Funnels tasks to a work sink (e.g. server) |
| [Fan-Out](messaging/fan_out.md) | Distributes tasks among workers (e.g. producer) |
| TODO: [Futures & Promises](messaging/futures_promises.md) | Acts as a place-holder of a result that is initially unknown for synchronization purposes |
| [Publish/Subscribe](messaging/publish_subscribe.md) | Passes information to a collection of recipients who subscribed to a topic |
| TODO: [Push & Pull](messaging/push_pull.md) | Distributes messages to multiple workers, arranged in a pipeline |

__Stability Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [Bulkheads](stability/bulkhead.md)  | Enforces a principle of failure containment (i.e. prevents cascading failures) |
| [Circuit-Breaker](stability/circuit_breaker.md) | Stops the flow of the requests when requests are likely to fail |
| TODO: [Deadline](stability/deadline.md) | Allows clients to stop waiting for a response once the probability of response becomes low (e.g. after waiting 10 seconds for a page refresh)|
| TODO: [Fail-Fast](stability/fail_fast.md) | Checks the availability of required resources at the start of a request and fails if the requirements are not satisfied |
| TODO: [Handshaking](stability/handshaking.md) | Asks a component if it can take any more load, if it can't the request is declined |
| TODO: [Steady-State](stability/steady_state.md) | For every service that accumulates a resource, some other service must recycle that resource |

__Profiling Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [Timing Functions](profiling/timing.md) | Wraps a function and logs the execution |

__Idioms__:

| Pattern | Description |
|:-------:| ----------- |
| [Functional Options](idiom/functional-options.md) | Allows creating clean APIs with sane defaults and idiomatic overrides |

__Anti-Patterns__:

| Pattern | Description |
|:-------:| ----------- |
| TODO: [Cascading Failures](antipatterns/cascading_failures.md) | A failure in a system of interconnected parts in which the failure of a part causes a domino effect |

__Other Patterns__:

| Pattern | Description |
|:-------:| ----------- |

# License

[![Creative Commons License](http://i.creativecommons.org/l/by/4.0/88x31.png)](http://creativecommons.org/licenses/by/4.0/)

This work is licensed under a [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/).
