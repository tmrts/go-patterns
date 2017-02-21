# Factory Method Pattern

Factory method creational design pattern allows creating objects without having
to specify the exact type of the object that will be created.

## Implementation

The example implementation shows how to provide a data store with different
backends such as in-memory, disk storage.

### Types

```go
package data

import "io"

type Store interface {
    Open(string) (io.ReadWriteCloser, error)
}
```

### Different Implementations

```go
package data

type StorageType int

const (
    DiskStorage StorageType = 1 << iota
    TempStorage
    MemoryStorage
)

func NewStore(t StorageType) Store {
    switch t {
    case MemoryStorage:
        return newMemoryStorage( /*...*/ )
    case DiskStorage:
        return newDiskStorage( /*...*/ )
    default:
        return newTempStorage( /*...*/ )
    }
}
```

## Usage

With the factory method, the user can specify the type of storage they want.

```go
s, _ := data.NewStore(data.MemoryStorage)
f, _ := s.Open("file")

n, _ := f.Write([]byte("data"))
defer f.Close()
```
