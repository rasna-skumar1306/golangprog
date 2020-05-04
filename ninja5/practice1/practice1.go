//program that combines struct with slice
package main

import "fmt"

type nume struct { //declaration of struct
	first string
	last  string
	fav   []string
}

func main() {
	p1 := nume{
		first: "James",
		last:  "Bond",
		fav: []string{
			"vodka",
			"martini",
			"vesper",
		},
	}
	p2 := nume{
		first: "Miss",
		last:  "Moneypenny",
		fav: []string{
			"chocolate",
			"gin",
			"flowers",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.fav {
		fmt.Printf("%v\t%v\n", i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.fav {
		fmt.Printf("%v\t%v\n", i, v)
	}
}
