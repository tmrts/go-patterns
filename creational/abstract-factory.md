# Abstract Factory Method Pattern

Abstract Factory method creational design pattern allows providing interface that creating objects without having to specify the exact type of the object that will be created.

## Implementation

The example implementation shows how to make a phone from different copmanies.

### Types

```go
package company

type iCompnayFactory interface {
	makePhone() iPhone
}
```

### Different Implementations

```go
package company

type CompnayType int

const (
	SAMSNUG CompnayType = iota
	APPLE
)

func getCompnayFactory(compnay CompnayType) (iCompnayFactory, error) {
	switch compnay {
	case SAMSNUG:
		return &samsungFactory{}, nil
	case APPLE:
		return &appleFactory{}, nil
	default:
		return nil, fmt.Errorf(/* .. */)
	}
}
```

## Usage

With the abstract factory method, the user can provide an interface for creating families of releated objects.

```go
appleFactory, _ := getCompnayFactory(APPLE)
applePhone := appleFactory.makePhone();
applePhone.makeCall(/*...*/);

samsungFactory, _ := getCompnayFactory(SAMSNUG)
samsungPhone := samsungFactory.makePhone();
samsungPhone.makeCall(/*...*/);
```
