//program that uses anonymous function
package main

import "fmt"

func main() {
	func() { //anonymous function
		fmt.Println("This is the anonymous function speaking!!!!!")
	}()
	fmt.Println("done executing the anonymous function")
}
