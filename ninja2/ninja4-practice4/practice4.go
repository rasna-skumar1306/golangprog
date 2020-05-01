package main

import "fmt"

func main() {
	x := 31
	fmt.Printf("%d\t%b\t%#x\n", x, x, x) //here %d decimal %b binary %#x hexadecimal
	y := x << 1                          //bitshifting x by 1 and storing it to y
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
