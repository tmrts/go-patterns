# Proxy Pattern

The [proxy pattern](https://en.wikipedia.org/wiki/Proxy_pattern) provides an object that controls access to another object, intercepting all calls.

## Implementation

The proxy could interface to anything: a network connection, a large object in memory, a file, or some other resource that is expensive or impossible to duplicate.

Short idea of implementation:
```go
    type IObject interface {
        ObjDo(action string)
    }

    type Object struct {
        action string
    }

    func (obj *Object) ObjDo(action string) {
        // Action handler
        fmt.Printf("I do, %s", action)
    }

    type ProxyObject struct {
        object *Object
    }

    func (p *ProxyObject) ObjDo(action string) {
        if p.object == nil {
            p.object = new(Object)
        }
        if action == "Run" {
            p.object.ObjDo(action)
        }
    }

    func main() {
        newObj := new(ProxyObject)
        newObj.ObjDo("Run") // Prints: I can, Run
    }
```

## Usage
For usage, see [observer/main.go](proxy/main.go) or [view in the Playground](https://play.golang.org/p/cr8jEmDmw0).
