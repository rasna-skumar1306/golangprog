package main

import "fmt"

func main() {
	m := map[string][]string{	//here map is string key and []string value pair
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["ian_fleming"] = []string{`love for bond`,`director`,`novel`} // appending an key value pair to the map

	for k, v := range m {	//printing the map
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	delete(m,"no_dr")	//deleting no_dr from map m
	
	fmt.Println("after deleting an element")
	
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

/*
`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
*/
