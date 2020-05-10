//correcting chan type err
package main

import (
	"fmt"
)

func main() {
	cs := make(chan int) //before it was chan<- int (send only channel)

	go func() {
		cs <- 42 //we are recieving a value down here so I changed cs chan<- int to generic channel chan int
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
