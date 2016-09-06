# Observer Pattern

The [observer pattern](https://en.wikipedia.org/wiki/Observer_pattern) allows a type instance to "publish" events to other type instances ("observers") who wish to be updated when a particular event occurs.

## Implementation

In long-running applications&mdash;such as webservers&mdash;instances can keep a collection of observers that will receive notification of triggered events.

Implementations vary, but interfaces can be used to make standard observers and notifiers:

```go
type (
	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		// Data in this case is a simple int, but the actual
		// implementation would depend on the application.
		Data int64
	}

	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// OnNotify allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		OnNotify(Event)
	}

	// Notifier is the instance being observed. Publisher is perhaps another decent
	// name, but naming things is hard.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event)
	}
)
```

## Usage

For usage, see [observer/main.go](observer/main.go) or [view in the Playground](https://play.golang.org/p/cr8jEmDmw0).
