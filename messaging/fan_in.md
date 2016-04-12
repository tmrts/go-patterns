Fan-In Messaging Patterns
===================================
Fan-In is a messaging pattern used to create a funnel for work amongst workers (clients: source, server: destination).

We can model fan-in using the Go channels.

```go
// Merge different channels in one channel
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```

The `Merge` function converts a list of channels to a single channel by starting a goroutine for each inbound channel that copies the values to the sole outbound channel.

Once all the output goroutines have been started, `Merge` a goroutine is started to close the main channel.
