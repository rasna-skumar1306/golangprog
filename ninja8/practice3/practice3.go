//program that encodes a struct to json format
package main

import (
	"encoding/json" //for encoding to json
	"fmt"
	"os" // for writing the output to console using os.stdout this package is used
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	//encoder takes in a writer (in this case (os.stdout)) and Encode as method with argument users
	//instead of os.stdout we can also use other writers such as file writers, html writers, etc...
	err := json.NewEncoder(os.Stdout).Encode(users) //encoder encoding the user struct to json
	if err != nil {
		fmt.Println(err)
	}

}
