//program in which a function is sent as a parameter to another function
package main

import "fmt"

func geta(a func()) { // this function gets another function as a parameter
	fmt.Println("got the function dude!!")
	a()
}

func main() {
	s := func() { //anonymous function
		fmt.Println("whooooohooooo iam being passed to another function")
	}
	geta(s) //function variable s is passed as a arguement
}
