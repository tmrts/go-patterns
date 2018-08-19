package main
import (
	"fmt"
	"math/rand"
)

type Shape interface {
	draw()
}

type Circle struct{
	x,y,r int
	color string
}

func (c Circle)draw() {
	fmt.Println("Draw circle..here",c.x,c.y,c.color)
}

type ShapeFactory struct{   
   shapeMap map[string]Shape
   colorArray []string
}

func NewShapeFactory()ShapeFactory{
	shapeFactoryVar:=ShapeFactory{}
	shapeFactoryVar.shapeMap=make(map[string]Shape)
	return shapeFactoryVar
}

func(s ShapeFactory)getShape(color string) Shape{
  if circleVar,ok:=s.shapeMap[color];ok {
	  fmt.Println("Cirle object of "+color +" already created..")
     return circleVar
  } 
  circleVar:=Circle{}
  circleVar.x=rand.Intn(100)
  circleVar.y=rand.Intn(100)
  circleVar.color=color
  s.shapeMap[circleVar.color] =circleVar
  return circleVar
}

func main(){
	var shapevar Shape
	shapeFactoryVar:=NewShapeFactory()
	colorArray:=[]string{"red","yellow","blue","black","orange","green","purple","white"}
	for i:=1;i<=10;i++ {
		color:=colorArray[rand.Intn(len(colorArray)-1)]
		shapevar =shapeFactoryVar.getShape(color)
		shapevar.draw()
	}
	
}