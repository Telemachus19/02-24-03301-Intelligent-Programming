package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	id, name string
	age      int
}

func main() {
	var persons []Person
	fmt.Println(len(persons))
	fmt.Println(cap(persons))
	scanner := bufio.NewScanner(os.Stdin)

	for input := ""; true; input = "" {
		fmt.Print("For adding a new Person enter 1 or to the end to program press any key\n")
		fmt.Scanln(&input)
		fmt.Println("you have entered:", input)
		if input != "1" {
			break
		}
		persons = append(persons, func() Person {
			var id, name string
			var age int
			fmt.Print("Enter The person's id: ")
			fmt.Scanln(&id)
			fmt.Print("Enter the person's name: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Enter the person's age: ")
			fmt.Scanln(&age)
			return Person{id, name, age}
		}())
	}

	fmt.Println("Length of the persons slice:", len(persons))
	fmt.Println("Capacity of the persons slice:", cap(persons))
	func(persons_slice []Person) {
		for i := 0; i < len(persons_slice); i++ {
			fmt.Println("--------")
			fmt.Printf("Person %d\n", i)
			fmt.Printf("ID: %s \t Name: %s \t Age: %d\n", persons_slice[i].id, persons_slice[i].name, persons_slice[i].age)
			fmt.Println("--------")
		}
	}(persons)
}
