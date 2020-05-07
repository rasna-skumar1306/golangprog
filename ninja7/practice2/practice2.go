//program uses *struct to change name of struct using a function
package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeme(s *person, f string, l string) { //recieves a *person and changes the name
	(*s).first = f
	(*s).last = l
}

func main() {
	p := person{
		first: "rasna",
		last:  "senthilkumar",
	}
	fmt.Println(p)
	changeme(&p, "rithu", "senthilkumar") //address of struct is sent as parameter
	fmt.Println(p)
}
