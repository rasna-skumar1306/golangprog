package main

import (
	"fmt"
	"github.com/golangprog/ninja13/practice2/01-code-starting/quote"
	"github.com/golangprog/ninja13/practice2/01-code-starting/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
