//program that exhibits method set
package main

/*
In this program I have exhibited the property of method sets
reciever | struct type
-----------------------
(t T)	 | T, *T	reciever type T can recieve both actual and Pointer T
(t *T)	 | *T		reciever type *T can recieve only pointer T
*/
import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.last, ",", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	h := person{
		"Rasna",
		"Senthilkumar",
	}

	/* saySomething(h) */ //this won't work as per method set
	saySomething(&h)      //as this has the address of h, this will work
	h.speak()             //this is not a pointer value but it calls speak function and it still works
}
