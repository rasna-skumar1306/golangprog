//using custom sort function for sorting struct acc to age and last name
package main

import (
	"fmt"
	"sort" //package sort for sorting various data
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

//SortByAge sorts the user by Age
type SortByAge []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } //compares the age of user

//SortByLast sorts the user by last name
type SortByLast []user

func (l SortByLast) Len() int           { return len(l) }
func (l SortByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l SortByLast) Less(i, j int) bool { return l[i].Last < l[j].Last } //compares the first character of last name of the user

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
	fmt.Println("-------------------------------")
	fmt.Println("Actual struct elements") //printing the actual elements of users without any sorting
	fmt.Println("-------------")
	for i, v := range users {
		fmt.Println("person #", i)
		fmt.Println("\t", v.First, v.Last, "age", v.Age)
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}
	fmt.Println("-------------------------------")
	fmt.Println("After sorting by Age") //sorting the users according to age
	fmt.Println("-------------")
	sort.Sort(SortByAge(users)) //custom sort function for sorting by age is called using sort.Sort(func name(args))
	for i, v := range users {
		fmt.Println("person #", i)
		fmt.Println("\t", v.First, v.Last, "age", v.Age)
		sort.Strings(v.Sayings) //sorting the sayings of the user alphabetically
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}
	fmt.Println("-------------------------------")
	fmt.Println("After sorting by Last name") //sorting the users according to last name
	fmt.Println("-------------")
	sort.Sort(SortByLast(users)) //custom sort function for sorting by age is called

	for i, v := range users {
		fmt.Println("person #", i)
		fmt.Println("\t", v.First, v.Last, "age", v.Age)
		sort.Strings(v.Sayings) //sorting the sayings of the user alphabetically
		for _, val := range v.Sayings {
			fmt.Println("\t\t", val)
		}
	}

}
