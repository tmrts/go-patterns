# Decorator Pattern
Decorator structural pattern allows extending the function of an existing object dynamically without altering its internals.

Decorators provide a flexible method to extend functionality of objects.

## Implementation
### Decorating single function
`LogDecorate` decorates a function with the signature `func(int) int` that 
manipulates integers and adds input/output logging capabilities.

```go
type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

        return result
	}
}
```

### Usage
```go
func Double(n int) int {
    return n * 2
}

f := LogDecorate(Double)

f(5)
// Starting execution with the integer 5
// Execution is completed with the result 10
```

### Decorating interface 
To ease decoration of interface with multiple methods, you can declare base decorator. The base decorator should simply calls all methods, then you can just override only one method in your target decorator. 


```go
type PasswordService interface {
	CheckPassword(p string) bool
	ChangePassword(p string)
	// ... more methods
}

type BaseDecoratorPasswordService struct {
	delegate PasswordService
}

func (r BaseDecoratorPasswordService) CheckPassword(a string) bool {
	return r.delegate.CheckPassword(a)
}

func (r BaseDecoratorPasswordService) ChangePassword(b string) {
	r.delegate.ChangePassword(b)
}

```

### Usage
```go
type Implementation struct {
}

func (r Implementation) CheckPassword(p string) bool {
	// .. validating password
	return true
}

func (r Implementation) ChangePassword(p string) {
	fmt.Printf("Implementation::ChangePassword(%v)\n", p)
}


func NewBigInterfaceMethodBLogger(delegate PasswordService) PasswordService {
	return &ChangePasswordLogger{BaseDecoratorPasswordService{delegate}}
}

type ChangePasswordLogger struct {
	BaseDecoratorPasswordService
}

func (r ChangePasswordLogger) ChangePassword(b string) {
	r.BaseDecoratorPasswordService.ChangePassword(b)
	fmt.Println("ChangePassword() was called")
}

func main() {
	ci := NewBigInterfaceMethodBLogger(&Implementation{})
	ci.CheckPassword("echo123")
	ci.ChangePassword("qwerty")
}
```


## Rules of Thumb
- Unlike Adapter pattern, the object to be decorated is obtained by **injection**.
- Decorators should not alter the interface of an object.


