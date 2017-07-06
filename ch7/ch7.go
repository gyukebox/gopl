package main

import (
	"fmt"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func main() {
	myRect := rectangle{
		width:  5,
		height: 13.5,
	}
	myCir := circle{
		radius: 7.5,
	}

	fmt.Println(myRect.area())
	fmt.Println(myCir.perimeter())
	showArea(myCir, myRect)
}

func (rect rectangle) area() float64 {
	return rect.height * rect.width
}

func (rect rectangle) perimeter() float64 {
	return 2*rect.height + 2*rect.width
}

func (cir circle) area() float64 {
	return 3.14 * cir.radius * cir.radius
}

func (cir circle) perimeter() float64 {
	return 2 * 3.14 * cir.radius
}

func showArea(shapes ...shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}
