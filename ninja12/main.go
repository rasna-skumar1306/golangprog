package main

import (
	"fmt"

	"github.com/golangprog/ninja12/dog"
)

type retriver struct {
	name string
	age  int
}

func main() {
	rithish := retriver{
		name: "Rithish",
		age:  dog.Years(14),
	}
	fmt.Println(rithish)

}
