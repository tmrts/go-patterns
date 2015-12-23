package semaphore

import (
	"errors"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

type Interface interface {
	Acquire() error
	Release() error
}

type semaphore struct {
	sem     chan struct{}
	timeout time.Duration
}

func New(tickets int, timeout time.Duration) Semaphore {
	return &semaphore{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

func (s *semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *semaphore) Release() error {
	_, ok := <-s.sem
	if !ok {
		return ErrIllegalRelease
	}

	return nil
}
