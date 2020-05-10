//creating a channel and pulling out values from it
package main

import "fmt"

func main() {
	ch := make(chan int) //creating a generic channel that can send and receive values
	go func() {
		ch <- 33 //sending 33 as value
	}()

	fmt.Println(<-ch) //receiving value out from channel
}

/*
A channel blocks the flow of routine hence in this program we have created a seperate goroutine for channel ch.
*/
