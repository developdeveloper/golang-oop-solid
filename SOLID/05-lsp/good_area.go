package main

import (
  "fmt"
  "math"
)

type Areable interface {
  Area() float32 
}

type Rectangle struct {
  Width float32 
  Height float32 
}

func (r *Rectangle) Area() float32 {
  return r.Width * r.Height
}

type Circle struct {
  Rectangle
}

func (c * Circle) Side() float32 {
  return ((c.Width + c.Height) * 2)/(2 * math.Pi)
}
  
func (c * Circle) Area() float32 {
  side := c.Side()
  return math.Pi * side * side
}

func printArea(area Areable) {
  fmt.Println(area.Area())
}

func main() {
  r := &Rectangle{Width: 2, Height: 4}
  printArea(r)
  
  s:= &Circle{Rectangle{Width: 2, Height: 4}}
  printArea(s)
}