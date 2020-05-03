//program to declare a slice and print all the elements of slice
package main

import "fmt"

func main() {
	a := []int{72, 71, 234, 34, 5, 12, 445, 22, 123, 34} //declaration of slice
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)

}
