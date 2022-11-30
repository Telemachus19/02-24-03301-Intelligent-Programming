package main

import "fmt"

type Car struct {
	name, model, color string
	weight             float64
}

type Product struct {
	Name, ID        string
	ManufactureYear int
}

func printProduct(p Product) {
	fmt.Printf("Name%s\nID:%s\nManufacture Year:%d\n", p.Name, p.ID, p.ManufactureYear)
}

func main() {
	car1 := Car{"Y", "Tesla", "Black", 120.2}
	p1 := Product{"Shampoo", "02-024-0025", 2017}
	p2 := Product{"Somehitng", "02-024-0026", 2017}
	p3 := Product{"Shampoo2", "02-024-0027", 2017}
	printProduct(p1)
	printProduct(p2)
	printProduct(p3)
	fmt.Print(car1)
}
