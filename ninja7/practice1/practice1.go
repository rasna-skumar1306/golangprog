//printing the address of the variable
package main

import "fmt"

func main() {
	s := 2131
	fmt.Println("Iam a variable", s)
	fmt.Println("And my address is", &s)
}
