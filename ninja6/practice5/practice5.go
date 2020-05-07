//program that calculates the area of circle and square using struct,interface and method function
package main

import (
	"fmt"
	"math" //this is the package which contains Pi value in-built
)

type square struct { //user defined type square
	length float64
}

type circle struct { //user defined type circle
	radius float64
}
type shape interface { //interface that connects both circle and square struct
	area() float64
}

func (s square) area() float64 { //function that returns float64 area of square
	return (s.length * s.length)
}

func (c circle) area() float64 { //function that returns float64 area of circle
	return c.radius * c.radius * math.Pi
}
func info(a shape) { //function info that takes both shape structs and calls the area function accordingly
	x := a.area()
	fmt.Printf("This is a %T and the area is %v\n", a, x)
}

func main() {
	s := square{
		length: 4,
	}
	c := circle{
		radius: 5,
	}
	info(s) //function info for square
	info(c) //func info for circle is called
}
