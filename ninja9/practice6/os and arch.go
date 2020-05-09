//building and installing the entire program
package main

import (
	"fmt"
	"runtime" //package that contains counting function for goroutines
)

func main() {
	fmt.Println("OS:", runtime.GOOS)                   //prints the OS
	fmt.Println("ARCH:", runtime.GOARCH)               //prints the Architecture
	fmt.Println("CPUs:", runtime.NumCPU())             //prints the no of CPUs
	fmt.Println("Goroutines:", runtime.NumGoroutine()) //prints the no of Goroutines
}
