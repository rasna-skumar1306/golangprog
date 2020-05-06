package main

import "fmt"

func foo() int {
	x := 32
	return x
}

func bar() (int, string) {
	x := 123
	r := "rasna"
	return x, r
}

func main() {
	t := foo()
	x, a := bar()
	fmt.Println(t, x, a)
}
