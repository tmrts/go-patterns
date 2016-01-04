package semaphore_test

import (
	"testing"
	"time"

	"github.com/tmrts/go-patterns/semaphore"
)

func TestCreatesSemaphore(t *testing.T) {
	tickets, timeout := 1, 3*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		t.Errorf("semaphore.Acquire() got errors %s", err)
	}

	if err := s.Release(); err != nil {
		t.Errorf("semaphore.Release() got errors %s", err)
	}
}

func TestCreatesNonBlockingSemaphore(t *testing.T) {
	tickets, timeout := 0, 0
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != semaphore.ErrIllegalRelease {
		t.Errorf("non-blocking semaphore.Acquire() expected error %s got %s", semaphore.ErrIllegalRelease, err)
	}
}
