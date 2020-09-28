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


There are three components **messages**, **topics**, **subscriptions**.

```go
type Message struct {
    // Contents
}

type Subscription struct {
	closed bool
	inbox chan Message
}

func (s *Subscription) Next() (Message, error) {
	if s.closed {
		return Message{}, errors.New("subscription closed")
	}
	
	m, ok := <-s.inbox
	if !ok {
		return Message{}, errors.New("subscription closed")
	}
	
	return m, nil
}

func (s *Subscription) Unsubscribe(Subscription) error {
	s.closed = true
	close(s.inbox)
}
```

```go
type Topic struct {
	Subscribers    []Subscription
	MessageHistory []Message
}

func (t *Topic) Subscribe() (Subscription) {
	return Subscription{inbox: make(chan Message)}
}

func (t *Topic) Publish(msg Message) error {
	for _, sub := range t.Subscribers {
		if sub.closed {
			continue
		}
		
		go func() {
			sub.inbox <- msg
		}()
	}

	return nil
}
```

Improvements
============
Events can be published in a parallel fashion by utilizing stackless goroutines.

Performance can be improved by dealing with straggler subscribers
by using a buffered inbox and you stop sending events once the inbox is full.
