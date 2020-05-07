//assigning a function to a variable
package main

import "fmt"

func main() {
	s := func() { //function assigned to variable s
		fmt.Println("This is the anonymous function assigned to a variable")
	}

	s() //calling function s
	fmt.Println("done executing the function")
}
