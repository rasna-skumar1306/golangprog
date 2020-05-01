package main

import "fmt"

func main() {
	s := "Hello, World!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) //this splits each character of string s and stores the ascii value of each character
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) //this is of type []uint8
}
