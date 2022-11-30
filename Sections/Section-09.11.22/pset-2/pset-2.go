package main

import (
	"fmt"
)

func main() {
	var a [5]int
	var even_part []int
	var odd_part []int
	for i := 0; i < len(a); i++ {
		fmt.Printf("array[%d] = ", i)
		fmt.Scanln(&a[i])
	}
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			fmt.Printf("%d is an Even element whose index is %d\n", a[i], i)
			even_part = append(even_part, a[i])
		} else {
			fmt.Printf("%d is an Odd element whose index is %d\n", a[i], i)
			odd_part = append(odd_part, a[i])
		}
	}
	fmt.Println("Even part:", even_part)
	fmt.Println("Odd part:", odd_part)
}
