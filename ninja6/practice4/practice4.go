//program that uses a method of a used defined struct
package main

import "fmt"

type person struct { //user defined struct
	first string
	last  string
	age   int
}

func (p person) speak() { //method that listens to type person
	fmt.Println("I am,", p.first, " and my age is", p.age)
}

func main() {
	p := person{
		first: "rasna",
		last:  "senthilkumar",
		age:   19,
	}
	p.speak() //calling the method speak that prints the name and age of the person
}
