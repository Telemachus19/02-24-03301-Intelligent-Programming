package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(numbers)
	snumbers := numbers[:3]
	fmt.Println(snumbers)
	snumbers1 := numbers[2:]
	fmt.Println(snumbers1)
	snumbers[2] = 100
	fmt.Println(numbers)
	fmt.Println(snumbers)
	fmt.Println(snumbers1)
}
