//program that uses nested for loop for printing utf-8 capital alphabets
package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%v\t%#U\n", i, i)
		}
		fmt.Printf("\n")
	}
}
