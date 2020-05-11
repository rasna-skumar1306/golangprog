//error handling through function
package main

import (
	"encoding/json"
	"fmt"
	"log" //using log to print error
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1) //calling a custom function
	if err != nil {
		log.Fatalln(err) //log.fatal stops the execution of lines below it and throws the error
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil { //checking for error
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err) //returning a custom error
	}
	return bs, nil
}
