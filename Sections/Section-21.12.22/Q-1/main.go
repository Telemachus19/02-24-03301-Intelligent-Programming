package main

import "fmt"

func calculateSquares(num int) int {
	var sum int
	for num > 0 {
		r := num % 10
		sum += (r * r)
		num /= 10
	}
	return sum
}
func calculateCubes(num int) int {
	var sum int
	for num > 0 {
		r := num % 10
		sum += (r * r * r)
		num /= 10
	}
	return sum
}
func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	fmt.Println(calculateSquares(number) + calculateCubes(number))
}
