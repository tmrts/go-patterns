# Bridge Pattern
The [bridge pattern](https://en.wikipedia.org/wiki/Bridge_pattern) allows you to "decouple an abstraction from its implementation so that the two can vary independently". It does so by creating two hierarchies: Abstraction and Implementation.

```
  Abstraction       |           Implementation
   Hierarchy        |             Hierarchy
                    |
 -------------      |         ------------------
| Abstraction |     |   imp  |  <Implementor>   |
|-------------| ----|------> |------------------|
| + imp       |     |        | implementation() |
 -------------      |         ------------------
                    |                  ^
                    |                  |
                    |        ---------------------
                    |       | ConcreteImplementor |
                    |       |---------------------|
                    |       |   implementation()  |
                    |        ---------------------
```

Note: In the literature, the `Abstraction` class is commonly represented as an "Abstract Class", meaning, children should be defined to instantiate it. Since Go does not explicitly support inheritance (and it has good reasons), that part was simplified by a concrete class modeled as a Struct.

## Implementation
```go
    // Abstraction represents the concretion of the abstraction hierarchy of the bridge
    type Abstraction struct {
        imp Implementor
    }

    // Implementor represents the abstraction of the implementation hierarchy of the bridge
    type Implementor interface {
        implementation()
    }

    // ConcreteImplementor implements Implementor
    type ConcreteImplementor struct{}

    func (c *ConcreteImplementor) implementation() {
        fmt.Println(`Some implementation here...`)
    }
```

## Usage
```go
    myObj := Abstraction{&ConcreteImplementor{}}

    myObj.imp.implementation()
```
[view in the Playground](https://play.golang.org/p/qlFOfjYX5YQ)
