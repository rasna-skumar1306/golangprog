// program to print a multidimensional slice
package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	x := [][]string{a, b}
	fmt.Println(x)
	for i, v := range x {
		fmt.Printf("record %d\n", i)
		for j, val := range v {
			fmt.Printf("%d\t%v\n", j, val)
		}
	}
}
