# Read/Write Lock Pattern
Allows parallel read access, but only exclusive access on write operations to a resource

## Implementation

```go
package router

import (
	"sync"
)

type Socket  string
type Peer interface {
	socket Socket
	connection net.Conn
}

func (p *Peer) Socket() Socket {
	return p.socket
}

// Router hash table to associate Socket with Peers.
// Unstructured mesh architecture
// eg. {127.0.0.1:4000: Peer}
type Router struct {
	sync.RWMutex
	table map[Socket]Peer
}


// Return connection interface based on socket
func (r *Router) Query(socket Socket) *Peer {
	// Mutex for reading topics.
	// Do not write while topics are read.
	// Write Lock can’t be acquired until all Read Locks are released.
	r.RWMutex.RLock()
	defer r.RWMutex.RUnlock()

	if peer, ok := r.table[socket]; ok {
		return peer
	}

	return nil
}

// Add create new socket connection association
func (r *Router) Add(peer *Peer) {
	// Lock write table while add operation
	// A blocked Lock call excludes new readers from acquiring the lock.
	r.RWMutex.Lock()
	r.table[peer.Socket()] = peer
        r.RWMutex.Unlock()
}

// Delete removes a connection from router
func (r *Router) Delete(peer *Peer) {
	// Lock write table while delete operation
	// A blocked Lock call excludes new readers from acquiring the lock.
	r.RWMutex.Lock()
	delete(r.table, peer.Socket())
        r.RWMutex.Unlock()
}
```

## Usage
### Syncronize routing peers from incoming connections

```go

// New router 
router:= &Router{
	table: make(map[Socket]Peer)
}

// !Important: 
// 1 - Write Lock can’t be acquired until all Read Locks are released.
// 2 - A blocked Lock call excludes new readers from acquiring the lock.

// Writing operation
go func(r *Router){
	for {
		// this will be running waiting for new connections
		/// .. some code here
		conn, err := listener.Accept()
		// eg. 192.168.1.1:8080
		remote := connection.RemoteAddr().String()
		socket := Socket(address)
		// New peer
		peer := &Peer{
			socket: socket,
			connection: conn
		}
		// Here we need a write lock to avoid race condition
		r.Add(peer)
	}
}(router)

// Reading operation
// ...some code here

// reading operation 1
connection := router.Query("192.168.1.1:8080")
	
//... more code here 
// reading operation 2
otherQuery:= router.Query("192.168.1.1:8081")
// read locks are like counters.. until counter = 0 Write can be acquired

```
