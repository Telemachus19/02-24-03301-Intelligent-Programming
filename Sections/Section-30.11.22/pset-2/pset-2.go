package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	resultMatrix := [][]int{{0, 0}, {0, 0}}

	var sum = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			sum += matrix[i][j]
		}
	}
	fmt.Println("Summation of elements of the matrix", sum)
	printMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			resultMatrix[i][j] = matrix[i][j] * matrix[i][j]
		}
	}
	fmt.Println("Product Matrix")
	printMatrix(resultMatrix)
}

func printMatrix(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d \t", matrix[i][j])
		}
		fmt.Println()
	}
}
