//program shows the working of defer function
package main

import "fmt"

func foo() {
	defer func() { //anonymous defer func
		fmt.Println("hello this is defer of fooooooooooooooooooooooooooooooooooo!!!!")
	}()
	fmt.Println("foo code ran")
}

func main() {
	defer foo() //defered foo
	fmt.Println("this is the main function code placed after the defer of foo")
}

/*
THIS IS THE OUTPUT:
this is the main function code placed after the defer of foo
foo code ran
hello this is defer of fooooooooooooooooooooooooooooooooooo!!!!
*/
