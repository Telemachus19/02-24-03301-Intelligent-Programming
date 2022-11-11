package main

import "fmt"

type Coordinates struct {
	x_axis, y_axis float64
	color          string
}
type Profile struct {
	Name, PhoneNumber string
	age               int
}

func main() {
	var c1 Coordinates = Coordinates{x_axis: 12.5,
		y_axis: -12.5,
		color:  "red"}
	fmt.Println(c1)
	var p1 Profile = Profile{Name: "Ahmed Ashraf", PhoneNumber: "+20 101 364 8328", age: 21}
	fmt.Println(p1)

	fmt.Printf("Hello, There!\n Let my tell more about myself.\n My name is %s, my phone number is %s, and I'm %d years old.\n", p1.Name, p1.PhoneNumber, p1.age)
	fmt.Printf("The coordinates are\nx:%f\ny:%f\nColor:%s", c1.x_axis, c1.y_axis, c1.color)
}
