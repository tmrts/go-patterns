Publish & Subscribe Messaging Pattern
============
Publish-Subscribe is a messaging pattern used to communicate messages between 
different components without these components knowing anything about each other's identity.

It is similar to the Observer behavioral design pattern. 
The fundamental design principals of both Observer and Publish-Subscribe is the decoupling of
those interested in being informed about `Event Messages` from the informer (Observers or Publishers).
Meaning that you don't have to program the messages to be sent directly to specific receivers.

To accomplish this, an intermediary, called a "message broker" or "event bus", 
receives published messages, and then routes them on to subscribers.


There are three components **messages**, **topics**, **users**.

```go
type Message struct {
    // Contents
}


type Subscription struct {
	ch chan<- Message

	Inbox chan Message
}

func (s *Subscription) Publish(msg Message) error {
	if _, ok := <-s.ch; !ok {
		return errors.New("Topic has been closed")
	}

	s.ch <- msg

	return nil
}
```

```go
type Topic struct {
	Subscribers    []Session
	MessageHistory []Message
}

func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
    // Get session and create one if it's the first

    // Add session to the Topic & MessageHistory

    // Create a subscription
}

func (t *Topic) Unsubscribe(Subscription) error {
	// Implementation
}

func (t *Topic) Delete() error {
	// Implementation
}
```

```go
type User struct {
    ID uint64
    Name string
}

type Session struct {
    User User
    Timestamp time.Time
}
```

Improvements
============
Events can be published in a parallel fashion by utilizing stackless goroutines.

Performance can be improved by dealing with straggler subscribers
by using a buffered inbox and you stop sending events once the inbox is full.
