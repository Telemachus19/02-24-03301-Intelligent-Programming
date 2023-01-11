package main

import "fmt"

func main() {
	SquareSumChan := make(chan int)
	CubicSumChan := make(chan int)
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	go func(n int) {
		sum := 0
		for n > 0 {
			r := n % 10
			sum += (r * r)
			n /= 10
		}
		SquareSumChan <- sum
	}(n)
	go func(n int) {
		sum := 0
		for n > 0 {
			r := n % 10
			sum += (r * r * r)
			n /= 10
		}
		CubicSumChan <- sum
	}(n)
	SquareSum := <-SquareSumChan
	CubicSum := <-CubicSumChan
	fmt.Println("Square sum of the digits: ", SquareSum)
	fmt.Println("cubic  sum of the digits: ", CubicSum)
	fmt.Println("Summation of square sum and cubic sum of digits : ", SquareSum+CubicSum)
}
