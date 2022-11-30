package main

import "fmt"

func main() {
	var a [8]int
	var equalityNumber int
	for i := 0; i < len(a); i++ {
		fmt.Printf("Enter element[%d]: ", i)
		fmt.Scanln(&a[i])
	}
	fmt.Print("Enter a number to test equality: ")
	fmt.Scanln(&equalityNumber)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == equalityNumber {
				fmt.Printf("%d + %d = %d\n", a[i], a[j], equalityNumber)
			}
		}
	}
}
