package main

import "fmt"

func Q1() {
	var sum int64 = 0
	var i int64 = 3
	var j int64 = 5
	const MAXCOUNT = 10
	for i <= MAXCOUNT {
		sum += i
		i += 3
	}
	for j <= MAXCOUNT {
		sum += j
		j += 5
	}
	fmt.Println(sum)
}
func Q1Formula() {
	var sumOfThrees = 33 * (2*3 + (32 * 3))
	sumOfThrees = sumOfThrees / 2
	var sumOfFives = 20 * (2*5 + (19 * 5))
	sumOfFives = sumOfFives / 2
	fmt.Println(sumOfFives + sumOfThrees)
}
func Q2() {
	var squareOfSum int = 0
	var sumOfSquares int = 0
	const MAXCOUNT = 10
	for i := 0; i <= MAXCOUNT; i++ {
		squareOfSum += i
	}
	squareOfSum *= squareOfSum
	fmt.Println("Square of the sum: ", squareOfSum)
	for i := 0; i <= MAXCOUNT; i++ {
		sumOfSquares += i * i
	}
	fmt.Println("Sum of Squares: ", sumOfSquares)
	fmt.Println("Difference between them: ", squareOfSum-sumOfSquares)
}
func Q3If(input int) {
	fmt.Println("The following result used if Statements")
	if input == 1 {
		fmt.Println("Today is Saturday")
	} else if input == 2 {
		fmt.Println("Today is Sunday")
	} else if input == 3 {
		fmt.Println("Today is Monday")
	} else if input == 4 {
		fmt.Println("Today is Tuesday")
	} else if input == 5 {
		fmt.Println("Today is Wednesday")
	} else if input == 6 {
		fmt.Println("Today is Thursday")
	} else if input == 7 {
		fmt.Println("Today is Friday")
	} else {
		fmt.Println("I can't determine what day is it")
	}
}
func Q3Switch(input int) {
	fmt.Println("The following result used switch statements")
	switch input {
	case 1:
		fmt.Println("Today is Saturday")
	case 2:
		fmt.Println("Today is Sunday")
	case 3:
		fmt.Println("Today is Monday")
	case 4:
		fmt.Println("Today is Tuesday")
	case 5:
		fmt.Println("Today is Wednesday")
	case 6:
		fmt.Println("Today is Thursday")
	case 7:
		fmt.Println("Today is Friday")
	default:
		fmt.Println("I can't determine what day is it")
	}
}
func main() {
	Q1()
	Q1Formula()
	Q2()
	var day int = 0
	fmt.Print("Enter the number of day: ")
	fmt.Scanf("%d", &day)
	fmt.Println("The number of the day is: ", day)
	Q3If(day)
	Q3Switch(day)
}
