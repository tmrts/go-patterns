Fan-Out Messaging Pattern
=========================
Fan-Out is a messaging pattern used for distributing work amongst workers (producer: source, consumers: destination).

We can model fan-out using the Go channels.

```go
// Split a channel into n channels that receive messages in a round-robin fashion.
func Split(ch <-chan int, n int) []<-chan int {
	cs := make([]chan int)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	// Distributes the work in a round robin fashion among the stated number
	// of channels until the main channel has been closed. In that case, close
	// all channels and return.
	distributeToChannels := func(ch <-chan int, cs []chan<- int) {
		// Close every channel when the execution ends.
		defer func(cs []chan<- int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}
```

The `Split` function converts a single channel into a list of channels by using 
a goroutine to copy received values to channels in the list in a round-robin fashion.
