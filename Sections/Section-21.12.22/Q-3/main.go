package main

import (
	"bufio"
	"fmt"
	"os"
)

type item struct {
	id, price int
	name      string
}

type game struct {
	genre string
	i     item
}

func printGames(games []game) {
	fmt.Println("Id    name    Price    Genre")
	for _, i := range games {
		fmt.Printf("%v  %v    %v    %v\n", i.i.id, i.i.name, i.i.price, i.genre)
		fmt.Println("--------")
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	gameSlices := make([]game, 0)
	var id, price int
	var name, genre string
	var itemx item
	// Enter 1 to continue adding game, anything else to exit
	for input := ""; true; input = "" {
		fmt.Print("Enter 1 to add games, or any character to exit: ")
		fmt.Scanln(&input)
		if input != "1" {
			break
		}
		fmt.Print("Enter the id of the game: ")
		fmt.Scanln(&id)
		fmt.Print("Enter name of the game: ")
		scanner.Scan()
		name = scanner.Text()
		fmt.Print("Enter the price of the game: ")
		fmt.Scanln(&price)
		itemx = item{id: id, name: name, price: price}
		fmt.Print("Enter the genre of the game: ")
		scanner.Scan()
		genre = scanner.Text()
		gameSlices = append(gameSlices, game{i: itemx, genre: genre})
	}
	printGames(gameSlices)
}
