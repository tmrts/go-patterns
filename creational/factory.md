# Factory Method Pattern

Factory method creational design pattern allows creating objects without having
to specify the exact type of the object that will be created.  
工厂方法创建型设计模式允许创建对象，不用指定要创建对象的类型。


## Implementation

The example implementation shows how to provide a data store with different
backends such as in-memory, disk storage.  
示例实现展示了如何使用不同的后端，如内存、磁盘存储。

### Types(对象类型)

```go
package data

import "io"

type Store interface {
    Open(string) (io.ReadWriteCloser, error)
}
```

### Different Implementations(不同实现)

```go
package data

type StorageType int

// 类型实现对象枚举
const (
    DiskStorage StorageType = 1 << iota
    TempStorage
    MemoryStorage
)

// 工厂方法(实例化对象)
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
使用工厂方法，指定实现类型枚举。

```go
// 给工厂方法提供实现类型枚举
s, _ := data.NewStore(data.MemoryStorage)
f, _ := s.Open("file")

n, _ := f.Write([]byte("data"))
defer f.Close()
```
