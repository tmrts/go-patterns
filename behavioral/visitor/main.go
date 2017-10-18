package main

import (
	"fmt"
)

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

type JsonVisitor struct {
}

func (*JsonVisitor) visitCircle(circle Circle) string {
	return fmt.Sprintf(`{"type": "circle", "radius": "%v"}`, circle.Rad)
}

func (*JsonVisitor) visitLine(line Line) string {
	return fmt.Sprintf(`{"type": "line", "length": "%v"}`, line.Len)
}

type XmlVisitor struct {
}

func (*XmlVisitor) visitCircle(circle Circle) string {
	return fmt.Sprintf(`<circle><radius>%d</radius></circle>`, circle.Rad)
}

func (*XmlVisitor) visitLine(line Line) string {
	return fmt.Sprintf(`<line><length>%d</length></line>`, line.Len)
}

func main() {
	circle := Circle{12}
	line := Line{42}
	jsonVisitor := JsonVisitor{}
	fmt.Println(circle.accept(&jsonVisitor))
	fmt.Println(line.accept(&jsonVisitor))

	xmlVisitor := XmlVisitor{}
	fmt.Println(circle.accept(&xmlVisitor))
	fmt.Println(line.accept(&xmlVisitor))
}
