//program to print raw string literals
package main

import "fmt"

func main() {
	a := `this is a raw string 
	literal
	"you can tell by seeing"` //using ` symbol can generate raw string literals
	fmt.Println(a)
}
