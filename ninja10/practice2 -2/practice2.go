//correcting chan err part 2
package main

import (
	"fmt"
)

func main() {
	cr := make(chan int) //before it was <-chan int (receive only channel)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr) //we are sending a value down here so I changed cs <-chan int to generic channel chan int

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
