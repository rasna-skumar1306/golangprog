package word

import "strings"

//UseCount returns the length of argument passed
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns the length of argument passed to it
func Count(s string) int {
	xs:= strings.Fields(s)
	return len(xs)
}
