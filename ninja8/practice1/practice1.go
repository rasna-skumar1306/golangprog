//program that marshals a struct into json format
package main

import (
	"encoding/json" //json package
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3} //slice of user struct

	fmt.Println(users)

	b, err := json.Marshal(users) //marshalling struct user into json. Marshal returns []int and err as output
	if err != nil {
		fmt.Println(err) //catching error (if any) and printing it
	}
	fmt.Println(string(b)) //[]int to string and printed
}
