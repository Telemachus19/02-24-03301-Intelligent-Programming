package main

import (
	"bufio"
	"fmt"
	"os"
)

type employee struct {
	name, position string
	id             int
}

func main() {
	// 1 for entering new data
	// 2 for updating the record
	// 3 for deleting the record
	// Any other for exiting the main loop
	m := make(map[employee]float64)

	scanner := bufio.NewScanner(os.Stdin)
	var e employee
	var sar float64
	fmt.Print("Enter the employee name: ")
	scanner.Scan()
	e.name = scanner.Text()
	fmt.Print("Enter the employee's position: ")
	fmt.Scan(&e.position)
	fmt.Print("Enter the employee's id: ")
	fmt.Scan(&e.id)
	fmt.Print("enter the salary of the employee: ")
	fmt.Scan(&sar)
	m[e] = sar
	fmt.Print(m)
	m[e] = 1500.12
	fmt.Print(m)
	delete(m, e)
}
