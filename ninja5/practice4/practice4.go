//a complex program that has anonympus struct
package main

import "fmt"

func main() {

	s := struct { //declaration of anonymous struct
		first     string         //string field
		friends   map[string]int //a map field inside the struct
		favDrinks []string       //a slice field inside the struct
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
