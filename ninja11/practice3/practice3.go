//custom error function
package main

import (
	"fmt"
)

type customErr struct { //struct that contains the info of the error
	info string
}

func (ce customErr) Error() string { //function to print the error
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more clarity", //error value
	}

	foo(c1) //calling func foo
}

func foo(e error) {
	fmt.Println("foo ran -", e) //collects error customErr and prints
}
