//program that generate 10 go routines which add a channel by 10 times
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //declaration of waitgroup

func main() {
	ch := make(chan int)
	go send(ch) //calling the function that generates and sends values to channel
	receive(ch) //calling the function that receives and prints the values from the channel
}

//function that generates and sends the values
func send(ch chan<- int) { //getting a generic channel (chan int) as sending channel (chan<- int)

	for i := 0; i < 10; i++ {
		wg.Add(1) //adding a goroutine to waitgroup
		go func(i int) {
			for j := 1; j <= 10; j++ {
				ch <- i + j //sending a value to channel
			}
			wg.Done() //send a signal that waitgroup has finished
		}(i)

	}
	wg.Wait() //waits till control reaches this point
	close(ch) //closing the channel
}

//function that receives and prints the values
func receive(c <-chan int) { //getting a generic channel (chan int) as a receiving channel (<-chan int)
	v := 0
	for i := range c {
		v++
		fmt.Println(v, i) //prints the channel values along with index value v
	}
}
