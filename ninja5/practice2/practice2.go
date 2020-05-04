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
	m := map[string]nume{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Printf("%v\n", v.first)
		fmt.Printf("%v\n", v.last)
		for i, val := range v.fav {
			fmt.Printf("%v\t %v\n", i, val)
		}
		fmt.Println("----------------------------------")
	}
}
