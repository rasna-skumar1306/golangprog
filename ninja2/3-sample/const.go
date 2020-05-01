package main

import "fmt"

const a = 42 //this is one way to declare a const

const (
	b = 42.36
	c = "Rasswanth" //this is the collective way to declare a const
	d = iota        //this is a iota variable which is predefined to 0.....n
	e = iota        //automatically gets the value 3
	f = iota        //automatically gets the value 4
)

const ( //or iota can be intialized like this
	g = iota
	h
	i
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d, e, f)
	fmt.Printf("%T\n", d)
	fmt.Println(g, h, i)
}
