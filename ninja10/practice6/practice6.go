//two function generate and receive the values through channel
package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c) //creating a separate goroutine for generating and sending values
	receive(c) //calling the receiving function
	fmt.Println("Exiting....")
}

func send(c chan<- int) { //getting a generic channel as sending channel
	for i := 0; i < 100; i++ {
		c <- i //sending values to c
	}
	close(c) //closing the channel c
}

func receive(c <-chan int) { //getting a generic channel as receiving channel
	for i := range c {
		fmt.Println(i)
	}
}
