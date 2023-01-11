package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	var num int
	fmt.Println("Enter a number <= 10: ")
	fmt.Scan(&num)
	factMap := make(map[int]int)

	for i := 1; i <= num; i++ {
		factMap[i] = fact(i)
	}

	fmt.Println(factMap)
}
