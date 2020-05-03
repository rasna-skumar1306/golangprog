//program that deletes certain elements of slice
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} //declaration of slice
	x = append(x[:3], x[6:]...)                        //slicing [42, 43, 44, 48, 49, 50, 51] from x

	fmt.Println(x)

}
