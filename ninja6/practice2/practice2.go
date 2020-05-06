//program in which 2 types of parameters are passed to same kind of function
package main

import "fmt"

func foo(y ...int) int { //function that takes a variadic parameter and returns the sum of elements
	s := 0
	for _, v := range y {
		s += v
	}
	return s
}

func bar(y []int) int { //function that takes a []int as parameter and returns the sum
	s := 0
	for _, v := range y {
		s += v
	}
	return s
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := foo(a...) //unfurling the slice of int
	b := bar(a)    //sending a []int as parameter
	fmt.Println(f) //printing the result of foo
	fmt.Println(b) //printing the result of bar
}
