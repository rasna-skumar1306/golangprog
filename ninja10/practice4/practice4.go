//using select statement to receive and print values
package main

import (
	"fmt"
)

func main() {
	q := make(chan int) //generic channel
	c := gen(q)         //calling value generation function and passing channel q

	receive(c, q) //calling receiving function to print values

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int { //getting a generic channel as receiving channel and also returning a receiving channel
	c := make(chan int) //creating generic channel

	go func() {
		for i := 0; i < 100; i++ {
			c <- i //passing values into channel
		}
		close(c) //closing channel c
	}()

	return c //return channel c
}

func receive(c, q <-chan int) { //getting a generic channel as receiving channel
	for v := range c {
		select {
		case v = <-c: //if value from c is receive the below statement is executed
			fmt.Println(v)
		case <-q: //if value from q is received return is executed
			return
		}
	}

}
