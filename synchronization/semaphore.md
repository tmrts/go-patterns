# Semaphore Pattern
A semaphore is a synchronization pattern/primitive that imposes mutual exclusion on a limited number of resources. 

## Implementation

```go
package semaphore

var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

// Interface contains the behavior of a semaphore that can be acquired and/or released.
type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}

	return nil
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}
```

## Usage
### Semaphore with Timeouts

```go
tickets, timeout := 1, 3*time.Second
s := semaphore.New(tickets, timeout)

if err := s.Acquire(); err != nil {
    panic(err)
}

// Do important work

if err := s.Release(); err != nil {
    panic(err)
}
```
### Semaphore without Timeouts (Non-Blocking)

```go
tickets, timeout := 0, 0
s := semaphore.New(tickets, timeout)

if err := s.Acquire(); err != nil {
    if err != semaphore.ErrNoTickets {
        panic(err)
    }

    // No tickets left, can't work :(
    os.Exit(1)
}
```
