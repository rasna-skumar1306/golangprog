package main

import "fmt"

func main() {
	i := 0

	s := "Hello, World!"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) //this splits each character of string s and stores the ascii value of each character
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) //this is of type []uint8

	for i = 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i]) //this %#U is used to print the UTF-8 encoding of the characters in s
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v) //this line prints ascii code of each character
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("position %d has hex value of %#x\n", i, v) //this line prints the hexadecimal value of each character
	}
}
