package main

import (
	"fmt"
	"math"
)

// here i am defining the squre struct
type square struct {
	length float64
	width  float64
}

// here i am defining the circle struct
type circle struct {
	radius float64
}

// these 2 functions are calculating the area or a squere and a circle, i gave them both the same name so they can be used in the interface
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return s.length * s.width
}

// here i created an interface named "shape" which holds the area function of both circle and squere attached to it.
type shape interface {
	area() float64
}

func main() {
	circle := circle{radius: 4}
	square := square{length: 5, width: 8}

	// info(circle, "circle")
	// info(square, "squere")

	fmt.Println("this is the radius of circle - ", info(circle))
	fmt.Println("this is the radius of squere - ", info(square))
}

// this func will take a type of shape which us assosiated with circle and square and then activate its area function respectevly
// func info(s shape, descrip string) {
// 	a := s.area()
// 	fmt.Printf("the area of %v is %v\n", descrip, a)
// }

func info(s shape) float64 {
	return s.area()
}

/*
Hands-on exercise #62 - interfaces
● create a type SQUARE
○ length float64
○ width float64
● create a type CIRCLE ○ radius float64
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
■ math.Pi
■ math.Pow
○ squarearea=L*W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/
