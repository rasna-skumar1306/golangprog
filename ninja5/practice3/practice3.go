//program to use multiple structs and embed them together
package main

import "fmt"

type vehicle struct {
	door  string
	color string
}
type truck struct {
	vehicle   //here vehicle struct is embedded in truck
	fourwheel bool
}
type sedan struct {
	vehicle //here vehicle struct is embedded in sedan
	luxury  bool
}

func main() {
	c := truck{
		vehicle: vehicle{
			door:  "metal",
			color: "red",
		},
		fourwheel: true,
	}

	d := sedan{
		vehicle: vehicle{
			door:  "metal",
			color: "blue",
		},
		luxury: true,
	}
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(c.fourwheel)
}
