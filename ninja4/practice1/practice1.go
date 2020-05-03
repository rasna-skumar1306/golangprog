//program to declare an array and print all the elements of the array using for loop
package main

import "fmt"

func main() {
	a := [5]int{72, 71, 234, 34, 5} //declaration of an array
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)

}
