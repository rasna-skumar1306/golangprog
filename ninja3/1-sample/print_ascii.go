//this program print the ascii equivalent of the decimal 
package main

import "fmt"

func main() {
	fmt.Println("Decimal\t\tUNICODE and ASCII")
	for i := 33; i <= 126; i++ {
		fmt.Printf("%v\t\t%#U\n", i, i)//%v prints the exact value present in i and %#U prints the ascii along with UTF-8 values
	}
}
