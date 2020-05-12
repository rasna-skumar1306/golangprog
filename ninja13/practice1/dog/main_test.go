package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("got", n, "expected", 70)
	}
}
func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("got", n, "expected", 70)
	}
}
func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}
func BenchmarkYears(b *testing.B) {
	v := 1
	for i := 0; i < b.N; i++ {
		Years(v)
		v++
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
