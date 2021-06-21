package main

import (
  "fmt"
)

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

func (c * Circle) Area() float32 {
  side := (c.Width + c.Height) * 2 / 4
  return side * side 
}

func printArea(rect *Rectangle) {
  fmt.Println(rect.Area())
}

func main() {
  r := &Rectangle{Width: 2, Height: 3}
  printArea(r)
  
  s:= &Circle{Rectangle{Width: 2, Height: 3}}
  printArea(s)
}  