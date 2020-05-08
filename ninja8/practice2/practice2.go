//program to unmarshal a json to a string of message
package main

import (
	"encoding/json" //json package
	"fmt"
)

type person struct { //struct got from https://mholt.github.io/json-to-go/ by pasting the json in the site.
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

var people []person

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// here s is the json data in raw string literal
	err := json.Unmarshal([]byte(s), &people) //json being unmarshalled to struct. unmarshall returns err
	if err != nil {
		fmt.Println(err) //handling the err
	}
	fmt.Println(people) //print the struct

	for i, v := range people { //orderly printing the struct according to key value pair
		fmt.Println("person ", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}

}
