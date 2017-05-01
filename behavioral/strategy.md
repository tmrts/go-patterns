# Strategy Pattern
Strategy behavioral design pattern enables an algorithm's behavior to be selected at runtime.

It defines algorithms, encapsulates them, and uses them interchangeably.

## Implementation
Implementation of an interchangeable operator object that operates on integers.

```go
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}
```

## Usage
### Addition Operator
```go
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}
```

```go
add := Operation{Addition{}}
add.Operate(3, 5) // 8
```

### Multiplication Operator
```go
type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
```

```go
mult := Operation{Multiplication{}}

mult.Operate(3, 5) // 15
```

## Rules of Thumb
- Strategy pattern is similar to Template pattern except in its granularity.
- Strategy pattern lets you change the guts of an object. Decorator pattern lets you change the skin.
