package main

import "fmt"

func main() {
	const rows = 4
	const cols = 4
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
