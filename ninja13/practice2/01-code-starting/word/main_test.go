package word

import (
	"fmt"
	"testing"

	"github.com/golangprog/ninja13/practice2/01-code-starting/quote"
)

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}
func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "expected", 3)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
