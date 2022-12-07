package main

import "fmt"

func main() {
	var i int
	fmt.Print("Enter a number to generate a map: ")
	fmt.Scan(&i)

	m := make(map[int]int)

	for j := 1; j <= i; j++ {
		m[j] = j * j
	}

	fmt.Println(m)

}
