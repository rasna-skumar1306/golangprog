//using comma ok statement to get value
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 324
	v, ok := <-c //, ok returns whether channel open is true or false
	fmt.Println(v, ok)

	close(c) //closing channel

	v, ok = <-c //as the channel is closed it returns 0 and false
	fmt.Println(v, ok)
}
