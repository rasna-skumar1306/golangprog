//concurrency check of go
package main

import (
	"fmt"
	"runtime" //package that contains counting function for goroutines
	"sync"    //package that contains waitgroups
)

func foo(s int) {
	for i := 0; i < s; i++ {
		fmt.Println("f", i)
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine()) //prints the no of Goroutines
	wg.Done()                                          //after completion this returns DONE msg
}

func bar(s int) {
	for i := 0; i < s; i++ {
		fmt.Println("b", i*10)
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine()) //prints the no of Goroutines
	wg.Done()                                          //after completion this returns DONE msg
}

var wg sync.WaitGroup //declaration of waitgroup

func main() {
	wg.Add(2)   //add 2 waitgroups for foo and bar
	go foo(100) //go is used to make a new goroutine
	go bar(100)
	fmt.Println("This is the main go routine")
	fmt.Println("OS:", runtime.GOOS)                   //prints the OS
	fmt.Println("ARCH:", runtime.GOARCH)               //prints the Architecture
	fmt.Println("CPUs:", runtime.NumCPU())             //prints the no of CPUs
	fmt.Println("Goroutines:", runtime.NumGoroutine()) //prints the no of Goroutines
	wg.Wait()                                          //waits for the waitgroups to complete their execution
}
