package main

import "fmt"

const end = 10

func sol_1() {

	for i := 0; i < end; i++ {
		fmt.Printf("%d ", (i+1)*(i+1))
	}
}

func sol_2() {
	num := 1
	diff := 3
	for i := 0; i < end; i, diff = i+1, diff+2 {
		fmt.Printf("%d ", num)
		num = num + diff
	}
}
func main() {
	sol_1()
	fmt.Println()
	sol_2()
}
