# Visitor Pattern
Visitor behavioral design pattern provides a way to separate an algorithm from an object on which it operates.

It gives the ability to extend the existing object without modifying the object itself.

## Implementation
Implementation of a visitor that can add functionality to the shape structures.

```go
type Visitor interface {
	visitCircle(circle Circle) string
	visitLine(line Line) string
}

type Shape interface {
	accept(Visitor) string
}

type Circle struct {
	Rad int
}

func (c Circle) accept(v Visitor) string {
	return v.visitCircle(c)
}

type Line struct {
	Len int
}

func (l Line) accept(v Visitor) string {
	return v.visitLine(l)
}
```

## Usage
### JSON marshaller
Using visitor to marshal shapes to JSON
```go
type JsonVisitor struct {
}

func (*JsonVisitor) visitCircle(circle Circle) string {
	return fmt.Sprintf(`{"type": "circle", "radius": "%v"}`, circle.Rad)
}

func (*JsonVisitor) visitLine(line Line) string {
	return fmt.Sprintf(`{"type": "line", "length": "%v"}`, line.Len)
}
```

```go
circle := Circle{12}
line := Line{42}
jsonVisitor := JsonVisitor{}
fmt.Println(circle.accept(&jsonVisitor))
fmt.Println(line.accept(&jsonVisitor))
```

### XML marshaller
Using visitor to marshal shapes to XML
```go
type XmlVisitor struct {
}

func (*XmlVisitor) visitCircle(circle Circle) string {
	return fmt.Sprintf(`<circle><radius>%d</radius></circle>`, circle.Rad)
}

func (*XmlVisitor) visitLine(line Line) string {
	return fmt.Sprintf(`<line><length>%d</length></line>`, line.Len)
}
```

```go
circle := Circle{12}
line := Line{42}
xmlVisitor := XmlVisitor{}
fmt.Println(circle.accept(&xmlVisitor))
fmt.Println(line.accept(&xmlVisitor))
```
