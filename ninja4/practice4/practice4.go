//playing more with slices
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} //declaration of slice
	y := []int{56, 57, 58, 59, 60}
	x = append(x, 52) //appending one element to x
	fmt.Println(x)
	x = append(x, 53, 54, 55) //appending 3 or more in one append command
	fmt.Println(x)
	y = append(x, y...) //appending one slice to another
	fmt.Println(y)
}
