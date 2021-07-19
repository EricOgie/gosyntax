package main

import (
	"fmt"
	"math"
)

// Rectangle struct
type Rectangle struct {
	legth float64
	width float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Define shape interface
type shape interface {
	area() float64 // Calculate and output the shape area

}

// Implementing the shape interface of the structs

func (r Rectangle) area() float64 {
	return r.legth * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {

	rect := Rectangle{legth: 10, width: 4}
	cir := Circle{radius: 3}
	fmt.Printf("Rectangle Area : %f\n", rect.area())
	fmt.Printf("Circle Area : %f\n", cir.area())

	fmt.Printf("Circle Area : %f\n", getArea(cir)) // This will give same result

	// fmt.Println("Area of Recatangle : Area of Circle = " + strconv.FormatFloat(getArea(rect), 'E', -1, 64) + " : " + strconv.FormatFloat(getArea(cir), 'E', -1, 64))

}
