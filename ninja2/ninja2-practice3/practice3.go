//program to use TYPED and UNTYPED variables
package main

import "fmt"

const (
	a = 31
	c = "rasna"
)

const (
	b int     = 13
	d float32 = 13.06
)

func main() {
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", d, d)
}
