//program to print decimal,binary,hexadecimal values

package main

import "fmt"

var c = "rasna"

func main() {
	fmt.Println("decimal\tbinary\thexadecimal")
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d\t%b\t%#x\n", c[i], c[i], c[i]) //here %d decimal %b binary %#x hexadecimal
	}
}
