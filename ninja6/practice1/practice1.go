//program to use return values in functions
package main

import "fmt"

func foo() int { //function that returns an int
	x := 32
	return x
}

func bar() (int, string) { //function that returns an int and a string
	x := 123
	r := "rasna"
	return x, r
}

func main() {
	t := foo()
	x, a := bar()
	fmt.Println(t, x, a)
}
