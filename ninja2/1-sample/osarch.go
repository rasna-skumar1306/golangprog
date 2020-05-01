package main

import (
	"fmt"
	"runtime"
)

func main() {
	//this program displays the OS and Architecture of the system which it is running
	fmt.Println(runtime.GOOS)   //displays the os of the system
	fmt.Println(runtime.GOARCH) //displays the architecture of the system
}
