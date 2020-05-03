//program to slice a "slice" and print the different slices
package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	y := append(x[:5])
	fmt.Println(y)
	z := append(x[5:])
	fmt.Println(z)
	b := append(x[2:7])
	fmt.Println(b)
	a := append(x[1:6])
	fmt.Println(a)
}

/*
[42 43 44 45 46]  => y
[47 48 49 50 51]  => z
[44 45 46 47 48]  => b
[43 44 45 46 47]  => a
*/
