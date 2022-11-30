package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type phoneBook struct {
	personName, phoneNumber, address string
}

func main() {
	var phonebook []phoneBook
	var name, number, address string
	var input string
	var found bool
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter 1 to Add a person to the phone book")
		fmt.Println("Enter 2 to search a person in the phone book")
		fmt.Println("Enter e to exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		if strings.Compare(input, "1") == 0 {
			fmt.Print("Enter person's name: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Enter person's number: ")
			scanner.Scan()
			number = scanner.Text()
			fmt.Print("Enter person's Address: ")
			scanner.Scan()
			address = scanner.Text()

			phonebook = append(phonebook, phoneBook{personName: name, phoneNumber: number, address: address})
		} else if strings.Compare(input, "2") == 0 {
			found = false
			fmt.Print("Enter the person's name to search: ")
			scanner.Scan()
			name = scanner.Text()
			for i := 0; i < len(phonebook); i++ {
				if strings.Compare(phonebook[i].personName, name) == 0 {
					fmt.Print("The person was found\nPrinting data ...\n")
					printData(phonebook[i])
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Person was not found")
			}
		} else {
			break
		}
	}
}
func printData(phonebook phoneBook) {
	fmt.Println("---")
	fmt.Println("Person's name:", phonebook.personName)
	fmt.Println("Person's Phone Number:", phonebook.phoneNumber)
	fmt.Println("Person's address:", phonebook.address)
	fmt.Println("---")
}
