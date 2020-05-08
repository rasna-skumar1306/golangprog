//sorting the elements of []int and []string
package main

import (
	"fmt"
	"sort" //sort package
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Printf("before sorting\n%v\n", xi)
	fmt.Printf("before sorting\n%v\n", xs)
	sort.Ints(xi) //sorting the []int
	fmt.Println(xi)
	sort.Strings(xs) //sorting the []string
	fmt.Println(xs)
}
