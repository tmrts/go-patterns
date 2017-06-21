# Proxy Pattern

The [proxy pattern](https://en.wikipedia.org/wiki/Proxy_pattern) provides an object that controls access to another object, intercepting all calls.

## Implementation

The proxy could interface to anything: a network connection, a large object in memory, a file, or some other resource that is expensive or impossible to duplicate.

Short idea of implementation:
```go
    // To use proxy and to object they must implement same methods
    type IObject interface {
        ObjDo(action string)
    }

    // Object represents real objects which proxy will delegate data
    type Object struct {
        action string
    }

    // ObjDo implements IObject interface and handel's all logic
    func (obj *Object) ObjDo(action string) {
        // Action behavior
        fmt.Printf("I can, %s", action)
    }

    // ProxyObject represents proxy object with intercepts actions
    type ProxyObject struct {
        object *Object
    }

    // ObjDo are implemented IObject and intercept action before send in real Object
    func (p *ProxyObject) ObjDo(action string) {
        if p.object == nil {
            p.object = new(Object)
        }
        if action == "Run" {
            p.object.ObjDo(action) // Prints: I can, Run
        }
    }
```

## Usage
More complex usage of proxy as example: User creates "Terminal" authorizes and PROXY send execution command to real Terminal object
See [proxy/main.go](proxy/main.go) or [view in the Playground](https://play.golang.org/p/mnjKCMaOVE).
