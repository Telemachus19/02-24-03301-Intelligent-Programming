package main

import "fmt"

func main() {
	// define a fixed-size array
	const len = 5
	var a [len]int

	for i := 0; i < len; i++ {
		fmt.Printf("array[%d] = ", i)
		fmt.Scanln(&a[i])
	}
	for i := 0; i < len; i++ {
		if a[i]%2 == 0 {
			fmt.Printf("%d is an even element whose index is %d \n", a[i], i)
		} else {
			fmt.Printf("%d is an odd element whose index is %d\n", a[i], i)
		}
	}

}
