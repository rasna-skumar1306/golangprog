//using channel buffers to store values
package main

import "fmt"

func main() {
	ch := make(chan int, 2) //creating a generic channel that also has a buffer of size 2
	ch <- 321               //sending 321
	ch <- 123               //sending 123
	fmt.Println(<-ch)       //receiving out 321 out of channel
	fmt.Println(<-ch)       //receiving out 123 out of channel
}

/*
here we are not using seperate goroutine for channel, instead we are using buffers to store values
*/
