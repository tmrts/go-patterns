# Lazy Initialisation Pattern

While not a traditional design pattern, lazy initialization is a common technique in Go. It involves deferring the creation of an object until it is actually needed. You can use this technique to optimize resource usage.

## Implementation

```go
package main

import (
	"fmt"
	"sync"
)

// LazyInitializer represents a generic lazy initializer for any type.
type LazyInitializer struct {
	mu       sync.Mutex
	instance interface{}
	initialized bool
}

// NewLazyInitializer creates a new LazyInitializer.
func NewLazyInitializer() *LazyInitializer {
	return &LazyInitializer{}
}

// GetInstance returns the instance, creating it lazily if necessary.
func (li *LazyInitializer) GetInstance(factory func() interface{}) interface{} {
	li.mu.Lock()
	defer li.mu.Unlock()

	if !li.initialized {
		li.instance = factory()
		li.initialized = true
	}
	return li.instance
}
```

## Usage

Given below is an example usage for lazy initialisation.

```go
	// Create a LazyInitializer instance.
	initializer := NewLazyInitializer()

	// Access the resource, which will be created lazily.
	resource1 := initializer.GetInstance(func() interface{} {
		fmt.Println("Creating a resource...")
		return "Resource 1"
	})
	fmt.Println("Resource 1:", resource1)

	// Access the resource again, which should reuse the existing one.
	resource2 := initializer.GetInstance(func() interface{} {
		fmt.Println("Creating a resource...")
		return "Resource 2"
	})
	fmt.Println("Resource 2:", resource2)

	// The resource is created only once, and subsequent accesses reuse it.
```

## Rules of Thumb

- The lazy initialization design pattern is used when you want to defer the creation of an object or resource until it's actually needed, rather than creating it immediately.
- It's important to note that while lazy initialization can be a useful optimization technique, it should be used judiciously.
-  In some cases, it might introduce additional complexity or latency if used excessively.
