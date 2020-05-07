//program in which a function returns another function
package main

import "fmt"

func ret() func() { //function that returns another function
	fmt.Println("This function returns a function")
	return func() {
		fmt.Println("This is the returned function")
	}
}

func main() {
	s := ret()
	s()
}
