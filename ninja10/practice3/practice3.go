//using functions for sending and receiving values through channels
package main

import (
	"fmt"
)

func main() {
	c := gen() //this function generates channels

	receive(c) //this function receives the values generated by gen()

	fmt.Println("about to exit")
}

func gen() <-chan int { //this function creates and sends value into channel. Here c is sending channel
	c := make(chan int)

	go func() { //a seperate go routine is set to keep all channel running
		for i := 0; i < 100; i++ {
			c <- i //passing values
		}
		close(c) //closing channel c
	}()

	return c
}

func receive(c <-chan int) { //this function is used to receive values. Here c is receiving channel
	for v := range c {
		fmt.Println(v) //printing the values of channel c
	}
}
